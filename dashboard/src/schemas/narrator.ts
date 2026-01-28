import z from "zod";
import { userBioSchema, uuidSchema, passwordSchema, userNameSchema } from "./_internal";

export const NarratorSchema = z
    .object({
        id: uuidSchema,
        name: userNameSchema,
        bio: z.optional(userBioSchema),
    })
    .and(passwordSchema);

export type Narrator = z.infer<typeof NarratorSchema>;

export const CreateNarratorRequestSchema = z.object({
    name: userNameSchema,
    bio: z.optional(userBioSchema),
});

export type CreateNarratorRequest = z.infer<typeof CreateNarratorRequestSchema>;

export const UpdateNarratorRequestSchema = z.object({
    name: z.optional(userNameSchema),
    bio: z.optional(userBioSchema),
});

export type UpdateNarratorRequest = z.infer<typeof UpdateNarratorRequestSchema>;
