import { describe, it, expect, vi, beforeEach } from 'vitest';
import { ordersApi } from '../lib/api/orders'; // Путь к вашему файлу ordersApi
import { API_URL, getAuthHeaders } from '../lib/api/config';

describe('ordersApi', () => {
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

	it('should checkout an order', async () => {
		const mockResponse = { success: true };
		const shippingDetails = { address: '123 Main St', city: 'Anytown', zip: '12345' };
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await ordersApi.checkout(shippingDetails);

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/orders/checkout`, {
			method: 'POST',
			headers: getAuthHeaders(),
			body: JSON.stringify({ shipping_details: shippingDetails })
		});
		expect(response).toEqual(mockResponse);
	});

	it('should get all orders', async () => {
		const mockResponse = [{ id: 1, status: 'shipped' }];
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await ordersApi.getAll();

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/orders`, {
			headers: getAuthHeaders()
		});
		expect(response).toEqual(mockResponse);
	});

	it('should get an order by id', async () => {
		const mockResponse = { id: 1, status: 'shipped' };
		const orderId = 1;
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await ordersApi.getById(orderId);

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/orders/${orderId}`, {
			headers: getAuthHeaders()
		});
		expect(response).toEqual(mockResponse);
	});

	it('should update the status of an order', async () => {
		const mockResponse = { success: true };
		const orderId = 1;
		const status = 'delivered';
		global.fetch = vi.fn().mockResolvedValue({
			ok: true,
			json: vi.fn().mockResolvedValue(mockResponse)
		});

		const response = await ordersApi.updateStatus(orderId, status);

		expect(fetch).toHaveBeenCalledWith(`${API_URL}/orders/${orderId}/status`, {
			method: 'PUT',
			headers: getAuthHeaders(),
			body: JSON.stringify({ status })
		});
		expect(response).toEqual(mockResponse);
	});
});
