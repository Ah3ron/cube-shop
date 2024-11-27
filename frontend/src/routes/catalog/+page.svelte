<script>
	import { fade, fly, scale } from 'svelte/transition';
	import { onMount } from 'svelte';
	import AddToCartButton from '$lib/AddToCartButton.svelte';
	import { productsApi } from '$lib/api/products.js';

	let products = [];
	let categories = [];
	let difficulties = [];
	let brands = [];
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

			categories = [...new Set(products.map((product) => product.category))];
			difficulties = [...new Set(products.map((product) => product.difficulty))];
			brands = [...new Set(products.map((product) => product.brand))];
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	function applyFilters() {
		loadProducts();
	}

	function resetFilters() {
		filters = {
			category: '',
			difficulty: '',
			brand: '',
			sort: ''
		};
		loadProducts();
	}

	onMount(loadProducts);
</script>

<div class="container mx-auto px-4 py-8">
	<div class="flex flex-col md:flex-row gap-8">
		<!-- Фильтры -->
		<div class=" w-full md:w-1/4">
			<div class="bg-base-100 p-4 rounded-box shadow-md">
				<h2 class="text-xl font-bold mb-4">Фильтры</h2>

				<div class="mb-4">
					<label for="category" class="block mb-2">Категория</label>
					<select id="category" class="select select-bordered w-full" bind:value={filters.category}>
						<option value="">Все категории</option>
						{#each categories as category}
							<option value={category}>{category}</option>
						{/each}
					</select>
				</div>

				<div class="mb-4">
					<label for="difficulty" class="block mb-2">Сложность</label>
					<select
						id="difficulty"
						class="select select-bordered w-full"
						bind:value={filters.difficulty}
					>
						<option value="">Любая сложность</option>
						{#each difficulties as difficulty}
							<option value={difficulty}>{difficulty}</option>
						{/each}
					</select>
				</div>

				<div class="mb-4">
					<label for="brand" class="block mb-2">Бренд</label>
					<select id="brand" class="select select-bordered w-full" bind:value={filters.brand}>
						<option value="">Все бренды</option>
						{#each brands as brand}
							<option value={brand}>{brand}</option>
						{/each}
					</select>
				</div>

				<div class="mb-4">
					<label for="sort" class="block mb-2">Сортировка</label>
					<select id="sort" class="select select-bordered w-full" bind:value={filters.sort}>
						<option value="">По умолчанию</option>
						<option value="price_asc">Цена (по возрастанию)</option>
						<option value="price_desc">Цена (по убыванию)</option>
						<option value="name_asc">Название (А-Я)</option>
						<option value="name_desc">Название (Я-А)</option>
					</select>
				</div>

				<div class="flex gap-2 flex-wrap">
					<button class="btn btn-primary flex-1" on:click={applyFilters}>Применить</button>
					<button class="btn btn-outline flex-1" on:click={resetFilters}>Сбросить</button>
				</div>
			</div>
		</div>

		<!-- Список продуктов -->
		<div class="w-full md:w-3/4">
			{#if loading}
				<div class="flex justify-center">
					<span class="loading loading-spinner loading-lg"></span>
				</div>
			{:else if error}
				<div class="alert alert-error shadow-lg">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-6 w-6"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
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
			{:else if products.length === 0}
				<div class="text-center py-8">
					<p class="text-xl">Нет продуктов, соответствующих выбранным фильтрам.</p>
				</div>
			{:else}
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each products as product}
						<div
							class="card bg-base-100 shadow-xl hover:shadow-2xl transition-all duration-300"
							transition:fade
						>
							<figure>
								<img src={product.image_url} alt={product.name} class="w-full h-48 object-cover" />
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
	</div>
</div>
