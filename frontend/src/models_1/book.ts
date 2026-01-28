export type Book = {
    id: string;
    title: string;
    duration_seconds: number;
    rating: number;
    release_date: Date;
    cover_image_url: string;
    audio_preview_url: string;
    is_abridged: boolean;
    created_at: Date;
};

export type CreateBookRequest = {
    title: string;
    duration_seconds: number;
    rating: number;
    release_date: Date;
    cover_image_url: string;
    audio_preview_url: string;
    is_abridged: boolean;
};

export type UpdateBookRequest = Partial<Pick<CreateBookRequest, "title" | "rating" | "is_abridged">>;

// ------------------------------------------------------------------

export type BookAuthor = {
    book_id: string;
    author_id: string;
};

export type AddAuthorToBookRequest = {
    author_id: string;
};

export type AddAuthorToBookResponse = {
    book_id: string;
    author_id: string;
};

export type RemoveAuthorFromBookRequest = {
    author_id: string;
};

// ------------------------------------------------------------------

export type BookNarrator = {
    book_id: string;
    narrator_id: string;
};

export type AddNarratorToBookRequest = {
    narrator_id: string;
};

export type AddNarratorToBookResponse = {
    book_id: string;
    narrator_id: string;
};

export type RemoveNarratorFromBookRequest = {
    narrator_id: string;
};

// ------------------------------------------------------------------

export type BookGenre = {
    book_id: string;
    genre_id: string;
};

export type AddGenreToBookRequest = {
    genre_id: string;
};

export type AddGenreToBookResponse = {
    book_id: string;
    genre_id: string;
};

export type RemoveGenreFromBookRequest = {
    genre_id: string;
};
