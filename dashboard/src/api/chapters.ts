import { ApiUrl } from "../consts";
import type { Chapter } from "../models/models";
import type { CreateChapterRequest } from "../models/requests";
import { request } from "./_utils";

export const createChapter = async (
	genreData: CreateChapterRequest,
): Promise<Chapter> => {
	const url = ApiUrl + "/chapters";
	const response = await fetch(url, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(genreData),
	});

	if (!response.ok) {
		const errorData = await response.json().catch(() => ({}));
		throw new Error(
			errorData.message || `Error ${response.status}: Failed to create book`,
		);
	}

	return await response.json();
};

export const getChapters = async (): Promise<Chapter[]> => {
	return request<Chapter[]>("/chapters");
};

export const getChapterById = async (id: string): Promise<Chapter> => {
	return request<Chapter>(`/chapters/${id}`);
};
