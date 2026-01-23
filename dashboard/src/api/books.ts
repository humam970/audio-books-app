import { ApiUrl } from "../consts";
import type { CreateBookRequest } from "../models/requests";
import type { Book } from "./../models/models";
import { request } from "./_utils";

export const createBook = async (
	bookData: CreateBookRequest,
): Promise<Book> => {
	const url = `${ApiUrl}/books`;
	const response = await fetch(url, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(bookData),
	});

	if (!response.ok) {
		const errorData = await response.json().catch(() => ({}));
		throw new Error(
			errorData.message || `Error ${response.status}: Failed to create book`,
		);
	}

	return await response.json();
};

export const getBooks = async (): Promise<Book[]> => {
	return request<Book[]>("/books");
};

export const getBookById = async (id: string): Promise<Book> => {
	return request<Book>(`/books/${id}`);
};
