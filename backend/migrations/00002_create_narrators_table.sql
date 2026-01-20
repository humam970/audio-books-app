-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS narrators (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    bio  TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS narrators;
-- +goose StatementEnd
