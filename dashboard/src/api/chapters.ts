import api from "./index";
import type { Chapter, CreateChapterRequest, UpdateChapterRequest } from "@/schemas/chapter";

export async function createChapter(genreData: CreateChapterRequest): Promise<Chapter> {
    const res = await api.post<Chapter>("/chapters", genreData);
    return res.data;
}

export async function listChapters(): Promise<Chapter[]> {
    const res = await api.get<Chapter[]>("/chapters");
    return res.data;
}

export async function getChapter(id: string): Promise<Chapter> {
    const res = await api.get<Chapter>(`/chapters/${id}`);
    return res.data;
}

export async function updateChapter(id: string, updateData: UpdateChapterRequest): Promise<Chapter> {
    const res = await api.put<Chapter>(`/chapters/${id}`, updateData);
    return res.data;
}

export async function deleteChapter(id: string): Promise<void> {
    await api.delete(`/chapters/${id}`);
}
