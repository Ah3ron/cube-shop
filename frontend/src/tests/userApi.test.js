import { describe, it, expect, vi, beforeEach } from 'vitest';
import { userApi } from '../lib/api/user'; // Путь к вашему файлу userApi
import { API_URL, getAuthHeaders } from '../lib/api/config';

describe('userApi', () => {
  beforeEach(() => {
    // Сброс всех моков перед каждым тестом
    vi.clearAllMocks();

    global.localStorage = {
      getItem: vi.fn().mockReturnValue('mocked_token'),
      setItem: vi.fn(),
      removeItem: vi.fn(),
      clear: vi.fn(),
    };
  });

  it('should get user profile', async () => {
    const mockResponse = { id: 1, name: 'John Doe' };
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await userApi.getProfile();
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/users/me`, {
      headers: getAuthHeaders(),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should update user profile', async () => {
    const mockResponse = { success: true };
    const profileData = { name: 'Jane Doe', email: 'jane@example.com' };
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await userApi.updateProfile(profileData);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/users/me`, {
      method: 'PUT',
      headers: getAuthHeaders(),
      body: JSON.stringify(profileData),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should update user password', async () => {
    const mockResponse = { success: true };
    const currentPassword = 'oldPassword';
    const newPassword = 'newPassword';
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await userApi.updatePassword(currentPassword, newPassword);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/users/me/password`, {
      method: 'PUT',
      headers: getAuthHeaders(),
      body: JSON.stringify({
        current_password: currentPassword,
        new_password: newPassword,
      }),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should delete user account', async () => {
    const mockResponse = { success: true };
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await userApi.deleteAccount();
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/users/me`, {
      method: 'DELETE',
      headers: getAuthHeaders(),
    });
    expect(response).toEqual(mockResponse);
  });
});
