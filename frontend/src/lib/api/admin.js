import { apiRequest } from './client';

export const adminApi = {
  getStats: () =>
    apiRequest('/admin/stats'),
    
  getAllOrders: () =>
    apiRequest('/admin/orders'),
    
  getAllUsers: () =>
    apiRequest('/admin/users'),
    
  updateOrderStatus: (orderId, status) =>
    apiRequest(`/admin/orders/${orderId}/status`, {
      method: 'PUT',
      body: JSON.stringify({ status })
    })
}; 