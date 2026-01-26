import api from "./index";
import type { CreateGenreRequest, Genre } from "../models/genre";

export async function createGenre(genreData: CreateGenreRequest): Promise<Genre> {
    const res = await api.post<Genre>("/genres", genreData);
    return res.data;
}

export async function listGenres(): Promise<Genre[]> {
    const res = await api.get<Genre[]>("/genres");
    return res.data;
}

export async function getGenre(id: string): Promise<Genre> {
    const res = await api.get<Genre>(`/genres/${id}`);
    return res.data;
}

export async function deleteGenre(id: string): Promise<void> {
    await api.delete(`/genres/${id}`);
}
