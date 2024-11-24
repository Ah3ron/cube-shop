import { apiRequest } from './client';

export const cartApi = {
  get: () => 
    apiRequest('/cart'),
    
  add: (productId, quantity) =>
    apiRequest('/cart', {
      method: 'POST',
      body: JSON.stringify({ product_id: productId, quantity })
    }),
    
  update: (itemId, quantity) =>
    apiRequest(`/cart/${itemId}`, {
      method: 'PATCH',
      body: JSON.stringify({ quantity })
    }),
    
  remove: (itemId) =>
    apiRequest(`/cart/${itemId}`, {
      method: 'DELETE'
    })
}; 