-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chapters (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id     UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    title       TEXT NOT NULL,
    start_time  INT NOT NULL, 
    end_time    INT NOT NULL,
    order_index INT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_chapters_book_order ON chapters(book_id, order_index);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_chapters_book_order;
DROP TABLE IF EXISTS chapters;
-- +goose StatementEnd

