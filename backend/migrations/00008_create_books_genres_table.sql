-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS book_genres (
    book_id  UUID REFERENCES books(id) ON DELETE CASCADE,
    genre_id UUID REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, genre_id)
);

CREATE INDEX IF NOT EXISTS idx_book_genres_genre ON book_genres(genre_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_book_genres_genre;
DROP TABLE IF EXISTS book_genres;
-- +goose StatementEnd
