import z from "zod";
import { userBioSchema, uuidSchema, userNameSchema } from "./_internal";

export const AuthorSchema = z.object({
    id: uuidSchema,
    name: userNameSchema,
    bio: userBioSchema,
});

export type Author = z.infer<typeof AuthorSchema>;

export const CreateAuthorRequestSchema = z.object({
    name: userNameSchema,
    bio: userBioSchema,
});

export type CreateAuthorRequest = z.infer<typeof CreateAuthorRequestSchema>;

export const createAuthorRequestDefaults: CreateAuthorRequest = {
    name: "",
    bio: "",
};

export const UpdateAuthorRequestSchema = CreateAuthorRequestSchema.partial();

export type UpdateAuthorRequest = z.infer<typeof UpdateAuthorRequestSchema>;

export const updateAuthorRequestDefaults: UpdateAuthorRequest = {
    name: "",
    bio: "",
};
