import { useForm } from "@tanstack/react-form";

export const useCreateAuthorForm = () => {
    return useForm({
        defaultValues: {
            email: "",
            password: "",
        },
        onSubmit: async ({ value }) => {
            console.log("Submitting:", value);
        },
    });
};
