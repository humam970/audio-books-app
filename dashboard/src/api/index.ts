import { ApiUrl } from "../consts";
import axios, { AxiosError, HttpStatusCode } from "axios";

const api = axios.create({
	baseURL: ApiUrl,
	withCredentials: true,
	headers: {
		"Content-Type": "application/json",
	},
});
export default api;

api.interceptors.response.use(undefined, (error: AxiosError) => {
	if (error.response) {
		const status = error.response.status;
		const data = error.response.data as any;

		switch (status) {
			case HttpStatusCode.Unauthorized:
				console.error("Token expired. Redirecting to login...");
				// Optional: clear localStorage and redirect
				break;
			case HttpStatusCode.Forbidden:
				console.error("Access denied. You don't have permission.");
				break;
			case HttpStatusCode.NotFound:
				console.error("Resource not found.");
				break;
			case HttpStatusCode.UnprocessableEntity:
				console.error("Validation failed:", data.errors);
				break;
			case HttpStatusCode.InternalServerError:
				console.error("Server crashed. Please try again later.");
				break;
		}
	} else if (error.request) {
		// The request was made but no response was received (Network Error)
		console.error("Network error: Please check your internet connection.");
	} else {
		// Something happened in setting up the request
		console.error("Axios Setup Error:", error.message);
	}

	return Promise.reject(error);
});
