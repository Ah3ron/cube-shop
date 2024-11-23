<script>
	import { fade, fly } from 'svelte/transition';
	import { onMount } from 'svelte';
	import AddToCartButton from '$lib/AddToCartButton.svelte';

	let product = null;
	let loading = true;
	let error = null;
	let imgLoading = true;
	let activeTab = 'description';

	async function fetchProduct() {
		try {
			const url = new URL(window.location.href);
			const productId = url.searchParams.get('id');

			const response = await fetch(`/api/products/${productId}`);

			if (!response.ok) {
				throw new Error('Product not found');
			}

			product = await response.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	onMount(async () => {
		fetchProduct();
	});
</script>

<div class="container mx-auto px-4 mt-16">
	<a href="/catalog" class="btn btn-ghost">
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
				d="M10 19l-7-7m0 0l7-7m-7 7h18"
			/>
		</svg>
		Back to Catalog
	</a>

	{#if loading}
		<div class="flex justify-center" in:fade={{ duration: 200 }}>
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error" in:fly={{ x: -20, duration: 500 }}>
			<span>{error}</span>
		</div>
	{:else}
		<div class="flex flex-col lg:flex-row gap-8" in:fly={{ y: 50, duration: 1000 }}>
			<div class="lg:w-1/2">
				<div class="relative">
					{#if imgLoading}
						<div class="absolute inset-0 flex items-center justify-center bg-base-200 rounded-xl">
							<span class="loading loading-spinner loading-lg"></span>
						</div>
					{/if}
					<img
						src={product.image_url || '/placeholder.png'}
						alt={product.name}
						class="w-full rounded-xl shadow-lg transition-opacity duration-300"
						class:opacity-0={imgLoading}
						class:opacity-100={!imgLoading}
						on:load={() => (imgLoading = false)}
					/>
				</div>
			</div>

			<div class="lg:w-1/2">
				<div class="bg-base-100 p-8 rounded-xl shadow-lg">
					<h1 class="text-4xl font-bold mb-2">{product.name}</h1>

					<div class="divider my-4"></div>

					<div class="flex items-center gap-6 mb-6">
						<div class="flex flex-col">
							<span class="text-sm text-base-content/70">Price</span>
							<span class="text-4xl font-bold text-primary">${product.price}</span>
						</div>
						<AddToCartButton {product} />
					</div>
				</div>

				<div class="mt-4 bg-base-100 rounded-xl shadow-lg overflow-hidden">
					<div class="tabs tabs-boxed bg-base-100 p-2 gap-2">
						<button
							class="tab flex-1 {activeTab === 'description' ? 'tab-active' : ''}"
							on:click={() => (activeTab = 'description')}
						>
							Description
						</button>
						<button
							class="tab flex-1 {activeTab === 'details' ? 'tab-active' : ''}"
							on:click={() => (activeTab = 'details')}
						>
							Details
						</button>
						<button
							class="tab flex-1 {activeTab === 'shipping' ? 'tab-active' : ''}"
							on:click={() => (activeTab = 'shipping')}
						>
							Shipping
						</button>
					</div>

					<div class="p-6" in:fade={{ duration: 200 }}>
						{#if activeTab === 'description'}
							<div class="prose max-w-none">
								<p>{product.description}</p>
							</div>
						{:else if activeTab === 'details'}
							<div class="space-y-4">
								<div class="stats stats-vertical shadow w-full">
									<div class="stat">
										<div class="stat-title">Category</div>
										<div class="stat-value text-lg">{product.category}</div>
									</div>
									<div class="stat">
										<div class="stat-title">Difficulty</div>
										<div class="stat-value text-lg">{product.difficulty}</div>
									</div>
									<div class="stat">
										<div class="stat-title">Brand</div>
										<div class="stat-value text-lg">{product.brand}</div>
									</div>
									<div class="stat">
										<div class="stat-title">Stock</div>
										<div class="stat-value text-lg">{product.stock} units</div>
									</div>
								</div>
							</div>
						{:else if activeTab === 'shipping'}
							<div class="space-y-4">
								<div class="card bg-primary text-primary-content">
									<div class="card-body">
										<h3 class="card-title">Standard Shipping</h3>
										<p>3-5 business days</p>
									</div>
								</div>
								<div class="card bg-secondary text-secondary-content">
									<div class="card-body">
										<h3 class="card-title">Express Shipping</h3>
										<p>1-2 business days</p>
									</div>
								</div>
							</div>
						{/if}
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
