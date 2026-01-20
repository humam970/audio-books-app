export interface Author {
	id: number;
	name: string;
	bio: string;
}

export interface Chapter {
	id: number;
	book_id: number;
	title: string;
	start_time: number;
	end_time: number;
}

export interface Book {
	id: number;
	title: string;
	author_id: number;
	author_name: string; // From the JOIN in your SQL query
	narrator: string;
	duration_hours: number;
	duration_minutes: number;
	rating: number;
	release_date: string; // ISO Date string from Postgres
	cover_image_url: string;
	audio_preview_url: string;
	is_abridged: boolean;

	genres: string[];
	chapters?: Chapter[];
}

export interface Genre {
	id: number;
	name: string;
}

export type LibraryResponse = Book[];
