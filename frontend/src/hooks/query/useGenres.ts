import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { genreKeys } from "./keys";
import { createGenre, deleteGenre, getGenre, listGenres } from "@/api/genres";

export function useCreateGenre() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: createGenre,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: genreKeys.lists() });
        },
    });
}

export function useListGenres() {
    return useQuery({
        queryKey: genreKeys.lists(),
        queryFn: listGenres,
        staleTime: 1000 * 60 * 5,
    });
}

export function useGetGenre(id: string) {
    return useQuery({
        queryKey: genreKeys.detail(id),
        queryFn: () => getGenre(id),
        enabled: !!id,
    });
}

export function useDeleteGenre() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => deleteGenre(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: genreKeys.lists() });
        },
    });
}
