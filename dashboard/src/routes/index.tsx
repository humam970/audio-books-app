import { ModeToggle } from "@/components/mode-toggle";
import { Button } from "@/components/ui/button";
import { Field, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { useAppForm } from "@/hooks/form/form";
import { userNameSchema } from "@/schemas/_internal";
import { createAuthorRequestDefaults } from "@/schemas/author";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
    component: MyApp,
});

function MyApp() {
    const form = useAppForm({
        defaultValues: createAuthorRequestDefaults,
        onSubmit: (data) => {
            console.log(data);
        },
    });

    return (
        <div>
            <ModeToggle />
            <form.AppForm>
                <form
                    className="*:mb-5 mt-5"
                    onSubmit={(e) => {
                        e.preventDefault();
                        form.handleSubmit();
                    }}
                >
                    <form.Field
                        name="name"
                        validators={{
                            onChange: userNameSchema,
                        }}
                    >
                        {(field) => {
                            return (
                                <div>
                                    <pre>{JSON.stringify(field.state, null, 4)}</pre>
                                    <Field>
                                        <FieldLabel htmlFor={field.name}>Author Name</FieldLabel>
                                        <Input
                                            id={field.name}
                                            onChange={(e) => field.handleChange(e.target.value)}
                                            onBlur={() => field.handleBlur()}
                                        />
                                    </Field>
                                </div>
                            );
                        }}
                    </form.Field>

                    <form.Subscribe selector={(state) => state.isDirty}>
                        {(isDirty) => <Button disabled={!isDirty}>Submit</Button>}
                    </form.Subscribe>
                </form>
            </form.AppForm>
        </div>
    );
}
