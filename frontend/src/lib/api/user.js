import { API_URL, handleResponse, getAuthHeaders } from './config';

export const userApi = {
	// Получение профиля пользователя
	getProfile: async () => {
		const response = await fetch(`${API_URL}/users/me`, {
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	},

	// Обновление профиля пользователя
	updateProfile: async (data) => {
		const response = await fetch(`${API_URL}/users/me`, {
			method: 'PUT',
			headers: getAuthHeaders(),
			body: JSON.stringify(data)
		});
		return handleResponse(response);
	},

	// Обновление пароля
	updatePassword: async (currentPassword, newPassword) => {
		const response = await fetch(`${API_URL}/users/me/password`, {
			method: 'PUT',
			headers: getAuthHeaders(),
			body: JSON.stringify({
				current_password: currentPassword,
				new_password: newPassword
			})
		});
		return handleResponse(response);
	},

	// Удаление аккаунта
	deleteAccount: async () => {
		const response = await fetch(`${API_URL}/users/me`, {
			method: 'DELETE',
			headers: getAuthHeaders()
		});
		return handleResponse(response);
	}
};
