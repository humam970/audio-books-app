import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { createBook, getBookById, getBooks } from "./../../api/books";
import { bookKeys } from "./keys";

export function useCreateBook() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: createBook,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
		},
	});
}

export function useBooks() {
	return useQuery({
		queryKey: bookKeys.lists(),
		queryFn: getBooks,
		staleTime: 1000 * 60 * 5,
	});
}

export function useBook(id: string) {
	return useQuery({
		queryKey: bookKeys.detail(id),
		queryFn: () => getBookById(id),
		enabled: !!id,
	});
}
