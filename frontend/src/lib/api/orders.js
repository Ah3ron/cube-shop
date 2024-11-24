import { apiRequest } from './client';

export const ordersApi = {
  getAll: () => 
    apiRequest('/orders'),
    
  getById: (id) =>
    apiRequest(`/orders/${id}`),
    
  updateStatus: (id, status) =>
    apiRequest(`/orders/${id}/status`, {
      method: 'PUT',
      body: JSON.stringify({ status })
    }),
    
  cancel: (id) =>
    apiRequest(`/orders/${id}/cancel`, {
      method: 'PUT'
    })
}; 