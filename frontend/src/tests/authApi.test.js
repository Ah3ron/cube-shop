import { describe, it, expect, vi, beforeEach } from 'vitest';
import { authApi } from '../lib/api/auth'; // Путь к вашему файлу authApi
import { API_URL, getAuthHeaders } from '../lib/api/config';

describe('authApi', () => {
	beforeEach(() => {
		// Сброс всех моков перед каждым тестом
		vi.clearAllMocks();

		global.localStorage = {
			getItem: vi.fn().mockReturnValue('mocked_token'),
			setItem: vi.fn(),
			removeItem: vi.fn(),
			clear: vi.fn()
		};
	});

	it('should register a user', async () => {
		const mockResponse = { success: true };
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await authApi.register('test@example.com', 'password', 'Test User');

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/auth/register`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ email: 'test@example.com', password: 'password', name: 'Test User' })
		});
		expect(response).toEqual(mockResponse);
	});

	it('should login a user', async () => {
		const mockResponse = { success: true };
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await authApi.login('test@example.com', 'password');

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/auth/login`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ email: 'test@example.com', password: 'password' })
		});
		expect(response).toEqual(mockResponse);
	});

	it('should logout a user', async () => {
		const mockResponse = { success: true };
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await authApi.logout();

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/auth/logout`, {
			method: 'POST',
			headers: getAuthHeaders()
		});
		expect(response).toEqual(mockResponse);
	});
});
