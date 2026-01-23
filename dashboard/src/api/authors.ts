import { ApiUrl } from "../consts";
import type { CreateAuthorRequest } from "../models/requests";
import type { Author } from "./../models/models";
import { request } from "./_utils";

export const createAuthor = async (
	authorData: CreateAuthorRequest,
): Promise<Author> => {
	const url = ApiUrl + "/authors";
	const response = await fetch(url, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(authorData),
	});

	if (!response.ok) {
		const errorData = await response.json().catch(() => ({}));
		throw new Error(
			errorData.message || `Error ${response.status}: Failed to create book`,
		);
	}

	return await response.json();
};

export const getAuthors = async (): Promise<Author[]> => {
	return request<Author[]>("/authors");
};

export const getAuthorById = async (id: string): Promise<Author> => {
	return request<Author>(`/authors/${id}`);
};
