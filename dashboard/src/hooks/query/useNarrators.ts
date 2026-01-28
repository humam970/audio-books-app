import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { authorKeys } from "./keys";
import { createNarrator, deleteNarrator, getNarrator, listNarrators, updateNarrator } from "../../api/narrators";
import type { UpdateNarratorRequest } from "@/models/narrator";

export function useCreateNarrator() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: createNarrator,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
        },
    });
}

export function useListNarrators() {
    return useQuery({
        queryKey: authorKeys.lists(),
        queryFn: listNarrators,
        staleTime: 1000 * 60 * 5,
    });
}

export function useGetNarrator(id: string) {
    return useQuery({
        queryKey: authorKeys.detail(id),
        queryFn: () => getNarrator(id),
        enabled: !!id,
    });
}

export function useUpdateNarrator() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: UpdateNarratorRequest }) => updateNarrator(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
        },
    });
}

export function useDeleteNarrator() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => deleteNarrator(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: authorKeys.lists() });
        },
    });
}
