import { ApiUrl } from "../consts";
import type { Narrator } from "../models/models";
import type { CreateNarratorRequest } from "../models/requests";
import { request } from "./_utils";

export const createNarrator = async (
	genreData: CreateNarratorRequest,
): Promise<Narrator> => {
	const url = ApiUrl + "/narrators";
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

export const getNarrators = async (): Promise<Narrator[]> => {
	return request<Narrator[]>("/narrators");
};

export const getNarratorById = async (id: string): Promise<Narrator> => {
	return request<Narrator>(`/narrators/${id}`);
};
