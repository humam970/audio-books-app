export type Author = {
	id: string;
	name: string;
	bio: string;
};

export type Book = {
	id: string;
	title: string;
	duration_seconds: number;
	rating: number;
	release_date: Date;
	cover_image_url: string;
	audio_preview_url: string;
	is_abridged: boolean;
	created_at: Date;
};

export type BookAuthor = {
	book_id: string;
	author_id: string;
};

export type BookGenre = {
	book_id: string;
	genre_id: string;
};

export type BookNarrator = {
	book_id: string;
	narrator_id: string;
};

export type Chapter = {
	id: string;
	book_id: string;
	title: string;
	start_time: number;
	end_time: number;
	order_index: number;
};

export type Genre = {
	id: string;
	name: string;
};

export type Narrator = {
	id: string;
	name: string;
	bio: string | null;
};
