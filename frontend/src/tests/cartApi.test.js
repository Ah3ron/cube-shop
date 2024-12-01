import { describe, it, expect, vi, beforeEach } from 'vitest';
import { cartApi } from '../lib/api/cart'; // Путь к вашему файлу cartApi
import { API_URL, getAuthHeaders } from '../lib/api/config';

describe('cartApi', () => {
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

  it('should get cart items', async () => {
    const mockResponse = { items: [] };
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await cartApi.get();
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/cart`, {
      headers: getAuthHeaders(),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should add an item to the cart', async () => {
    const mockResponse = { success: true };
    const productId = 1;
    const quantity = 2;
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await cartApi.addItem(productId, quantity);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/cart/items`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ product_id: productId, quantity }),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should update an item in the cart', async () => {
    const mockResponse = { success: true };
    const itemId = 1;
    const quantity = 3;
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await cartApi.updateItem(itemId, quantity);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/cart/items/${itemId}`, {
      method: 'PUT',
      headers: getAuthHeaders(),
      body: JSON.stringify({ quantity }),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should remove an item from the cart', async () => {
    const mockResponse = { success: true };
    const itemId = 1;
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await cartApi.removeItem(itemId);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/cart/items/${itemId}`, {
      method: 'DELETE',
      headers: getAuthHeaders(),
    });
    expect(response).toEqual(mockResponse);
  });

  it('should clear the cart', async () => {
    const mockResponse = { success: true };
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await cartApi.clear();
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/cart`, {
      method: 'DELETE',
      headers: getAuthHeaders(),
    });
    expect(response).toEqual(mockResponse);
  });
});
