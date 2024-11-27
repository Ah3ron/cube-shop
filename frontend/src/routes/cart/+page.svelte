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

<div class="container mx-auto px-4 py-8">
	{#if loading}
		<div class="loading">Загрузка...</div>
	{:else if error}
		<div class="error">{error}</div>
	{:else if cart.items.length === 0}
		<div class="empty">Корзина пуста</div>
	{:else}
		<div class="space-y-4">
			{#each cart.items as item}
				<div class="card p-4 flex justify-between items-center">
					<img src={item.product.imageUrl} alt={item.product.name} class="w-16 h-16 object-cover" />
					<div class="flex-1 px-4">
						<h3 class="font-bold">{item.product.name}</h3>
						<p class="text-gray-600">{item.product.price} ₽ x {item.quantity}</p>
					</div>
					<div class="flex items-center space-x-2">
						<button class="btn btn-sm" on:click={() => updateQuantity(item.id, item.quantity - 1)}>-</button>
						<span>{item.quantity}</span>
						<button class="btn btn-sm" on:click={() => updateQuantity(item.id, item.quantity + 1)}>+</button>
						<button class="btn btn-error btn-sm" on:click={() => removeItem(item.id)}>Удалить</button>
					</div>
				</div>
			{/each}
			<div class="text-right text-xl font-bold">
				Итого: {cart.total} ₽
			</div>
			<div class="text-right">
				<a href="/checkout" class="btn btn-primary">Оформить заказ</a>
			</div>
		</div>
	{/if}
</div>
