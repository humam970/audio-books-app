export const bookKeys = {
	all: ["books"] as const,
	detail: (id: string) => [...bookKeys.all, "detail", id] as const,
	search: (query: string) => [...bookKeys.all, "search", query] as const,
};
