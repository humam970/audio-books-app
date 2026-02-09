import { createFileRoute } from "@tanstack/react-router";
import GenreSelector from "./-components/GenreSelect";
import FeaturedBook from "./-components/FeaturedBook";

export const Route = createFileRoute("/")({
    component: RouteComponent,
});

function RouteComponent() {
    return (
        <div className="mt-section-separator *:mb-section-separator px-gutter">
            <FeaturedBook />
            <GenreSelector />
        </div>
    );
}
