import { API_URL, handleResponse, getAuthHeaders } from './config';

export const productsApi = {
    getAll: async (filters = {}) => {
        const queryString = new URLSearchParams(filters).toString();
        const url = queryString ? `${API_URL}/products?${queryString}` : `${API_URL}/products`;
        const response = await fetch(url);
        return handleResponse(response);
    },

    getById: async (id) => {
        const response = await fetch(`${API_URL}/products/${id}`);
        return handleResponse(response);
    },

    create: async (productData) => {
        const response = await fetch(`${API_URL}/products`, {
            method: 'POST',
            headers: getAuthHeaders(),
            body: JSON.stringify(productData)
        });
        return handleResponse(response);
    },

    update: async (id, productData) => {
        const response = await fetch(`${API_URL}/products/${id}`, {
            method: 'PUT',
            headers: getAuthHeaders(),
            body: JSON.stringify(productData)
        });
        return handleResponse(response);
    },

    delete: async (id) => {
        const response = await fetch(`${API_URL}/products/${id}`, {
            method: 'DELETE',
            headers: getAuthHeaders()
        });
        return handleResponse(response);
    },

    search: async (query) => {
        const response = await fetch(`${API_URL}/products/search?q=${query}`);
        return handleResponse(response);
    }
};