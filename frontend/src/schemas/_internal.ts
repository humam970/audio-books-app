import z from "zod";

export const uuidSchema = z.uuid();
export const userNameSchema = z.string().min(4, "Name must be more than 4").max(9);
export const userBioSchema = z.string().min(4).max(9);

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
