import BookCard from "./book_card";

function BookCarousel({ books }: { books: string[] }) {
	return (
		<div className="flex gap-2 overflow-x-auto [&>*:first-child]:ml-4">
			{books.map((book) => (
				<BookCard src="books/black_dots.png" title={book} author={book} />
			))}
		</div>
	);
}

export default BookCarousel;
