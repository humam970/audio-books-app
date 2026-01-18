-- +goose UP
-- +goose StatementBegin
CREATE TABLE authors (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    bio  TEXT NOT NULL
);

CREATE TABLE narrators (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    bio  TEXT
);

CREATE TABLE books (
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

CREATE TABLE book_authors (
    book_id   UUID REFERENCES books(id) ON DELETE CASCADE,
    author_id UUID REFERENCES authors(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
);

CREATE TABLE book_narrators (
    book_id     UUID REFERENCES books(id) ON DELETE CASCADE,
    narrator_id UUID REFERENCES narrators(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, narrator_id)
);

CREATE TABLE chapters (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id     UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    title       TEXT NOT NULL,
    start_time  INT NOT NULL, 
    end_time    INT NOT NULL,
    order_index INT NOT NULL
);

CREATE TABLE genres (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE book_genres (
    book_id  UUID REFERENCES books(id) ON DELETE CASCADE,
    genre_id UUID REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, genre_id)
);

CREATE INDEX idx_book_authors_author ON book_authors(author_id);
CREATE INDEX idx_book_narrators_narrator ON book_narrators(narrator_id);
CREATE INDEX idx_book_genres_genre ON book_genres(genre_id);
CREATE INDEX idx_chapters_book_id ON chapters(book_id);
-- +goose StatementEnd
