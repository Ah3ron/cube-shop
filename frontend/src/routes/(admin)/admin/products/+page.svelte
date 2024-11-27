<script>
	import { onMount } from 'svelte';

	let products = [];
	let loading = true;
	let error = null;

	let editingProduct = null;
	let showEditModal = false;

	onMount(async () => {
		try {
			const response = await fetch('/api/v1/products');
			products = await response.json();
		} catch (err) {
			console.log(err.message);
		} finally {
			loading = false;
		}
	});

	async function handleDelete(id) {
		if (!confirm('Вы уверены, что хотите удалить этот товар?')) return;

		try {
			await fetch(`/api/v1/products/${id}`, {
				method: 'DELETE',
				headers: {
					Authorization: `Bearer ${localStorage.getItem('token')}`
				}
			});
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
			const response = await fetch(`/api/v1/products/${editingProduct.ID}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${localStorage.getItem('token')}`
				},
				body: JSON.stringify(editingProduct)
			});

			const updatedProduct = await response.json();
			products = products.map((p) => (p.ID === updatedProduct.ID ? updatedProduct : p));
			showEditModal = false;
		} catch (err) {
			alert('Ошибка при обновлении товара: ' + err.message);
		}
	}
	let showAddModal = false;
	let newProduct = {
		name: '',
		price: 0,
		stock: 0,
		description: '',
		image_url: ''
	};
	async function handleAddProduct() {
		try {
			const response = await fetch('/api/v1/products', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${localStorage.getItem('token')}`
				},
				body: JSON.stringify(newProduct)
			});

			const addedProduct = await response.json();
			products = [...products, addedProduct];
			showAddModal = false;
			// Reset the form
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

<div class="p-4 bg-base-100 rounded-box shadow-md">
	<div class="flex justify-between items-center mb-6">
		<h1 class="text-2xl font-bold">Управление товарами</h1>
		<div class="flex justify-between items-center mb-6">
			<h1 class="text-2xl font-bold">Управление товарами</h1>
			<button class="btn btn-primary" on:click={() => (showAddModal = true)}>Добавить товар</button>
		</div>
	</div>

	{#if loading}
		<div class="flex justify-center">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error">{error}</div>
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
								<button class="btn btn-sm btn-error" on:click={() => handleDelete(product.ID)}>
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
