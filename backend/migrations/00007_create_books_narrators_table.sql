-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS book_narrators (
    book_id     UUID REFERENCES books(id) ON DELETE CASCADE,
    narrator_id UUID REFERENCES narrators(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, narrator_id)
);

CREATE INDEX IF NOT EXISTS idx_book_narrators_narrator ON book_narrators(narrator_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_book_narrators_narrator;
DROP TABLE IF EXISTS book_narrators;
-- +goose StatementEnd
