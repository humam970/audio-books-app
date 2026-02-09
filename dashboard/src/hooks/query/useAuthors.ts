import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { authorKeys } from "./keys";
import { createAuthor, deleteAuthor, getAuthor, listAuthors, updateAuthor } from "@/api/authors";
import type { UpdateAuthorRequest } from "@/schemas/author";

export function useCreateAuthor() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: createAuthor,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
        },
    });
}

export function useListAuthors() {
    return useQuery({
        queryKey: authorKeys.lists(),
        queryFn: listAuthors,
        staleTime: 1000 * 60 * 5,
    });
}

export function useGetAuthor(id: string) {
    return useQuery({
        queryKey: authorKeys.detail(id),
        queryFn: () => getAuthor(id),
        enabled: !!id,
    });
}

export function useUpdateAuthor() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: UpdateAuthorRequest }) => updateAuthor(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
        },
    });
}

export function useDeleteAuthor() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => deleteAuthor(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
        },
    });
}
