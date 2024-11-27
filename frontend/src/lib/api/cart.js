import { API_URL, handleResponse, getAuthHeaders } from './config';

export const cartApi = {
	get: async () => {
		const response = await fetch(`${API_URL}/cart`, {
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	},

	addItem: async (productId, quantity) => {
		const response = await fetch(`${API_URL}/cart/items`, {
			method: 'POST',
			headers: getAuthHeaders(),
			body: JSON.stringify({ product_id: productId, quantity })
		});
		return handleResponse(response);
	},

	updateItem: async (id, quantity) => {
		const response = await fetch(`${API_URL}/cart/items/${id}`, {
			method: 'PUT',
			headers: getAuthHeaders(),
			body: JSON.stringify({ quantity })
		});
		return handleResponse(response);
	},

	removeItem: async (id) => {
		const response = await fetch(`${API_URL}/cart/items/${id}`, {
			method: 'DELETE',
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	},

	clear: async () => {
		const response = await fetch(`${API_URL}/cart`, {
			method: 'DELETE',
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	}
};
