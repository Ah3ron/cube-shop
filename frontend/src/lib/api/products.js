import { API_URL, getAuthHeaders, handleResponse } from './config';

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
	},

	delete: async (id) => {
		const response = await fetch(`${API_URL}/products/${id}`, {
			method: 'DELETE',
			headers: getAuthHeaders()
		});

		if (!response.ok) {
			const error = await response.json();
			throw new Error(error.error || 'Failed to delete product');
		}

		return response.status === 204 ? true : await response.json();
	}
};
