import api from "./index";
import type { Author, CreateAuthorRequest, UpdateAuthorRequest } from "@/schemas/author";

export async function createAuthor(req: CreateAuthorRequest): Promise<Author> {
    const res = await api.post<Author>("/authors", req);
    return res.data as Author;
}

export async function listAuthors(): Promise<Author[]> {
    const res = await api.get<Author[]>("/authors");
    return res.data;
}

export async function getAuthor(id: string): Promise<Author> {
    const res = await api.get<Author>(`/authors/${id}`);
    return res.data;
}

export async function updateAuthor(id: string, req: UpdateAuthorRequest): Promise<Author> {
    const res = await api.put<Author>(`/authors/${id}`, req);
    return res.data;
}

export async function deleteAuthor(id: string): Promise<void> {
    await api.delete(`/authors/${id}`);
}
