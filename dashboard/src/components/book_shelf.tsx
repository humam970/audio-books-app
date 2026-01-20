import BookCarousel from "./books_carousel";

const books = ["book", "book", "book", "book", "book", "book"];

function BooksShelf({ genre }: { genre: string }) {
	return (
		<div>
			<div className="flex justify-between mx-4">
				<span className="text-lg font-medium">{genre}</span>
				<a href="#" className="text-sm text-[#E36166] font-normal">
					See all
				</a>
			</div>

			<BookCarousel books={books} />
		</div>
	);
}

export default BooksShelf;
