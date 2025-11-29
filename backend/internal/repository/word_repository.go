package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jp-ru-dict/internal/models"

	"github.com/lib/pq"
)

// WordRepository определяет контракт для работы со словами в базе данных
type WordRepository interface {
	GetWords(ctx context.Context, limit, offset int, search, tags, on, kun string) ([]models.Word, int, error)
	GetWordByID(ctx context.Context, id int) (*models.Word, error)
	CreateWord(ctx context.Context, word *models.Word) error
	UpdateWord(ctx context.Context, id int, word *models.Word) error
	DeleteWord(ctx context.Context, id int) error
}

type wordRepository struct {
	db *sql.DB
}

// NewWordRepository создает новый экземпляр репозитория слов
func NewWordRepository(db *sql.DB) WordRepository {
	return &wordRepository{db: db}
}

// GetWords возвращает список слов с пагинацией и фильтрацией
// Также возвращает общее количество слов для пагинации
func (r *wordRepository) GetWords(ctx context.Context, limit, offset int, search, tags, on, kun string) ([]models.Word, int, error) {
	// Получаем общее количество для метаданных
	countQuery := "SELECT COUNT(*) FROM words WHERE 1=1"
	countArgs := []any{}
	countQuery, countArgs = r.addSearchConditions(countQuery, countArgs, search, tags, on, kun)

	var total int
	err := r.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count words: %w", err)
	}

	// Затем получаем сами данные
	query := `
        SELECT id, ru, jp, on, kun, ex_jp, ex_ru, tags, created_at, updated_at
        FROM words WHERE 1=1
    `
	args := []any{}
	query, args = r.addSearchConditions(query, args, search, tags, on, kun)
	query += " ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query words: %w", err)
	}
	defer rows.Close()

	var words []models.Word
	for rows.Next() {
		var word models.Word
		err := rows.Scan(
			&word.ID, &word.Ru, &word.Jp, &word.On, &word.Kun,
			&word.ExJP, &word.ExRU, &word.Tags, &word.CreatedAt, &word.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan word: %w", err)
		}
		words = append(words, word)
	}

	return words, total, nil
}

// addSearchConditions добавляет условия поиска к SQL запросу
func (r *wordRepository) addSearchConditions(query string, args []interface{}, search, tags, on, kun string) (string, []interface{}) {
	argPos := len(args) + 1

	if search != "" {
		// Поиск по массивам ru и jp
		query += fmt.Sprintf(" AND (EXISTS (SELECT 1 FROM unnest(ru) AS r WHERE r ILIKE $%d) OR EXISTS (SELECT 1 FROM unnest(jp) AS j WHERE j ILIKE $%d))", argPos, argPos)
		args = append(args, "%"+search+"%")
		argPos++
	}
	if tags != "" {
		query += fmt.Sprintf(" AND $%d = ANY(tags)", argPos)
		args = append(args, tags)
		argPos++
	}
	if on != "" {
		query += fmt.Sprintf(" AND $%d = ANY(on)", argPos)
		args = append(args, on)
		argPos++
	}
	if kun != "" {
		query += fmt.Sprintf(" AND $%d = ANY(kun)", argPos)
		args = append(args, kun)
		argPos++
	}

	return query, args
}

// GetWordByID возвращает слово по его идентификатору
func (r *wordRepository) GetWordByID(ctx context.Context, id int) (*models.Word, error) {
	query := `SELECT id, ru, jp, on, kun, ex_jp, ex_ru, tags, created_at, updated_at 
              FROM words WHERE id = $1`

	var word models.Word
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&word.ID, &word.Ru, &word.Jp, &word.On, &word.Kun,
		&word.ExJP, &word.ExRU, &word.Tags, &word.CreatedAt, &word.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("word not found")
		}
		return nil, fmt.Errorf("failed to get word: %w", err)
	}

	return &word, nil
}

// CreateWord создает новое слово в базе данных
func (r *wordRepository) CreateWord(ctx context.Context, word *models.Word) error {
	query := `
        INSERT INTO words (ru, jp, on, kun, ex_jp, ex_ru, tags, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `

	// Преобразуем slice в pq.StringArray для работы с PostgreSQL массивами
	ruArray := pq.Array(word.Ru)
	jpArray := pq.Array(word.Jp)
	onArray := pq.Array(word.On)
	kunArray := pq.Array(word.Kun)
	exJPArray := pq.Array(word.ExJP)
	exRUArray := pq.Array(word.ExRU)
	tagsArray := pq.Array(word.Tags)

	err := r.db.QueryRowContext(ctx, query,
		ruArray, jpArray, onArray, kunArray, exJPArray, exRUArray, tagsArray,
	).Scan(&word.ID, &word.CreatedAt, &word.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create word: %w", err)
	}

	return nil
}

// UpdateWord обновляет существующее слово в базе данных
func (r *wordRepository) UpdateWord(ctx context.Context, id int, word *models.Word) error {
	query := `
        UPDATE words 
        SET ru = $1, jp = $2, on = $3, kun = $4, 
            ex_jp = $5, ex_ru = $6, tags = $7, updated_at = NOW()
        WHERE id = $8
        RETURNING updated_at
    `

	ruArray := pq.Array(word.Ru)
	jpArray := pq.Array(word.Jp)
	onArray := pq.Array(word.On)
	kunArray := pq.Array(word.Kun)
	exJPArray := pq.Array(word.ExJP)
	exRUArray := pq.Array(word.ExRU)
	tagsArray := pq.Array(word.Tags)

	err := r.db.QueryRowContext(ctx, query,
		ruArray, jpArray, onArray, kunArray, exJPArray, exRUArray, tagsArray, id,
	).Scan(&word.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("word not found")
		}
		return fmt.Errorf("failed to update word: %w", err)
	}

	word.ID = id
	return nil
}

// DeleteWord удаляет слово из базы данных по идентификатору
func (r *wordRepository) DeleteWord(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM words WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete word: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("word not found")
	}

	return nil
}
