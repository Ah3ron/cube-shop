import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { cartApi } from '$lib/api/cart';

function createCartStore() {
	const { subscribe, set, update } = writable({ items: [], total: 0 });

	return {
		subscribe,
		set,
		update,
		async fetch() {
			if (!browser) return;

			try {
				const token = localStorage.getItem('token');
				if (!token) {
					set({ items: [], total: 0 });
					return;
				}

				const cart = await cartApi.get();
				set(cart);
			} catch (err) {
				console.error('Error fetching cart:', err);
				set({ items: [], total: 0 });
			}
		},
		async addItem(productId, quantity = 1) {
			if (!browser) return;

			try {
				await cartApi.addItem(productId, quantity);
				await this.fetch();
			} catch (err) {
				console.error('Error adding item to cart:', err);
				throw err;
			}
		},
		async removeItem(itemId) {
			if (!browser) return;

			try {
				await cartApi.removeItem(itemId);
				await this.fetch();
			} catch (err) {
				console.error('Error removing item from cart:', err);
				throw err;
			}
		},
		async updateQuantity(itemId, quantity) {
			if (!browser) return;

			try {
				await cartApi.updateItem(itemId, quantity);
				await this.fetch();
			} catch (err) {
				console.error('Error updating cart item quantity:', err);
				throw err;
			}
		},
		async clear() {
			if (!browser) return;

			try {
				await cartApi.clear();
				set({ items: [], total: 0 });
			} catch (err) {
				console.error('Error clearing cart:', err);
				throw err;
			}
		}
	};
}

export const cart = createCartStore();
