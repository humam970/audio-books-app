import api from "./index";
import type {
    AddAuthorToBookRequest,
    AddAuthorToBookResponse,
    AddGenreToBookRequest,
    AddGenreToBookResponse,
    AddNarratorToBookRequest,
    AddNarratorToBookResponse,
    Book,
    CreateBookRequest,
    UpdateBookRequest,
} from "@/schemas/book";

export async function createBook(req: CreateBookRequest): Promise<Book> {
    const res = await api.post<Book>("/books", req);
    return res.data;
}

export async function listBooks(): Promise<Book[]> {
    const res = await api.get<Book[]>("/books");
    return res.data;
}

export async function getBook(id: string): Promise<Book> {
    const res = await api.get<Book>(`/books/${id}`);
    return res.data;
}

export async function updateBook(id: string, req: UpdateBookRequest): Promise<Book> {
    const res = await api.put<Book>(`/books/${id}`, req);
    return res.data;
}

export async function deleteBook(id: string): Promise<void> {
    await api.delete(`/books/${id}`);
}

export async function addAuthorToBook(id: string, req: AddAuthorToBookRequest): Promise<AddAuthorToBookResponse> {
    const res = await api.post<AddAuthorToBookResponse>(`/books/${id}/authors`, req);
    return res.data;
}

export async function addGenreToBook(id: string, req: AddGenreToBookRequest): Promise<AddGenreToBookResponse> {
    const res = await api.post<AddGenreToBookResponse>(`/books/${id}/genres`, req);
    return res.data;
}

export async function addNarratorToBook(id: string, req: AddNarratorToBookRequest): Promise<AddNarratorToBookResponse> {
    const res = await api.post<AddNarratorToBookResponse>(`/books/${id}/narrators`, req);
    return res.data;
}

export async function removeAuthorFromBook(bookId: string, authorId: string): Promise<void> {
    await api.delete(`/books/${bookId}/authors/${authorId}`);
}

export async function removeNarratorFromBook(bookId: string, narratorId: string): Promise<void> {
    await api.delete(`/books/${bookId}/narrators/${narratorId}`);
}

export async function removeGenreFromBook(bookId: string, genreId: string): Promise<void> {
    await api.delete(`/books/${bookId}/genres/${genreId}`);
}
