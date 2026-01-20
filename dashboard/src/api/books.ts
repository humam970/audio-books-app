import type { LibraryData } from "./../models/books";

export const fetchLibrary = async (): Promise<LibraryData> => {
	try {
		const response = await fetch("/data/books.json");

		if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);

		const data: LibraryData = await response.json();
		return data;
	} catch (error) {
		console.error("Error fetching library data:", error);
		if (error instanceof Error) {
			throw new Error(error.message);
		}
		throw new Error("An unexpected error occurred while fetching books.");
	}
};
