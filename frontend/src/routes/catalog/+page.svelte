<script>
	import { fade, fly, scale } from 'svelte/transition';
	import { onMount } from 'svelte';
	import AddToCartButton from '$lib/AddToCartButton.svelte';

	let products = [];
	let categories = [];
	let loading = true;
	let error = null;

	let filters = {
		category: '',
		difficulty: '',
		brand: '',
		sort: ''
	};

	async function loadProducts() {
		try {
			loading = true;
			products = await productsApi.getAll(filters);
			categories = await productsApi.getCategories();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	$: {
		if (typeof window !== 'undefined') {
			loadProducts();
		}
	}
</script>

<div class="container mx-auto px-4 py-8">
	{#if loading}
		<div class="loading">Загрузка...</div>
	{:else if error}
		<div class="error">{error}</div>
	{:else}
		<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
			{#each products as product}
				<div class="card">
					<img src={product.imageUrl} alt={product.name} class="w-full h-48 object-cover" />
					<div class="p-4">
						<h3 class="text-lg font-bold">{product.name}</h3>
						<p class="text-gray-600">{product.price} ₽</p>
						<button class="btn btn-primary mt-2" on:click={() => addToCart(product.id)}>
							В корзину
						</button>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
