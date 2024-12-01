import { describe, it, expect, vi, beforeEach } from 'vitest';
import { productsApi } from '../lib/api/products'; // Путь к вашему файлу productsApi
import { API_URL } from '../lib/api/config';

describe('productsApi', () => {
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

  it('should get all products with filters', async () => {
    const mockResponse = [{ id: 1, name: 'Product 1' }];
    const filters = { category: 'electronics', sort: 'price' };
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await productsApi.getAll(filters);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/products?category=electronics&sort=price`);
    expect(response).toEqual(mockResponse);
  });

  it('should get a product by id', async () => {
    const mockResponse = { id: 1, name: 'Product 1' };
    const productId = 1;
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await productsApi.getById(productId);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/products/${productId}`);
    expect(response).toEqual(mockResponse);
  });

  it('should search for products', async () => {
    const mockResponse = [{ id: 1, name: 'Product 1' }];
    const query = 'laptop';
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue(mockResponse),
    });

    const response = await productsApi.search(query);
    
    expect(fetch).toHaveBeenCalledWith(`${API_URL}/products/search?q=${query}`);
    expect(response).toEqual(mockResponse);
  });
});
