import MyForm from "@/lib/authorForm/new";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/authors")({
    component: RouteComponent,
});

function RouteComponent() {
    return <MyForm />;
}
