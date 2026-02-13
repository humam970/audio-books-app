import { Button } from "@/components/ui/button";

const genres = ["All", "Fantasy", "Mystery", "Si-Fi", "Romance", "Business", "Classic"];

function GenreSelector() {
    return (
        <div className="flex gap-3">
            {genres.map((genre) => (
                <GenreItem name={genre} />
            ))}
        </div>
    );
}

function GenreItem({ name }: { name: string }) {
    return (
        <Button asChild variant={name === "All" ? "default" : "outline"}>
            <a href="#">{name}</a>
        </Button>
    );
}

export default GenreSelector;
