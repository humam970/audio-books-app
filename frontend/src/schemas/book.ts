import z from "zod";
import { uuidSchema } from "./_internal";

export const BookSchema = z.object({
    id: uuidSchema,
    title: z.string(),
    duration_seconds: z.number().positive(),
    rating: z.string(),
    release_date: z.date(),
    cover_image_url: z.string(),
    audio_preview_url: z.string(),
    is_abridged: z.boolean(),
    created_at: z.date(),
});

export type Book = z.infer<typeof BookSchema>;

export const CreateBookRequestSchema = z.object({
    title: z.string(),
    duration_seconds: z.number().positive(),
    rating: z.string(),
    release_date: z.date(),
    cover_image_url: z.string(),
    audio_preview_url: z.string(),
    is_abridged: z.boolean(),
});

export type CreateBookRequest = z.infer<typeof CreateBookRequestSchema>;

export const UpdateBookRequestSchema = z.object({
    title: z.optional(z.string()),
    rating: z.optional(z.string()),
    is_abridged: z.optional(z.boolean()),
});

export type UpdateBookRequest = z.infer<typeof UpdateBookRequestSchema>;

// ------------------------------------------------------------------

export const BookAuthorSchema = z.object({
    book_id: uuidSchema,
    author_id: uuidSchema,
});

export type BookAuthor = z.infer<typeof BookAuthorSchema>;

export const AddAuthorToBookRequestSchema = z.object({
    author_id: uuidSchema,
});

export type AddAuthorToBookRequest = z.infer<typeof AddAuthorToBookRequestSchema>;

export const AddAuthorToBookResponseSchema = z.object({
    book_id: uuidSchema,
    author_id: uuidSchema,
});

export type AddAuthorToBookResponse = z.infer<typeof AddAuthorToBookResponseSchema>;

export const RemoveAuthorFromBookRequestSchema = z.object({
    author_id: uuidSchema,
});

export type RemoveAuthorFromBookRequest = z.infer<typeof RemoveAuthorFromBookRequestSchema>;

// ------------------------------------------------------------------

export const BookNarratorSchema = z.object({
    book_id: uuidSchema,
    narrator_id: uuidSchema,
});

export type BookNarrator = z.infer<typeof BookNarratorSchema>;

export const AddNarratorToBookRequestSchema = z.object({
    narrator_id: uuidSchema,
});

export type AddNarratorToBookRequest = z.infer<typeof AddNarratorToBookRequestSchema>;

export const AddNarratorToBookResponseSchema = z.object({
    book_id: uuidSchema,
    narrator_id: uuidSchema,
});

export type AddNarratorToBookResponse = z.infer<typeof AddNarratorToBookResponseSchema>;

export const RemoveNarratorFromBookRequestSchema = z.object({
    narrator_id: uuidSchema,
});

export type RemoveNarratorFromBookRequest = z.infer<typeof RemoveNarratorFromBookRequestSchema>;

// ------------------------------------------------------------------

export const BookGenreSchema = z.object({
    book_id: uuidSchema,
    genre_id: uuidSchema,
});

export type BookGenre = z.infer<typeof BookGenreSchema>;

export const AddGenreToBookRequestSchema = z.object({
    genre_id: uuidSchema,
});

export type AddGenreToBookRequest = z.infer<typeof AddGenreToBookRequestSchema>;

export const AddGenreToBookResponseSchema = z.object({
    book_id: uuidSchema,
    genre_id: uuidSchema,
});

export type AddGenreToBookResponse = z.infer<typeof AddGenreToBookResponseSchema>;

export const RemoveGenreFromBookRequestSchema = z.object({
    genre_id: uuidSchema,
});

export type RemoveGenreFromBookRequest = z.infer<typeof RemoveGenreFromBookRequestSchema>;
