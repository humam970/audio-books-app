-- name: CreateAuthor :one
INSERT INTO authors (name, bio)
VALUES (@name, @bio)
RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors WHERE id = @id LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: UpdateAuthor :one
UPDATE authors
SET name = @name, bio = @bio
WHERE id = @id
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = @id;

-- name: AddAuthorToBook :exec
INSERT INTO book_authors (book_id, author_id)
VALUES (@book_id, @author_id) ON CONFLICT DO NOTHING;

-- name: RemoveAuthorFromBook :exec
DELETE FROM book_authors
WHERE book_id = @book_id AND author_id = @author_id;

-------------------------------------------------------------------------------------------------------

-- name: CreateNarrator :one
INSERT INTO narrators (name, bio)
VALUES (@name, @bio)
RETURNING *;

-- name: GetNarrator :one
SELECT * FROM narrators WHERE id = @id LIMIT 1;

-- name: ListNarrators :many
SELECT * FROM narrators ORDER BY name;

-- name: UpdateNarrator :one
UPDATE narrators
SET name = @name, bio = @bio
WHERE id = @id
RETURNING *;

-- name: DeleteNarrator :exec
DELETE FROM narrators WHERE id = @id;

-- name: AddNarratorToBook :exec
INSERT INTO book_narrators (book_id, narrator_id)
VALUES (@book_id, @narrator_id) ON CONFLICT DO NOTHING;

-- name: RemoveNarratorFromBook :exec
DELETE FROM book_narrators
WHERE book_id = @book_id AND narrator_id = @narrator_id;

-------------------------------------------------------------------------------------------------------

-- name: CreateGenre :one
INSERT INTO genres (name) VALUES (@name) RETURNING *;

-- name: GetGenre :one
SELECT * FROM genres WHERE id = @id LIMIT 1;

-- name: ListGenres :many
SELECT * FROM genres ORDER BY name;

-- name: DeleteGenre :exec
DELETE FROM genres WHERE id = @id;

-- name: AddGenreToBook :exec
INSERT INTO book_genres (book_id, genre_id)
VALUES (@book_id, @genre_id) ON CONFLICT DO NOTHING;

-- name: RemoveGenreFromBook :exec
DELETE FROM book_genres
WHERE book_id = @book_id AND genre_id = @genre_id;

-- name: GetGenresForBook :many
SELECT g.* FROM genres g
JOIN book_genres bg ON g.id = bg.genre_id
WHERE bg.book_id = @id;

-------------------------------------------------------------------------------------------------------

-- name: CreateBook :one
INSERT INTO books (
    title,
    duration_seconds,
    rating,
    release_date,
    cover_image_url,
    audio_preview_url,
    is_abridged
) VALUES (
    @title,
    @duration_seconds,
    @rating,
    @release_date,
    @cover_image_url,
    @audio_preview_url,
    @is_abridged
)
RETURNING *;

-- name: GetBook :one
SELECT
    b.*,
    array_agg(DISTINCT a.name)::text[] as authors,
    array_agg(DISTINCT n.name)::text[] as narrators
FROM books b
LEFT JOIN book_authors ba ON b.id = ba.book_id
LEFT JOIN authors a ON ba.author_id = a.id
LEFT JOIN book_narrators bn ON b.id = bn.book_id
LEFT JOIN narrators n ON bn.narrator_id = n.id
WHERE b.id = @id
GROUP BY b.id
LIMIT 1;

-- name: ListBooks :many
SELECT
    b.id,
    b.title,
    b.release_date,
    b.cover_image_url,
    b.rating,
    array_agg(DISTINCT a.name)::text[] as authors
FROM books b
LEFT JOIN book_authors ba ON b.id = ba.book_id
LEFT JOIN authors a ON ba.author_id = a.id
GROUP BY b.id
ORDER BY b.release_date DESC;

-- name: UpdateBook :one
UPDATE books
SET
    title           = @title,
    rating          = @rating,
    is_abridged     = @is_abridged
WHERE id = @id
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = @id;

-------------------------------------------------------------------------------------------------------

-- name: CreateChapter :one
INSERT INTO chapters (book_id, title, start_time, end_time, order_index)
VALUES (@book_id, @title, @start_time, @end_time, @order_index)
RETURNING *;

-- name: ListChaptersByBook :many
SELECT * FROM chapters
WHERE book_id = @book_id
ORDER BY order_index ASC;

-- name: UpdateChapter :one
UPDATE chapters
SET
    title      = @title,
    start_time = @start_time,
    end_time   = @end_time,
    order_index = @order_index
WHERE id = @id
RETURNING *;

-- name: DeleteChapter :exec
DELETE FROM chapters WHERE id = @id;

-------------------------------------------------------------------------------------------------------
