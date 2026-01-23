export const bookKeys = {
	all: ["books"] as const,
	lists: () => [...bookKeys.all, "list"] as const,
	details: () => [...bookKeys.all, "detail"] as const,
	detail: (id: string) => [...bookKeys.details(), id] as const,
};

export const authorKeys = {
	all: ["author"] as const,
	lists: () => [...authorKeys.all, "list"] as const,
	details: () => [...authorKeys.all, "detail"] as const,
	detail: (id: string) => [...authorKeys.details(), id] as const,
};

export const narratorKeys = {
	all: ["narrator"] as const,
	lists: () => [...narratorKeys.all, "list"] as const,
	details: () => [...narratorKeys.all, "detail"] as const,
	detail: (id: string) => [...narratorKeys.details(), id] as const,
};

export const genreKeys = {
	all: ["genre"] as const,
	lists: () => [...genreKeys.all, "list"] as const,
	details: () => [...genreKeys.all, "detail"] as const,
	detail: (id: string) => [...genreKeys.details(), id] as const,
};

export const chapterKeys = {
	all: ["chapter"] as const,
	lists: () => [...chapterKeys.all, "list"] as const,
	details: () => [...chapterKeys.all, "detail"] as const,
	detail: (id: string) => [...chapterKeys.details(), id] as const,
};
