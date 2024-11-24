<script>
	import { onMount, onDestroy } from 'svelte';
	import { Chart } from 'chart.js/auto';
	import { adminApi } from '$lib/api/admin';
	import { productsApi } from '$lib/api/products';

	let activeTab = 'catalog';
	let isLoading = true;
	let error = null;
	let products = [];
	let orders = [];
	let stats = {};
	let chartCanvas;
	let chart;

	const validStatuses = ['pending', 'approved', 'cancelled'];

	async function fetchProducts() {
		try {
			products = await productsApi.getAll();
		} catch (err) {
			error = err.message;
		}
	}

	async function fetchOrders() {
		try {
			orders = await adminApi.getAllOrders();
		} catch (err) {
			error = err.message;
		}
	}

	async function updateStatus(orderId, newStatus) {
		try {
			await adminApi.updateOrderStatus(orderId, newStatus);
			await fetchOrders();
		} catch (err) {
			error = err.message;
		}
	}

	async function deleteProduct(productId) {
		try {
			if (confirm('Вы уверены, что хотите удалить этот продукт?')) {
				await productsApi.delete(productId);
				await fetchProducts();
			}
		} catch (err) {
			error = err.message;
		}
	}

	async function fetchStats() {
		try {
			stats = await adminApi.getStats();
		} catch (err) {
			error = err.message;
		}
	}

	function initChart() {
		if (chartCanvas && stats.monthlySales) {
			try {
				chart = new Chart(chartCanvas, {
					type: 'line',
					data: {
						labels: stats.monthlySales.map(sale => sale.month),
						datasets: [{
							label: 'Monthly Sales',
							data: stats.monthlySales.map(sale => sale.amount),
							borderColor: 'rgb(75, 192, 192)',
							tension: 0.1
						}]
					}
				});
			} catch (err) {
				error = 'Failed to initialize chart';
				console.error(err);
			}
		}
	}

	onMount(async () => {
		try {
			isLoading = true;
			await Promise.all([fetchProducts(), fetchOrders(), fetchStats()]);
			initChart();
		} catch (err) {
			error = err.message;
		} finally {
			isLoading = false;
		}
	});

	onDestroy(() => {
		if (chart) {
			chart.destroy();
		}
	});

	$: if (activeTab === 'catalog') {
		fetchProducts();
	} else if (activeTab === 'orders') {
		fetchOrders();
	}
</script>

<div class="container mx-auto px-3 py-8">
	<h1 class="text-3xl font-bold mb-8">Admin Dashboard</h1>

	<div class="tabs tabs-boxed mb-6">
		<button
			class="tab {activeTab === 'catalog' ? 'tab-active' : ''}"
			on:click={() => (activeTab = 'catalog')}
		>
			Catalog Management
		</button>
		<button
			class="tab {activeTab === 'orders' ? 'tab-active' : ''}"
			on:click={() => (activeTab = 'orders')}
		>
			Orders
		</button>
		<button
			class="tab {activeTab === 'stats' ? 'tab-active' : ''}"
			on:click={() => (activeTab = 'stats')}
		>
			Statistics
		</button>
	</div>

	{#if isLoading}
		<div class="flex justify-center">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error shadow-lg">
			<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
					  d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
			</svg>
			<span>{error}</span>
		</div>
	{:else}
		{#if activeTab === 'catalog'}
			<div in:fade={{ duration: 300 }}>
				<div class="card bg-base-100 shadow">
					<div class="card-body">
						<h2 class="card-title">Product Management</h2>
						<div class="overflow-x-auto">
							<table class="table">
								<thead>
									<tr>
										<th>ID</th>
										<th>Name</th>
										<th>Price</th>
										<th>Stock</th>
										<th>Actions</th>
									</tr>
								</thead>
								<tbody>
									{#each products as product}
										<tr>
											<td>{product.ID}</td>
											<td>{product.name}</td>
											<td>${product.price}</td>
											<td>{product.stock}</td>
											<td>
												<button class="btn btn-sm btn-outline" 
														on:click={() => handleEditProduct(product.ID)}>
													<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
														<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
															  d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
													</svg>
												</button>
												<button class="btn btn-sm btn-error ml-2" 
														on:click={() => deleteProduct(product.ID)}>
													<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
														<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
															  d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
													</svg>
												</button>
											</td>
										</tr>
									{/each}
								</tbody>
							</table>
						</div>
						<button class="btn btn-primary mt-4">Add New Product</button>
					</div>
				</div>
			</div>
		{/if}

		{#if activeTab === 'orders'}
			<div class="card bg-base-100 shadow">
				<div class="card-body">
					<h2 class="card-title">Order Management</h2>
					<div class="overflow-x-auto">
						<table class="table">
							<thead>
								<tr>
									<th>Order ID</th>
									<th>User ID</th>
									<th>Date</th>
									<th>Status</th>
									<th>Total</th>
									<th>Actions</th>
								</tr>
							</thead>
							<tbody>
								{#each orders as order}
									<tr>
										<td>{order.ID}</td>
										<td>{order.user_id}</td>
										<td>{order.created_at}</td>
										<td>
											<select
												class="select select-bordered select-sm"
												on:change={(e) => updateStatus(order.id, e.target.value)}
											>
												{#each validStatuses as status}
													<option>{status}</option>
												{/each}
											</select>
										</td>
										<td>${order.total_price}</td>
										<td>
											<button class="btn btn-sm btn-outline">View</button>
										</td>
									</tr>
								{/each}
							</tbody>
						</table>
					</div>
				</div>
			</div>
		{/if}

		{#if activeTab === 'stats'}
			<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
				<div class="card bg-base-100 shadow">
					<div class="card-body">
						<h2 class="card-title">Sales Overview</h2>
						<canvas bind:this={chartCanvas}></canvas>
					</div>
				</div>

				<div class="card bg-base-100 shadow">
					<div class="card-body">
						<h2 class="card-title">Quick Stats</h2>
						<div class="stats stats-vertical shadow">
							<div class="stat">
								<div class="stat-figure text-primary">
									<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
											  d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
									</svg>
								</div>
								<div class="stat-title">Общие продажи</div>
								<div class="stat-value text-primary">
									${orders.reduce((sum, order) => sum + order.total_price, 0).toFixed(2)}
								</div>
							</div>

							<div class="stat">
								<div class="stat-title">Total Orders</div>
								<div class="stat-value">{stats.stats.totalOrders}</div>
							</div>

							<div class="stat">
								<div class="stat-title">Active Products</div>
								<div class="stat-value">{products.reduce((sum, product) => sum + product.stock, 0)}</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		{/if}
	{/if}
</div>

<style>
	:global(.chart-container) {
		position: relative;
		height: 300px;
		width: 100%;
	}
</style>