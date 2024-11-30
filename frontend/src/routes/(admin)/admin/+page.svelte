<script>
	import { onMount } from 'svelte';
	import { ordersApi } from '$lib/api/orders.js';
	import { productsApi } from '$lib/api/products';

	// State variables
	let products = [];
	let orders = [];
	let loading = { products: true, orders: true };
	let error = { products: null, orders: null };
	let activeTab = 'orders';

	// Modal states
	let editingProduct = null;
	let showEditModal = false;
	let showAddModal = false;
	let newProduct = {
		name: '',
		price: 0,
		stock: 0,
		description: '',
		image_url: ''
	};

	onMount(async () => {
		await Promise.all([loadProducts(), loadOrders()]);
	});

	function switchTab(tab) {
		activeTab = tab;
	}

	// Orders functions
	async function loadOrders() {
		try {
			orders = await ordersApi.getAll();
		} catch (err) {
			error.orders = err.message;
		} finally {
			loading.orders = false;
		}
	}

	async function updateOrderStatus(orderId, newStatus) {
		try {
			await ordersApi.updateStatus(orderId, newStatus);
			orders = orders.map((order) =>
				order.ID === orderId ? { ...order, status: newStatus } : order
			);
		} catch (err) {
			alert('Ошибка при обновлении статуса: ' + err.message);
		}
	}

	function getStatusColor(status) {
		switch (status) {
			case 'pending':
				return 'badge-warning';
			case 'completed':
				return 'badge-success';
			case 'cancelled':
				return 'badge-error';
			default:
				return 'badge-ghost';
		}
	}

	// Products functions
	async function loadProducts() {
		try {
			products = await productsApi.getAll();
		} catch (err) {
			error.products = err.message;
		} finally {
			loading.products = false;
		}
	}

	async function handleDelete(id) {
		if (!confirm('Вы уверены, что хотите удалить этот товар?')) return;

		try {
			await productsApi.delete(id);
			products = products.filter((p) => p.ID !== id);
		} catch (err) {
			alert('Ошибка при удалении товара: ' + err.message);
		}
	}

	async function handleEdit(product) {
		editingProduct = { ...product };
		showEditModal = true;
	}

	async function saveProduct() {
		try {
			const updatedProduct = await productsApi.update(editingProduct.ID, editingProduct);
			products = products.map((p) => (p.ID === updatedProduct.ID ? updatedProduct : p));
			showEditModal = false;
		} catch (err) {
			alert('Ошибка при обновлении товара: ' + err.message);
		}
	}

	async function handleAddProduct() {
		try {
			const addedProduct = await productsApi.create(newProduct);
			products = [...products, addedProduct];
			showAddModal = false;
			newProduct = {
				name: '',
				price: 0,
				stock: 0,
				description: '',
				image_url: ''
			};
		} catch (err) {
			alert('Ошибка при добавлении товара: ' + err.message);
		}
	}
</script>

<div class="drawer lg:drawer-open">
	<input id="admin-drawer" type="checkbox" class="drawer-toggle" />

	<div class="drawer-content">
		<div class="lg:hidden navbar bg-base-100 mb-4">
			<div class="flex-none">
				<label for="admin-drawer" class="btn btn-square btn-ghost drawer-button">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						class="inline-block w-5 h-5 stroke-current"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M4 6h16M4 12h16M4 18h16"
						></path>
					</svg>
				</label>
			</div>
			<div class="flex-1">
				<span class="text-xl font-bold"> Админ панель </span>
			</div>
		</div>

		<div class="p-4 space-y-4">
			{#if activeTab === 'orders'}
				<div class="bg-base-100 rounded-box p-4">
					<h1 class="text-2xl font-bold mb-6">Управление заказами</h1>
					{#if loading.orders}
						<div class="flex justify-center">
							<span class="loading loading-spinner loading-lg"></span>
						</div>
					{:else if error.orders}
						<div class="alert alert-error">{error.orders}</div>
					{:else}
						<div class="overflow-x-auto">
							<table class="table w-full">
								<thead>
									<tr>
										<th>ID</th>
										<th>Дата</th>
										<th>Клиент</th>
										<th>Сумма</th>
										<th>Статус</th>
										<th>Действия</th>
									</tr>
								</thead>
								<tbody>
									{#each orders as order}
										<tr>
											<td>{order.ID}</td>
											<td>{new Date(order.created_at).toLocaleDateString()}</td>
											<td>{order.shipping_details.name}</td>
											<td>${order.total_price}</td>
											<td>
												<span class="badge {getStatusColor(order.status)}">
													{order.status}
												</span>
											</td>
											<td>
												<select
													class="select select-bordered select-sm"
													value={order.status}
													on:change={(e) => updateOrderStatus(order.ID, e.target.value)}
												>
													<option value="pending">В ожидании</option>
													<option value="completed">Завершен</option>
													<option value="cancelled">Отменен</option>
												</select>
											</td>
										</tr>
									{/each}
								</tbody>
							</table>
						</div>
					{/if}
				</div>
			{:else}
				<div class="bg-base-100 rounded-box p-4">
					<div class="flex justify-between items-center mb-6">
						<h1 class="text-2xl font-bold">Управление товарами</h1>
						<button class="btn btn-primary" on:click={() => (showAddModal = true)}>
							Добавить товар
						</button>
					</div>
					{#if loading.products}
						<div class="flex justify-center">
							<span class="loading loading-spinner loading-lg"></span>
						</div>
					{:else if error.products}
						<div class="alert alert-error">{error.products}</div>
					{:else}
						<div class="overflow-x-auto">
							<table class="table w-full">
								<thead>
									<tr>
										<th>ID</th>
										<th>Изображение</th>
										<th>Название</th>
										<th>Цена</th>
										<th>Остаток</th>
										<th>Действия</th>
									</tr>
								</thead>
								<tbody>
									{#each products as product}
										<tr>
											<td>{product.ID}</td>
											<td>
												<img
													src={product.image_url}
													alt={product.name}
													class="w-16 h-16 object-cover rounded"
												/>
											</td>
											<td>{product.name}</td>
											<td>${product.price}</td>
											<td>{product.stock}</td>
											<td class="flex gap-2">
												<button class="btn btn-sm btn-info" on:click={() => handleEdit(product)}>
													Редактировать
												</button>
												<button
													class="btn btn-sm btn-error"
													on:click={() => handleDelete(product.ID)}
												>
													Удалить
												</button>
											</td>
										</tr>
									{/each}
								</tbody>
							</table>
						</div>
					{/if}
				</div>
			{/if}
		</div>
	</div>

	<div class="drawer-side">
		<label for="admin-drawer" class="drawer-overlay"></label>
		<ul class="menu p-4 w-60 h-full bg-base-200 text-base-content">
			<h2 class="text-lg font-bold mb-4">Админ панель</h2>
			<li class="mb-2">
				<button
					class:active={activeTab === 'orders'}
					on:click|preventDefault={() => switchTab('orders')}
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
						/>
					</svg>
					Заказы
				</button>
			</li>
			<li>
				<button
					class:active={activeTab === 'products'}
					on:click|preventDefault={() => switchTab('products')}
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"
						/>
					</svg>
					Товары
				</button>
			</li>
		</ul>
	</div>
</div>

{#if showEditModal}
	<div class="modal modal-open">
		<div class="modal-box">
			<h3 class="font-bold text-lg mb-4">Редактировать товар</h3>
			<form on:submit|preventDefault={saveProduct} class="space-y-4">
				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Название</label>
					<input
						type="text"
						bind:value={editingProduct.name}
						class="input input-bordered"
						required
					/>
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Цена</label>
					<input
						type="number"
						bind:value={editingProduct.price}
						class="input input-bordered"
						required
						min="0"
						step="0.01"
					/>
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Остаток</label>
					<input
						type="number"
						bind:value={editingProduct.stock}
						class="input input-bordered"
						required
						min="0"
					/>
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Описание</label>
					<textarea
						bind:value={editingProduct.description}
						class="textarea textarea-bordered"
						rows="3"
					></textarea>
				</div>

				<div class="modal-action">
					<button type="submit" class="btn btn-primary">Сохранить</button>
					<button type="button" class="btn" on:click={() => (showEditModal = false)}>Отмена</button>
				</div>
			</form>
		</div>
	</div>
{/if}

{#if showAddModal}
	<div class="modal modal-open">
		<div class="modal-box">
			<h3 class="font-bold text-lg mb-4">Добавить новый товар</h3>
			<form on:submit|preventDefault={handleAddProduct} class="space-y-4">
				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Название</label>
					<input type="text" bind:value={newProduct.name} class="input input-bordered" required />
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Цена</label>
					<input
						type="number"
						bind:value={newProduct.price}
						class="input input-bordered"
						required
						min="0"
						step="0.01"
					/>
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Остаток</label>
					<input
						type="number"
						bind:value={newProduct.stock}
						class="input input-bordered"
						required
						min="0"
					/>
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">URL изображения</label>
					<input
						type="url"
						bind:value={newProduct.image_url}
						class="input input-bordered"
						required
					/>
				</div>

				<div class="form-control">
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label class="label">Описание</label>
					<textarea bind:value={newProduct.description} class="textarea textarea-bordered" rows="3"
					></textarea>
				</div>

				<div class="modal-action">
					<button type="submit" class="btn btn-primary">Добавить</button>
					<button type="button" class="btn" on:click={() => (showAddModal = false)}>
						Отмена
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
