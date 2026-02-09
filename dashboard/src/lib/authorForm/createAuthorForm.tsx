import { Button } from "@/components/ui/button";
import { FieldGroup } from "@/components/ui/field";
import { useAppForm } from "@/hooks/form/form";
import { useCreateAuthor } from "@/hooks/query/useAuthors";
import { createAuthorRequestDefaults, CreateAuthorRequestSchema } from "@/schemas/author";

function CreateAuthorForm() {
    const { mutate, isPending, isError, error } = useCreateAuthor();

    const form = useAppForm({
        defaultValues: createAuthorRequestDefaults,
        validators: {
            onChange: CreateAuthorRequestSchema,
            onSubmit: ({ formApi, value }) => {
                mutate(value);
            },
        },
    });

    return (
        <form
            className=" w-200"
            onSubmit={(e) => {
                e.preventDefault();
                form.handleSubmit();
            }}
        >
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

                <form.AppField name="bio">
                    {(field) => (
                        <field.TextField_LDI
                            id="bio"
                            label="Author Bio"
                            placeholder="Author Bio"
                            description="Enter The Author Bio"
                        />
                    )}
                </form.AppField>

                <form.Subscribe selector={(state) => [state.canSubmit, state.isDirty]}>
                    {([canSubmit, isDirty]) => (
                        <Button type="submit" disabled={!canSubmit || !isDirty}>
                            Submit
                        </Button>
                    )}
                </form.Subscribe>
            </FieldGroup>
        </form>
    );
}

export default CreateAuthorForm;
