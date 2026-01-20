import GenreItem from "./genre_item";

const genres = ["All", "Detective", "Drama", "Historic"];

function GenreCarousel() {
	return (
		<div className="flex gap-2 overflow-x-auto [&>*:first-child]:ml-4">
			{genres.map((genre) => (
				<GenreItem itemName={genre} listName={genre} />
			))}
		</div>
	);
}

export default GenreCarousel;
