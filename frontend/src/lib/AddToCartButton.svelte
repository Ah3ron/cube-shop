<script>
	import { cart } from '$lib/stores/cart';

	export let product;
	let loading = false;

	async function handleAddToCart() {
		loading = true;
		try {
			const token = localStorage.getItem('token');
			const response = await fetch('/api/cart', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({
					product_id: product.ID,
					quantity: 1
				})
			});

			if (!response.ok) {
				const error = await response.json();
				throw new Error(error.error);
			}

			await cart.fetch();
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
