import { Star } from "lucide-react";

type Book = {
    title: string;
    genre: string;
    author: string;
};

const books: Book[] = [
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
    { title: "The Green Garden", author: "Ali Abbas", genre: "Fantasy" },
];

type BookCardProps = {
    title: string;
    genre: string;
    authorName: string;
};

function BookCard({ title, authorName, genre }: BookCardProps) {
    return (
        <figure
            className="aspect-3/4
                w-[calc((100cqi-var(--spacing)*6*1)/2)]
             sm:w-[calc((100cqi-var(--spacing)*6*2)/3)]
             lg:w-[calc((100cqi-var(--spacing)*6*3)/4)]
             xl:w-[calc((100cqi-var(--spacing)*6*4)/5)]
            2xl:w-[calc((100cqi-var(--spacing)*6*5)/6)]
            "
        >
            <img src="/featured_book.png" className="w-full object-fill rounded-lg mb-2xs" />
            <figcaption className="flex flex-col text-gray-600">
                <cite className="not-italic text-black text-base font-bold">{title}</cite>
                <span className="text-sm mb-3xs">by {authorName}</span>
                <span className="flex justify-between text-xs">
                    {genre}
                    <span className="flex items-center gap-1">
                        <Star size={12} className="fill-yellow-600 stroke-yellow-600" />
                        4.7
                    </span>
                </span>
            </figcaption>
        </figure>
    );
}

function BooksGrid() {
    return (
        <section className="@container flex flex-wrap gap-6">
            {books.map((book) => (
                <BookCard title={book.title} genre={book.genre} authorName={book.author} />
            ))}
        </section>
    );
}

export default BooksGrid;
