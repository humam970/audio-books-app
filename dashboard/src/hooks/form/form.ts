import { fieldContext, formContext } from "./form-context";
import { createFormHook } from "@tanstack/react-form";
import LabelAndInput from "@/components/form/label_and_input";

export const { useAppForm, withForm, withFieldGroup } = createFormHook({
    fieldContext: fieldContext,
    formContext: formContext,
    fieldComponents: { LabelAndInput },
    formComponents: {},
});
