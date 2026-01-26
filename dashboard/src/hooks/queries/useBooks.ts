import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
    addAuthorToBook,
    addGenreToBook,
    addNarratorToBook,
    createBook,
    deleteBook,
    getBook,
    listBooks,
    removeAuthorFromBook,
    removeGenreFromBook,
    removeNarratorFromBook,
    updateBook,
} from "@/api/books";
import { bookKeys } from "./keys";
import type {
    AddAuthorToBookRequest,
    AddGenreToBookRequest,
    AddNarratorToBookRequest,
    UpdateBookRequest,
} from "@/models/book";

export function useCreateBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: createBook,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}

export function useListBooks() {
    return useQuery({
        queryKey: bookKeys.lists(),
        queryFn: listBooks,
        staleTime: 1000 * 60 * 5,
    });
}

export function useGetBook(id: string) {
    return useQuery({
        queryKey: bookKeys.detail(id),
        queryFn: () => getBook(id),
        enabled: !!id,
    });
}

export function useUpdateBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: UpdateBookRequest }) => updateBook(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}

export function useDeleteBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => deleteBook(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}

export function useAddAuthorToBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: AddAuthorToBookRequest }) => addAuthorToBook(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}

export function useAddNarratorToBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: AddNarratorToBookRequest }) => addNarratorToBook(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}
export function useAddGenreToBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, req }: { id: string; req: AddGenreToBookRequest }) => addGenreToBook(id, req),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}

export function useRemoveAuthorFromBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ bookId, authorId }: { bookId: string; authorId: string }) =>
            removeAuthorFromBook(bookId, authorId),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}

export function useRemoveNarratorFromBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ bookId, narratorId }: { bookId: string; narratorId: string }) =>
            removeNarratorFromBook(bookId, narratorId),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}
export function useRemoveGenreFromBook() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ bookId, genreId }: { bookId: string; genreId: string }) => removeGenreFromBook(bookId, genreId),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: bookKeys.lists() });
        },
    });
}
