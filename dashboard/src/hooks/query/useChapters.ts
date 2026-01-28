import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { chapterKeys } from "./keys";
import { createChapter, deleteChapter, getChapter, listChapters, updateChapter } from "../../api/chapters";
import type { UpdateChapterRequest } from "@/models/chapter";

export function useCreateChapter() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: createChapter,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: chapterKeys.lists() });
        },
    });
}

export function useListChapters() {
    return useQuery({
        queryKey: chapterKeys.lists(),
        queryFn: listChapters,
        staleTime: 1000 * 60 * 5,
    });
}

export function useGetChapter(id: string) {
    return useQuery({
        queryKey: chapterKeys.detail(id),
        queryFn: () => getChapter(id),
        enabled: !!id,
    });
}

export function useUpdateChapter() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: UpdateChapterRequest }) => updateChapter(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: chapterKeys.lists() });
        },
    });
}

export function useDeleteChapter() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => deleteChapter(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: chapterKeys.lists() });
        },
    });
}
