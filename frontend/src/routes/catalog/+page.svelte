<script>
	import { fade, fly, scale } from 'svelte/transition';
	import { onMount } from 'svelte';
	import AddToCartButton from '$lib/AddToCartButton.svelte';

	let products = [];
	let loading = true;
	let imgLoading = true;
	let error = null;
	let debounceTimer;

	// Filter states
	let filters = {
		category: '',
		difficulty: '',
		brand: '',
		price: { min: 0, max: 0 },
		inStock: false
	};

	// Declare Set variables first
	let categories = new Set();
	let difficulties = new Set();
	let brands = new Set();

	// Reactive statement to update Sets
	$: {
		if (products.length > 0) {
			categories = new Set(products.map((p) => p.category));
			difficulties = new Set(products.map((p) => p.difficulty));
			brands = new Set(products.map((p) => p.brand));
			filters.price.max = Math.max(...products.map((p) => p.price));
		}
	}

	function debouncePriceUpdate(value, type) {
		clearTimeout(debounceTimer);
		debounceTimer = setTimeout(() => {
			filters.price[type] = value;
		}, 300);
	}

	onMount(async () => {
		loading = true;
		try {
			const response = await fetch('/api/products');

			if (!response.ok) {
				throw new Error('Failed to fetch products');
			}

			products = await response.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});

	$: filteredProducts = products.filter((p) => {
		return (
			(!filters.category || p.category === filters.category) &&
			(!filters.difficulty || p.difficulty === filters.difficulty) &&
			(!filters.brand || p.brand === filters.brand) &&
			(!filters.inStock || p.stock > 0) &&
			p.price >= filters.price.min &&
			p.price <= filters.price.max
		);
	});
</script>

<div class="container mx-auto px-3 py-8">
	<div class="flex flex-col md:flex-row gap-8">
		<!-- Filters -->
		<div class="w-full md:w-64 space-y-4">
			<div class="card bg-base-100 shadow">
				<div class="card-body">
					<h2 class="card-title">Filters</h2>

					<div class="form-control">
						<!-- svelte-ignore a11y_label_has_associated_control -->
						<label class="label">Category</label>
						<select class="select select-bordered" bind:value={filters.category}>
							<option value="">All</option>
							{#each [...categories] as category}
								<option value={category}>{category}</option>
							{/each}
						</select>
					</div>

					<div class="form-control">
						<!-- svelte-ignore a11y_label_has_associated_control -->
						<label class="label">Difficulty</label>
						<select class="select select-bordered" bind:value={filters.difficulty}>
							<option value="">All</option>
							{#each [...difficulties] as difficulty}
								<option value={difficulty}>{difficulty}</option>
							{/each}
						</select>
					</div>

					<div class="form-control">
						<!-- svelte-ignore a11y_label_has_associated_control -->
						<label class="label">Brand</label>
						<select class="select select-bordered" bind:value={filters.brand}>
							<option value="">All</option>
							{#each [...brands] as brand}
								<option value={brand}>{brand}</option>
							{/each}
						</select>
					</div>

					<div class="form-control">
						<!-- svelte-ignore a11y_label_has_associated_control -->
						<label class="label">Price Range</label>
						<div class="flex gap-2">
							<input
								type="number"
								class="input input-bordered w-full"
								value={filters.price.min}
								min="0"
								on:input={(e) => {
									if (e.target.value < 0) e.target.value = 0;
									if (e.target.value > filters.price.max) {
										filters.price.max = Number(e.target.value);
									}
									debouncePriceUpdate(Number(e.target.value), 'min');
								}}
							/>
							<input
								type="number"
								class="input input-bordered w-full"
								value={filters.price.max}
								min={filters.price.min}
								on:input={(e) => {
									if (e.target.value < filters.price.min) {
										e.target.value = filters.price.min;
									}
									debouncePriceUpdate(Number(e.target.value), 'max');
								}}
							/>
						</div>
					</div>

					<div class="form-control">
						<label class="label cursor-pointer">
							<span class="label-text">In Stock Only</span>
							<input type="checkbox" class="toggle" bind:checked={filters.inStock} />
						</label>
					</div>

					<button
						class="btn btn-outline btn-sm mt-4"
						on:click={() => {
							filters = {
								category: '',
								difficulty: '',
								brand: '',
								price: { min: 0, max: Math.max(...products.map((p) => p.price)) },
								inStock: false
							};
						}}
					>
						Reset Filters
					</button>
				</div>
			</div>
		</div>

		<!-- Products Grid -->
		<div class="flex-1" in:fly={{ y: 50, duration: 1000 }}>
			<h1 class="text-3xl font-bold mb-8" in:fly={{ y: -20, duration: 800 }}>Product Catalog</h1>

			{#if loading}
				<div class="flex justify-center" in:fade={{ duration: 200 }}>
					<span class="loading loading-spinner loading-lg"></span>
				</div>
			{:else if error}
				<div class="alert alert-error" in:fly={{ x: -20, duration: 500 }}>
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
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each filteredProducts as product, i}
						<a href={`/product/?id=${product.ID}`} class="block">
							<div
								class="card bg-base-100 shadow-xl hover:shadow-2xl hover:scale-1 transition-all duration-300"
								in:fly={{ y: 50, duration: 800, delay: i * 150 }}
								out:scale={{ duration: 300 }}
							>
								<figure class="px-4 pt-4 relative">
									{#if imgLoading}
										<div
											class="absolute inset-0 flex items-center justify-center bg-base-200 rounded-xl"
										>
											<span class="loading loading-spinner loading-lg"></span>
										</div>
									{/if}
									<img
										src={product.image_url || '/placeholder.png'}
										alt={product.name}
										class="rounded-xl h-48 w-full object-cover transition-opacity duration-300"
										class:opacity-0={imgLoading}
										class:opacity-100={!imgLoading}
										on:load={() => (imgLoading = false)}
									/>
								</figure>
								<div class="card-body">
									<h2 class="card-title">{product.name}</h2>
									<div class="flex flex-wrap gap-2 mb-2 w-2/3">
										<div class="badge badge-primary">{product.category}</div>
										<div class="badge badge-outline badge-secondary">{product.difficulty}</div>
										<div class="badge badge-outline badge-accent">{product.brand}</div>
										<div class="badge badge-outline badge-ghost">Stock: {product.stock}</div>
									</div>
									<p class="text-sm opacity-70">{product.description}</p>
									<div class="flex justify-between items-center mt-4 gap-3 flex-wrap">
										<span class="text-xl font-bold">${product.price}</span>
										<AddToCartButton {product} />
									</div>
								</div>
							</div>
						</a>
					{/each}
				</div>
			{/if}
		</div>
	</div>
</div>
