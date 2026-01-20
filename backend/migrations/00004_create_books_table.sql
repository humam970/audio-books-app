-- +goose UP
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title             TEXT NOT NULL,
    duration_seconds  INT NOT NULL, 
    rating            DECIMAL(3, 2) NOT NULL DEFAULT 0.00, 
    release_date      DATE NOT NULL,
    cover_image_url   TEXT NOT NULL,
    audio_preview_url TEXT NOT NULL,
    is_abridged       BOOLEAN NOT NULL DEFAULT FALSE,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
