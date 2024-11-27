<script>
	import { fade, fly, scale } from 'svelte/transition';
	import { onMount } from 'svelte';
	import AddToCartButton from '$lib/AddToCartButton.svelte';
	import { productsApi } from '$lib/api/products';

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
			const [productsData, categoriesData] = await Promise.all([
				productsApi.getAll(filters),
				productsApi.getCategories()
			]);
			
			products = productsData;
			categories = categoriesData;
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
		<div class="flex justify-center">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error shadow-lg">
			<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
			</svg>
			<span>{error}</span>
		</div>
	{:else}
		<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
			{#each products as product}
				<div class="card bg-base-100 shadow-xl hover:shadow-2xl transition-all duration-300" transition:fade>
					<figure>
						<img 
							src={product.image_url} 
							alt={product.name} 
							class="w-full h-48 object-cover"
						/>
					</figure>
					<div class="card-body">
						<h3 class="card-title">{product.name}</h3>
						<p class="text-primary text-xl font-bold">${product.price}</p>
						<div class="card-actions justify-end">
							<AddToCartButton {product} />
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
