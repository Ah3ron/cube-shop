<script>
	import { onMount } from 'svelte';
	import { cart } from '$lib/stores/cart';

	let loading = true;
	let error = null;

	async function handleCheckout() {
		const token = localStorage.getItem('token');
		if (!token) {
			window.location.href = '/auth/login';
			return;
		}

		try {
			const response = await fetch('/api/orders', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({
					items: $cart,
					total: $cart.reduce((sum, item) => sum + item.price * item.quantity, 0)
				})
			});

			if (!response.ok) throw new Error('Checkout failed');
			window.location.href = '/orders';
		} catch (err) {
			error = err.message;
		}
	}

	onMount(async () => {
		try {
			await cart.fetch();
		} finally {
			loading = false;
		}
	});
</script>

<div class="container mx-auto px-4 py-8 mt-16">
	<h1 class="text-3xl font-bold mb-8">Shopping Cart</h1>

	{#if loading}
		<div class="flex justify-center">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error">
			<span>{error}</span>
		</div>
	{:else if $cart.length === 0}
		<div class="text-center py-8">
			<p class="text-xl mb-4">Your cart is empty</p>
			<a href="/catalog" class="btn btn-primary">Continue Shopping</a>
		</div>
	{:else}
		<div class="flex flex-col lg:flex-row gap-8">
			<div class="flex-1">
				{#each $cart as item}
					<div class="card bg-base-100 shadow-xl mb-4">
						<div class="card-body flex flex-row">
							<div class="flex-1">
								<h2 class="card-title">{item.product.name}</h2>
								<p class="text-base-content/70">${item.product.price}</p>
							</div>
							<div class="flex items-center gap-4">
								<div class="join">
									<button
										class="btn btn-primary btn-sm join-item"
										on:click={() => cart.updateQuantity(item.id, item.quantity - 1)}>-</button
									>
									<input
										type="number"
										class="input input-bordered input-sm w-16 join-item"
										value={item.quantity}
										on:change={(e) => cart.updateQuantity(item.id, parseInt(e.target.value))}
									/>
									<button
										class="btn btn-sm join-item btn-primary"
										on:click={() => cart.updateQuantity(item.id, item.quantity + 1)}>+</button
									>
								</div>
								<!-- svelte-ignore a11y_consider_explicit_label -->
								<button
									class="btn btn-ghost btn-sm hover:btn-error hover:btn hover:btn-sm"
									on:click={() => cart.removeItem(item.id)}
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
											d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
										/>
									</svg>
								</button>
							</div>
						</div>
					</div>
				{/each}
			</div>

			<div class="w-full lg:w-80">
				<div class="card bg-base-100 shadow-xl">
					<div class="card-body">
						<h2 class="card-title">Order Summary</h2>
						<div class="text-lg">
							Total: ${$cart
								.reduce((sum, item) => sum + item.product.price * item.quantity, 0)
								.toFixed(2)}
						</div>
						<button class="btn btn-primary mt-4" on:click={handleCheckout}>
							Proceed to Checkout
						</button>
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
