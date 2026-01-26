import type { CreateNarratorRequest, Narrator, UpdateNarratorRequest } from "../models/narrator";
import api from "./index";

export async function createNarrator(genreData: CreateNarratorRequest): Promise<Narrator> {
    const res = await api.post<Narrator>("/narrators", genreData);
    return res.data;
}

export async function listNarrators(): Promise<Narrator[]> {
    const res = await api.get<Narrator[]>("/narrators");
    return res.data;
}

export async function getNarrator(id: string): Promise<Narrator> {
    const res = await api.get<Narrator>(`/narrators/${id}`);
    return res.data;
}

export async function updateNarrator(id: string, updateData: UpdateNarratorRequest): Promise<Narrator> {
    const res = await api.put<Narrator>(`/narrators/${id}`, updateData);
    return res.data;
}

export async function deleteNarrator(id: string): Promise<void> {
    await api.delete(`/narrators/${id}`);
}
