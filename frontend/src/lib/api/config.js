export const API_URL = '/api/v1';

export const handleResponse = async (response) => {
	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Что-то пошло не так');
	}
	if (response.status === 204) {
		return null;
	}
	return response.json();
};

export const getAuthHeaders = () => {
	const token = localStorage.getItem('token');
	return {
		'Content-Type': 'application/json',
		...(token && { Authorization: `Bearer ${token}` })
	};
};
