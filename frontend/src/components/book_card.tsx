function BookCard({
	src,
	title,
	author,
}: {
	src: string;
	title: string;
	author: string;
}) {
	return (
		<div className="flex flex-col gap-3.5 min-w-40">
			<img src={src} />
			<div>
				<span className="font-medium text-lg">{title}</span>
				<br />
				<span className="text-sm">{author}</span>
			</div>
		</div>
	);
}

export default BookCard;
