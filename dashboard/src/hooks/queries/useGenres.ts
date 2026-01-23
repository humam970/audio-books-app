import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { genreKeys } from "./keys";
import { createGenre, getGenreById, getGenres } from "../../api/genres";

export function useCreateGenre() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: createGenre,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: genreKeys.lists() });
		},
	});
}

export function useGenres() {
	return useQuery({
		queryKey: genreKeys.lists(),
		queryFn: getGenres,
		staleTime: 1000 * 60 * 5,
	});
}

export function useGenre(id: string) {
	return useQuery({
		queryKey: genreKeys.detail(id),
		queryFn: () => getGenreById(id),
		enabled: !!id,
	});
}
