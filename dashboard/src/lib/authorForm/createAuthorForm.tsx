import { Button } from "@/components/ui/button";
import { FieldGroup } from "@/components/ui/field";
import { useAppForm } from "@/hooks/form/form";
import { createAuthorRequestDefaults, CreateAuthorRequestSchema } from "@/schemas/author";
import { useStore } from "@tanstack/react-form";

function CreateAuthorForm() {
    const form = useAppForm({
        defaultValues: createAuthorRequestDefaults,
        defaultState: {
            canSubmit: false,
        },
        validators: {
            onChange: CreateAuthorRequestSchema,
            onSubmit: ({ formApi, value }) => {
                console.log(formApi.fieldMetaDerived);
                console.log(value);
            },
        },
    });

    const store = useStore(form.store, (state) => state);

    return (
        <form
            className=" w-200"
            onSubmit={(e) => {
                e.preventDefault();
                form.handleSubmit();
            }}
        >
            <pre>{JSON.stringify(store, null, 4)}</pre>
            <FieldGroup>
                <form.AppField name="name">
                    {(field) => (
                        <field.TextField_LDI
                            id="name"
                            label="Author Name"
                            placeholder="Author Name"
                            description="Enter The Author Name"
                        />
                    )}
                </form.AppField>

                <form.AppField name="bio" validators={{}}>
                    {(field) => (
                        <field.TextField_LDI
                            id="bio"
                            label="Author Bio"
                            placeholder="Author Bio"
                            description="Enter The Author Bio"
                        />
                    )}
                </form.AppField>

                <Button type="submit" variant="secondary" disabled={form.state.canSubmit}>
                    Submit
                </Button>
            </FieldGroup>
        </form>
    );
}

export default CreateAuthorForm;
