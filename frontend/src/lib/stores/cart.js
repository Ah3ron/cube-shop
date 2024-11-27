import { writable } from 'svelte/store';
import { browser } from '$app/environment';

function createCartStore() {
	const { subscribe, set, update } = writable([]);

	return {
		subscribe,
		set,
		update,
		async fetch() {
			if (!browser) return;

			try {
				const token = localStorage.getItem('token');
				if (!token) {
					set([]);
					return;
				}

				const response = await fetch('/api/cart', {
					headers: {
						Authorization: `Bearer ${token}`
					}
				});

				if (!response.ok) throw new Error('Failed to fetch cart');
				const items = await response.json();
				set(items.sort((a, b) => a.id - b.id));
			} catch (err) {
				console.error('Error fetching cart:', err);
				set([]);
			}
		},
		async removeItem(itemId) {
			if (!browser) return;

			try {
				const token = localStorage.getItem('token');
				if (!token) return;

				const response = await fetch(`/api/cart/${itemId}`, {
					method: 'DELETE',
					headers: {
						Authorization: `Bearer ${token}`
					}
				});

				if (!response.ok) throw new Error('Failed to remove item');
				await this.fetch();
			} catch (err) {
				console.error('Error removing item:', err);
			}
		},
		async updateQuantity(itemId, quantity) {
			if (!browser) return;

			try {
				const token = localStorage.getItem('token');
				if (!token) return;

				const response = await fetch(`/api/cart/${itemId}`, {
					method: 'PATCH',
					headers: {
						'Content-Type': 'application/json',
						Authorization: `Bearer ${token}`
					},
					body: JSON.stringify({ quantity })
				});

				if (!response.ok) throw new Error('Failed to update quantity');
				await this.fetch();
			} catch (err) {
				console.error('Error updating quantity:', err);
			}
		}
	};
}

export const cart = createCartStore();
