import z from "zod";
import { uuidSchema } from "./_internal";

export const ChapterSchema = z
    .object({
        id: uuidSchema,
        book_id: uuidSchema,
        title: z.string(),
        start_time: z.number(),
        end_time: z.number(),
        order_index: z.int().positive(),
    })
    .refine((data) => data.start_time < data.end_time, {
        error: "Chapter start time must be less than its end time",
        path: ["end_time"],
    });

export const CreateChapterRequestSchema = z.object({
    title: z.string(),
    start_time: z.number(),
    end_time: z.number(),
    order_index: z.int().positive(),
});

export type CreateChapterRequest = z.infer<typeof CreateChapterRequestSchema>;

export const UpdateChapterRequestSchema = z.object({
    title: z.optional(z.string()),
    start_time: z.optional(z.number()),
    end_time: z.optional(z.number()),
    order_index: z.optional(z.int().positive()),
});

export type UpdateChapterRequest = z.infer<typeof UpdateChapterRequestSchema>;
