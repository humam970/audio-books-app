import * as z from "zod";

export type Author = {
    id: string;
    name: string;
    bio: string;
};

export const idSchema = z.uuid();
export const usernameSchema = z.string().min(4).max(9);
export const bioSchema = z.string().min(4).max(9);

const internalPasswordSchema = z.string().min(8).max(20);

export const passwordSchema = z
    .object({
        password: internalPasswordSchema,
        confirmPassword: internalPasswordSchema,
    })
    .refine((data) => data.password === data.confirmPassword, {
        message: "Passwords don't match",
        path: ["confirmPassword"],
    });

export const AuthorSchema = z
    .object({
        id: idSchema,
        name: usernameSchema,
        bio: bioSchema,
    })
    .and(passwordSchema);

export type CreateAuthorRequest = {
    name: string;
    bio: string;
};

export const CreateAuthorRequestDefaults: CreateAuthorRequest = {
    name: "",
    bio: "",
};

export type UpdateAuthorRequest = Partial<CreateAuthorRequest>;
