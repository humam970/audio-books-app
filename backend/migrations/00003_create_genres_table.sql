-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS genres (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS genres;
-- +goose StatementEnd
