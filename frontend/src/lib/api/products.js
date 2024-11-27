import { API_URL, handleResponse, getAuthHeaders } from './config';

export const productsApi = {
	getAll: async (filters = {}) => {
		const params = new URLSearchParams(filters);
		const response = await fetch(`${API_URL}/products?${params}`);
		return handleResponse(response);
	},

	getById: async (id) => {
		const response = await fetch(`${API_URL}/products/${id}`);
		return handleResponse(response);
	},

	search: async (query) => {
		const response = await fetch(`${API_URL}/products/search?q=${query}`);
		return handleResponse(response);
	}
};
