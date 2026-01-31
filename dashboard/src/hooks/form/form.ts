import TextField_LDI from "@/components/form/field_with_label_and_input";
import { fieldContext, formContext } from "./form-context";
import { createFormHook } from "@tanstack/react-form";

export const { useAppForm, withForm, withFieldGroup } = createFormHook({
    fieldContext: fieldContext,
    formContext: formContext,
    fieldComponents: {
        TextField_LDI: TextField_LDI,
    },
    formComponents: {},
});
