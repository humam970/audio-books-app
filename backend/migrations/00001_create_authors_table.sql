-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS authors (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    bio  TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS authors;
-- +goose StatementEnd
