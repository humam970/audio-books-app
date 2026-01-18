function GenreItem({
	itemName,
	listName,
}: {
	itemName: string;
	listName: string;
}) {
	return (
		<label
			htmlFor={itemName}
			className="px-6 py-3.5 bg-[#E7E0CB] text-black rounded-full
				has-checked:bg-black
				has-checked:text-white
			"
		>
			<input id={itemName} className="sr-only" type="radio" name={listName} />
			{itemName}
		</label>
	);
}

export default GenreItem;
