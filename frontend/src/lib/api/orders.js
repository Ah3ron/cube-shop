import { API_URL, handleResponse, getAuthHeaders } from './config';

export const ordersApi = {
	checkout: async (shippingDetails) => {
		const response = await fetch(`${API_URL}/orders/checkout`, {
			method: 'POST',
			headers: getAuthHeaders(),
			body: JSON.stringify({ shipping_details: shippingDetails })
		});
		return handleResponse(response);
	},

	getAll: async () => {
		const response = await fetch(`${API_URL}/orders`, {
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	},

	getById: async (id) => {
		const response = await fetch(`${API_URL}/orders/${id}`, {
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	},

	updateStatus: async (id, status) => {
		const response = await fetch(`${API_URL}/orders/${id}/status`, {
			method: 'PUT',
			headers: getAuthHeaders(),
			body: JSON.stringify({ status })
		});
		return handleResponse(response);
	}
};
