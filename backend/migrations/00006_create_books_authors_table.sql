-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS book_authors (
    book_id   UUID REFERENCES books(id) ON DELETE CASCADE,
    author_id UUID REFERENCES authors(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
);

CREATE INDEX IF NOT EXISTS idx_book_authors_author ON book_authors(author_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_book_authors_author;
DROP TABLE IF EXISTS book_authors;
-- +goose StatementEnd
