import { useQuery, useQueryClient } from "@tanstack/react-query";
import { fetchLibrary } from "./../../api/books";
import { bookKeys } from "./keys";
import type { LibraryData } from "../../models/books";

export function useLibrary() {
	return useQuery({
		queryKey: bookKeys.all,
		queryFn: fetchLibrary,
		staleTime: 1000 * 60 * 5,
	});
}

export function useBook(id: string) {
	const queryClient = useQueryClient();

	return useQuery({
		queryKey: bookKeys.detail(id),
		queryFn: async () => {
			const data = await fetchLibrary();
			const found = data.books.find((b) => b.id === id);
			if (!found) throw new Error(`Book with ID ${id} not found`);
			return found;
		},
		enabled: !!id,
		// Step 1: Look in the 'all books' cache before fetching
		initialData: () => {
			const allData = queryClient.getQueryData<LibraryData>(bookKeys.all);
			return allData?.books.find((b) => b.id === id);
		},
		// Step 2: If we found it in initialData, consider it fresh for 5 mins
		initialDataUpdatedAt: () =>
			queryClient.getQueryState(bookKeys.all)?.dataUpdatedAt,
	});
}
