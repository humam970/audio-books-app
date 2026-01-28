import z from "zod";
import { uuidSchema } from "./_internal";

const genreNameSchema = z.string().min(5).max(25);

export const GenreSchema = z.object({
    id: uuidSchema,
    name: genreNameSchema,
});

export type Genre = z.infer<typeof GenreSchema>;

export const CreateGenreRequestSchema = z.object({
    name: genreNameSchema,
});

export type CreateGenreRequest = z.infer<typeof CreateGenreRequestSchema>;
