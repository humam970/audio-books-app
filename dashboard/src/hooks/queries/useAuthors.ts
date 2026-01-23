import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { authorKeys } from "./keys";
import { createAuthor, getAuthorById, getAuthors } from "../../api/authors";

export function useCreateAuthor() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: createAuthor,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
		},
	});
}

export function useAuthors() {
	return useQuery({
		queryKey: authorKeys.lists(),
		queryFn: getAuthors,
		staleTime: 1000 * 60 * 5,
	});
}

export function useAuthor(id: string) {
	return useQuery({
		queryKey: authorKeys.detail(id),
		queryFn: () => getAuthorById(id),
		enabled: !!id,
	});
}
