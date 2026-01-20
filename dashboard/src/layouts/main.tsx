import BooksShelf from "../components/book_shelf";
import GenreCarousel from "../components/genre_carousel";

function Main() {
	return (
		<main className="col-[full] *:mb-6">
			<GenreCarousel />
			<BooksShelf genre="Drama" />
		</main>
	);
}

export default Main;
