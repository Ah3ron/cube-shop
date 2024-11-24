import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { cartApi } from '$lib/api/cart';

function createCartStore() {
	const { subscribe, set, update } = writable([]);

	return {
		subscribe,
		set,
		update,
		async fetch() {
			if (!browser) return;

			try {
				const items = await cartApi.get();
				set(items.sort((a, b) => a.id - b.id));
			} catch (err) {
				console.error('Error fetching cart:', err);
				set([]);
			}
		},
		async removeItem(itemId) {
			if (!browser) return;

			try {
				await cartApi.remove(itemId);
				await this.fetch();
			} catch (err) {
				console.error('Error removing item:', err);
			}
		},
		async updateQuantity(itemId, quantity) {
			if (!browser) return;

			try {
				await cartApi.update(itemId, quantity);
				await this.fetch();
			} catch (err) {
				console.error('Error updating quantity:', err);
			}
		}
	};
}

export const cart = createCartStore();
