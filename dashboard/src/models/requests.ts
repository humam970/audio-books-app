export type CreateAuthorRequest = {
	name: string;
	bio: string;
};

export type UpdateAuthorRequest = Partial<CreateAuthorRequest>;

export type AddAuthorToBookRequest = {
	author_id: string;
};

export type RemoveAuthorFromBookRequest = {
	author_id: string;
};

export type CreateNarratorRequest = {
	name: string;
	Bio?: string | null;
};

export type UpdateNarratorRequest = Partial<CreateNarratorRequest>;

export type AddNarratorToBookRequest = {
	narrator_id: string;
};

export type RemoveNarratorFromBookRequest = {
	narrator_id: string;
};

export type CreateGenreRequest = {
	name: string;
};

export type AddGenreToBookRequest = {
	genre_id: string;
};

export type RemoveGenreFromBookRequest = {
	genre_id: string;
};

export type CreateChapterRequest = {
	title: string;
	start_time: number;
	end_time: number;
	order_index: number;
};

export type UpdateChapterRequest = Partial<CreateChapterRequest>;

export type CreateBookRequest = {
	title: string;
	duration_seconds: number;
	rating: number;
	release_date: Date;
	cover_image_url: string;
	audio_preview_url: string;
	is_abridged: boolean;
};

export type UpdateBookRequest = {
	title: string;
	rating: number;
	is_abridged: boolean;
};
