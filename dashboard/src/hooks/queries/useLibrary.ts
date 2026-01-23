import { useQuery } from "@tanstack/react-query";
import { getBookById, getBooks } from "./../../api/books";
import { bookKeys } from "./keys";

export function useBooks() {
	return useQuery({
		queryKey: bookKeys.all,
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
