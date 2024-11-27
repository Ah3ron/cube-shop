<script>
	import { onMount } from 'svelte';
	import { cartApi } from '$lib/api/cart.js';

	let cart = { items: [], total: 0 };

	let loading = true;
	let error = null;

	async function loadCart() {
		try {
			loading = true;
			cart = await cartApi.get();
			cart.items = cart.items.sort((a, b) => a.id - b.id);
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	async function updateQuantity(id, quantity) {
		try {
			await cartApi.updateItem(id, quantity);
			await loadCart();
		} catch (err) {
			error = err.message;
		}
	}

	async function removeItem(id) {
		try {
			await cartApi.removeItem(id);
			await loadCart();
		} catch (err) {
			error = err.message;
		}
	}

	onMount(loadCart);
</script>

<div class="container mx-auto px-4 py-8 max-w-4xl">
	{#if loading}
		<div class="flex justify-center items-center h-64">
			<span class="loading loading-spinner text-primary text-2xl"></span>
		</div>
	{:else if error}
		<div class="bg-error/10 text-error p-4 rounded-lg text-center shadow-sm">
			{error}
		</div>
	{:else if cart.items.length === 0}
		<div class="text-center py-16 bg-base-200 rounded-lg">
			<h2 class="text-2xl font-semibold text-base-content/70">Корзина пуста</h2>
			<a href="/catalog" class="btn btn-primary mt-4">Перейти в каталог</a>
		</div>
	{:else}
		<div class="space-y-6">
			<h1 class="text-3xl font-bold mb-8">Корзина</h1>

			<div class="space-y-4">
				{#each cart.items as item}
					<div class="card bg-base-100 shadow-lg hover:shadow-xl transition-shadow duration-200">
						<div class="p-6 flex justify-between items-center gap-4">
							<div class="flex items-center gap-6">
								<img
									src={item.product.image_url}
									alt={item.product.name}
									class="w-24 h-24 object-cover rounded-lg shadow-sm"
								/>
								<div class="flex-1">
									<h3 class="text-lg font-bold mb-2">{item.product.name}</h3>
									<p class="text-base-content/60">
										{item.product.price.toLocaleString()} $ × {item.quantity}
									</p>
								</div>
							</div>

							<div class="flex items-center gap-4">
								<div class="flex items-center bg-base-200 rounded-lg p-1">
									<button
										class="btn btn-primary btn-sm px-3"
										disabled={item.quantity <= 1}
										on:click={() => updateQuantity(item.id, item.quantity - 1)}
									>
										−
									</button>
									<span class="w-8 text-center font-medium">{item.quantity}</span>
									<button
										class="btn btn-primary btn-sm px-3"
										on:click={() => updateQuantity(item.id, item.quantity + 1)}
									>
										+
									</button>
								</div>
								<button class="btn btn-error btn-sm" on:click={() => removeItem(item.id)}>
									Удалить
								</button>
							</div>
						</div>
					</div>
				{/each}
			</div>

			<div class="card bg-base-100 shadow-lg p-6 mt-8">
				<div class="flex justify-between items-center">
					<span class="text-xl">Итого:</span>
					<span class="text-2xl font-bold">
						{cart.total.toLocaleString()} $
					</span>
				</div>
				<hr />
				<div class="flex justify-end mt-6">
					<a href="/checkout" class="btn btn-primary btn-lg"> Оформить заказ </a>
				</div>
			</div>
		</div>
	{/if}
</div>

<style>
</style>
