import { Field, FieldDescription, FieldError, FieldLabel } from "@/components/ui/field";
import { Input } from "../ui/input";
import { useFieldContext } from "@/hooks/form/form-context";
import type { HTMLInputTypeAttribute } from "react";

export default function LabelAndInput({
    id,
    label,
    description,
    placeholder,
    type = "text",
}: {
    id: string;
    label: string;
    description?: string;
    placeholder?: string;
    type?: HTMLInputTypeAttribute;
}) {
    const field = useFieldContext<string>();
    const shouldShowError = field.state.meta.isTouched || field.state.meta.isValid;

    return (
        <Field>
            <FieldLabel htmlFor={id}>{label}</FieldLabel>

            <Input
                id={id}
                type={type}
                placeholder={placeholder}
                value={field.state.value ?? ""}
                onBlur={field.handleBlur}
                onChange={(e) => field.handleChange(e.target.value)}
            />

            {description && <FieldDescription>{description}</FieldDescription>}

            {shouldShowError && (
                <FieldError>
                    {field.state.meta.errors.map((error: any) => error?.message ?? error).join(", ")}
                </FieldError>
            )}
        </Field>
    );
}
