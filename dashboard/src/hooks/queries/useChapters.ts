import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { chapterKeys } from "./keys";
import { createChapter, getChapterById, getChapters } from "../../api/chapters";

export function useCreateChapter() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: createChapter,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: chapterKeys.lists() });
		},
	});
}

export function useChapters() {
	return useQuery({
		queryKey: chapterKeys.lists(),
		queryFn: getChapters,
		staleTime: 1000 * 60 * 5,
	});
}

export function useChapter(id: string) {
	return useQuery({
		queryKey: chapterKeys.detail(id),
		queryFn: () => getChapterById(id),
		enabled: !!id,
	});
}
