import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { narratorKeys } from "./keys";
import {
	createNarrator,
	getNarratorById,
	getNarrators,
} from "../../api/narrators";

export function useCreateNarrator() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: createNarrator,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: narratorKeys.lists() });
		},
	});
}

export function useNarrators() {
	return useQuery({
		queryKey: narratorKeys.lists(),
		queryFn: getNarrators,
		staleTime: 1000 * 60 * 5,
	});
}

export function useNarrator(id: string) {
	return useQuery({
		queryKey: narratorKeys.detail(id),
		queryFn: () => getNarratorById(id),
		enabled: !!id,
	});
}
