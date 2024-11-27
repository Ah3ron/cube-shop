const API_URL = '/api/v1';

async function request(endpoint, options = {}) {
    if (typeof window === 'undefined') return {};

    const token = localStorage.getItem('token');
    const headers = {
        'Content-Type': 'application/json',
        ...(token && { Authorization: `Bearer ${token}` }),
        ...options.headers,
    };

    const response = await fetch(`${API_URL}${endpoint}`, {
        ...options,
        headers,
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Что-то пошло не так');
    }

    return response.json();
}

// Products API
export const productsApi = {
    getAll: (filters = {}) => {
        const params = new URLSearchParams(filters);
        return request(`/products?${params}`);
    },
    
    getById: (id) => request(`/products/${id}`),
    getCategories: () => request('/products/categories'),
    search: (query) => request(`/products/search?q=${query}`),
    
    create: (product) => request('/products', {
        method: 'POST',
        body: JSON.stringify(product),
    }),
    
    update: (id, product) => request(`/products/${id}`, {
        method: 'PUT',
        body: JSON.stringify(product),
    }),
    
    delete: (id) => request(`/products/${id}`, { method: 'DELETE' }),
};

// Cart API
export const cartApi = {
    get: () => request('/cart'),
    
    addItem: (productId, quantity) => request('/cart/items', {
        method: 'POST',
        body: JSON.stringify({ product_id: productId, quantity }),
    }),
    
    updateItem: (id, quantity) => request(`/cart/items/${id}`, {
        method: 'PUT',
        body: JSON.stringify({ quantity }),
    }),
    
    removeItem: (id) => request(`/cart/items/${id}`, { method: 'DELETE' }),
    clear: () => request('/cart', { method: 'DELETE' }),
};

// Orders API
export const ordersApi = {
    checkout: (shippingDetails) => request('/orders/checkout', {
        method: 'POST',
        body: JSON.stringify({ shipping_details: shippingDetails }),
    }),
    
    getAll: () => request('/orders'),
    getById: (id) => request(`/orders/${id}`),
    cancel: (id) => request(`/orders/${id}/cancel`, { method: 'PUT' }),
};

// Auth API
export const authApi = {
    register: (email, password, name) => request('/auth/register', {
        method: 'POST',
        body: JSON.stringify({ email, password, name }),
    }),
    
    login: (email, password) => request('/auth/login', {
        method: 'POST',
        body: JSON.stringify({ email, password }),
    }),
    
    logout: () => request('/auth/logout', { method: 'POST' }),
}; 