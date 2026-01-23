import { ApiUrl } from "../consts";
import type { Genre } from "../models/models";
import type { CreateGenreRequest } from "../models/requests";
import { request } from "./_utils";

export const createGenre = async (
	genreData: CreateGenreRequest,
): Promise<Genre> => {
	const url = ApiUrl + "/genres";
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

export const getGenres = async (): Promise<Genre[]> => {
	return request<Genre[]>("/genres");
};

export const getGenreById = async (id: string): Promise<Genre> => {
	return request<Genre>(`/genres/${id}`);
};
