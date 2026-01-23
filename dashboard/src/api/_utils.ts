import { ApiUrl } from "../consts";

export async function request<T>(
	endpoint: string,
	options?: RequestInit,
): Promise<T> {
	const url = ApiUrl + endpoint;
	try {
		const response = await fetch(url, options);

		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}

		return (await response.json()) as T;
	} catch (error) {
		console.error(`API Request Error [${endpoint}]:`, error);

		if (error instanceof Error) throw error;
		throw new Error("An unexpected error occurred.");
	}
}
