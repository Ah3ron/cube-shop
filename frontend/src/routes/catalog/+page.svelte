<script>
    import { onMount } from 'svelte';

    let products = [];
    let loading = true;
    let error = null;

    // Filter states
    let filters = {
        category: '',
        difficulty: '',
        brand: '',
        price: { min: 0, max: 100 },
        inStock: false
    };

    // Declare Set variables first
    let categories = new Set();
    let difficulties = new Set();
    let brands = new Set();

    // Reactive statement to update Sets
    $: {
        if (products.length > 0) {
            categories = new Set(products.map(p => p.category));
            difficulties = new Set(products.map(p => p.difficulty));
            brands = new Set(products.map(p => p.brand));
        }
    }

    onMount(async () => {
        loading = true;
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                window.location.href = '/auth/login';
                return;

            }

            const response = await fetch('/api/products', {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });

            if (response.status === 401) {
                localStorage.removeItem('token');
                window.location.href = '/auth/login';
                return;
            }

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
			(!filters.inStock || p.inStock) &&
			p.price >= filters.price.min &&
			p.price <= filters.price.max
		);
	});
</script>

<div class="container mx-auto px-4 py-8">
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
								bind:value={filters.price.min}
								min="0"
							/>
							<input
								type="number"
								class="input input-bordered w-full"
								bind:value={filters.price.max}
								min="0"
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
								price: { min: 0, max: 100 },
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
		<div class="flex-1">
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
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each filteredProducts as product}
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
	</div>
</div>
