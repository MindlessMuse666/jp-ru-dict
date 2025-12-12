package model

import (
	"time"
)

// Word представляет слово в словаре пользователя
// Содержит основную информацию о японском слове и его переводе
type Word struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Jp        []string  `json:"jp" db:"jp" validate:"required,min=1"`
	Ru        []string  `json:"ru" db:"ru" validate:"required,min=1"`
	On        []string  `json:"on,omitempty" db:"on"`
	Kun       []string  `json:"kun,omitempty" db:"kun"`
	ExJp      []string  `json:"ex_jp,omitempty" db:"ex_jp" validate:"max=3"`
	ExRu      []string  `json:"ex_ru,omitempty" db:"ex_ru" validate:"max=3"`
	Tags      []string  `json:"tags,omitempty" db:"tags" validate:"max=5"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// WordCreateRequest представляет запрос на создание слова
type WordCreateRequest struct {
	Jp   []string `json:"jp" binding:"required,min=1"`
	Ru   []string `json:"ru" binding:"required,min=1"`
	On   []string `json:"on,omitempty"`
	Kun  []string `json:"kun,omitempty"`
	ExJp []string `json:"ex_jp,omitempty" binding:"max=3"`
	ExRu []string `json:"ex_ru,omitempty" binding:"max=3"`
	Tags []string `json:"tags,omitempty" binding:"max=5"`
}

// WordUpdateRequest представляет запрос на обновление слова
type WordUpdateRequest struct {
	Jp   []string `json:"jp,omitempty" binding:"min=1"`
	Ru   []string `json:"ru,omitempty" binding:"min=1"`
	On   []string `json:"on,omitempty"`
	Kun  []string `json:"kun,omitempty"`
	ExJp []string `json:"ex_jp,omitempty" binding:"max=3"`
	ExRu []string `json:"ex_ru,omitempty" binding:"max=3"`
	Tags []string `json:"tags,omitempty" binding:"max=5"`
}

// WordsListResponse представляет ответ со списком слов
type WordsListResponse struct {
	Words      []*Word `json:"words"`
	NextCursor int     `json:"next_cursor,omitempty"`
	HasMore    bool    `json:"has_more"`
}
