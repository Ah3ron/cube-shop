import { apiRequest } from './client';

export const productsApi = {
  getAll: () => 
    apiRequest('/products'),
    
  getById: (id) =>
    apiRequest(`/products/${id}`),
    
  create: (product) =>
    apiRequest('/admin/products', {
      method: 'POST',
      body: JSON.stringify(product)
    }),
    
  update: (id, product) =>
    apiRequest(`/admin/products/${id}`, {
      method: 'PUT',
      body: JSON.stringify(product)
    }),
    
  delete: (id) =>
    apiRequest(`/admin/products/${id}`, {
      method: 'DELETE'
    })
}; 