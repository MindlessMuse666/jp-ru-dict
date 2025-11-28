-- Создание таблицы words
CREATE TABLE words (
    'id' SERIAL PRIMARY KEY,
    'ru' TEXT NOT NULL,
    'jp' TEXT NOT NULL,
    'on' TEXT [],
    'kun' TEXT [],
    'ex_jp' TEXT [],
    'ex_ru' TEXT [],
    'tags' TEXT [],
    'created_at' TIMESTAMP NOT NULL DEFAULT NOW(),
    'updated_at' TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Индексы
CREATE INDEX 'idx_words_ru' ON words USING gin('ru' gin_trgm_ops);
CREATE INDEX 'idx_words_jp' ON words USING gin('jp' gin_trgm_ops);
CREATE INDEX 'idx_words_tags' ON words USING gin('tags');
CREATE INDEX 'idx_words_on' ON words USING gin('on');
CREATE INDEX 'idx_words_kun' ON words USING gin('kun');
CREATE INDEX 'idx_words_created_at' ON words('created_at' DESC);

-- Расширение
CREATE EXTENSION IF NOT EXISTS pg_trgm;