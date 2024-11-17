<!-- src/routes/catalog/+page.svelte -->
<script>
	import { onMount } from 'svelte';

	let products = [];
	let loading = true;
	let error = null;

	onMount(async () => {
		try {
			const response = await fetch('/api/products');
			if (!response.ok) throw new Error('Failed to fetch products');
			products = await response.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});
</script>

<div class="container mx-auto px-4 py-8">
	<h1 class="text-3xl font-bold mb-8">Product Catalog</h1>

	{#if loading}
		<div class="flex justify-center">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="stroke-current shrink-0 h-6 w-6"
				fill="none"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
				/>
			</svg>
			<span>{error}</span>
		</div>
	{:else}
		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
			{#each products as product}
				<div class="card bg-base-100 shadow-xl hover:shadow-2xl transition-all duration-300">
					<figure class="px-4 pt-4">
						<img
							src={product.image || '/placeholder.png'}
							alt={product.name}
							class="rounded-xl h-48 w-full object-cover"
						/>
					</figure>
					<div class="card-body">
						<h2 class="card-title">{product.name}</h2>
						<p class="text-sm opacity-70">{product.description}</p>
						<div class="flex justify-between items-center mt-4">
							<span class="text-xl font-bold">${product.price}</span>
							<button class="btn btn-primary">Add to Cart</button>
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
