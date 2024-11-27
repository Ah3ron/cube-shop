<script>
	import { cart } from '$lib/stores/cart';
	import { cartApi } from '$lib/api/cart.js';

	export let product;
	let loading = false;

	async function handleAddToCart() {
		loading = true;
		try {
			await cartApi.addItem(product.ID, 1);
		} catch (error) {
			console.error('Failed to add to cart:', error);
		}
		loading = false;
	}
</script>

<button class="btn btn-primary w-fit relative" on:click={handleAddToCart} disabled={loading}>
	{#if loading}
		<span class="loading loading-infinity text-primary absolute inset-0 m-auto"></span>
		<span class="opacity-0">Add to Cart</span>
	{:else}
		Add to Cart
	{/if}
</button>
