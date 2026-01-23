import { useCreateAuthor } from "../hooks/queries/useAuthors";

export function AddAuthorForm() {
	const { mutate, isPending, isError, error } = useCreateAuthor();

	const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();

		const formData = new FormData(e.currentTarget);
		const name = formData.get("name") as string;
		const bio = formData.get("bio") as string;

		// Trigger the mutation
		mutate(
			{ bio, name },
			{
				onSuccess: () => {
					(e.target as HTMLFormElement).reset();
					alert("Author added!");
				},
			},
		);
	};

	return (
		<form onSubmit={handleSubmit}>
			<input name="name" placeholder="Author Name" required />
			<input name="bio" placeholder="Author bio" required />

			<button
				className="p-4 bg-amber-400 text-nowrap"
				type="submit"
				disabled={isPending}
			>
				{isPending ? "Adding..." : "Add Author"}
			</button>

			{isError && <p>Error: {error.message}</p>}
		</form>
	);
}

function Main() {
	return (
		<main className="bg-[#ccc] grid grid-cols-9">
			<AddAuthorForm />
		</main>
	);
}

export default Main;
