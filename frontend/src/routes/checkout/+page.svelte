<script>
	import { cart } from '$lib/stores/cart';
	import { ordersApi } from '$lib/api/orders';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	let loading = false;
	let error = null;
	let formData = {
		name: '',
		email: '',
		address: '',
		city: '',
		country: '',
		zipCode: ''
	};

	// Реактивное вычисление общей суммы корзины
	$: cartTotal =
		$cart?.items?.reduce((sum, item) => sum + item.product.price * item.quantity, 0) || 0;

	async function handleSubmit() {
		loading = true;
		error = null;

		try {
			const token = localStorage.getItem('token');
			if (!token) {
				goto('/auth/login');
				return;
			}

			await ordersApi.checkout({
				name: formData.name,
				email: formData.email,
				address: formData.address,
				city: formData.city,
				country: formData.country,
				zipCode: formData.zipCode
			});

			// Очищаем корзину и перенаправляем на страницу заказов
			await cart.clear();
			goto('/orders');
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	function validateZipCode(event) {
		const value = event.target.value;
		if (!/^\d{5}(-\d{4})?$/.test(value)) {
			event.target.setCustomValidity(
				'Введите корректный почтовый индекс (например: 12345 или 12345-6789)'
			);
		} else {
			event.target.setCustomValidity('');
		}
	}
</script>

<div class="container mx-auto px-4 py-8 mt-16">
	<h1 class="text-3xl font-bold mb-8">Secure Checkout</h1>

	{#if $cart.length === 0}
		<div class="text-center py-8">
			<p class="text-xl mb-4">Your cart is empty</p>
			<a href="/catalog" class="btn btn-primary btn-lg">Continue Shopping</a>
		</div>
	{:else}
		<div class="flex flex-col lg:flex-row gap-8">
			<div class="flex-1">
				<div class="card bg-base-100 shadow-xl">
					<div class="card-body">
						<h2 class="card-title mb-4">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="h-6 w-6 mr-2"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
								/>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
								/>
							</svg>
							Shipping Information
						</h2>

						{#if error}
							<div class="alert alert-error mb-4 shadow-lg">
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
										d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
									/>
								</svg>
								<span>{error}</span>
							</div>
						{/if}

						<form on:submit|preventDefault={handleSubmit} class="flex flex-col gap-4">
							<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
								<div class="form-control">
									<label class="label" for="name">
										<span class="label-text">Full Name</span>
									</label>
									<input
										type="text"
										id="name"
										bind:value={formData.name}
										class="input input-bordered focus:input-primary"
										required
										placeholder="John Doe"
									/>
								</div>

								<div class="form-control">
									<label class="label" for="email">
										<span class="label-text">Email</span>
									</label>
									<input
										type="email"
										id="email"
										bind:value={formData.email}
										class="input input-bordered focus:input-primary"
										required
										placeholder="john@example.com"
									/>
								</div>
							</div>

							<div class="form-control">
								<label class="label" for="address">
									<span class="label-text">Address</span>
								</label>
								<!-- svelte-ignore element_invalid_self_closing_tag -->
								<textarea
									id="address"
									bind:value={formData.address}
									class="textarea textarea-bordered focus:textarea-primary h-24"
									required
									placeholder="Street address, apartment, suite, etc."
								/>
							</div>

							<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
								<div class="form-control">
									<label class="label" for="city">
										<span class="label-text">City</span>
									</label>
									<input
										type="text"
										id="city"
										bind:value={formData.city}
										class="input input-bordered focus:input-primary"
										required
									/>
								</div>

								<div class="form-control">
									<label class="label" for="country">
										<span class="label-text">Country</span>
									</label>
									<select
										id="country"
										bind:value={formData.country}
										class="select select-bordered focus:select-primary"
										required
									>
										<option value="">Select country</option>
										<option value="US">United States</option>
										<option value="CA">Canada</option>
										<option value="GB">United Kingdom</option>
									</select>
								</div>

								<div class="form-control">
									<label class="label" for="zipCode">
										<span class="label-text">ZIP Code</span>
									</label>
									<input
										type="text"
										id="zipCode"
										bind:value={formData.zipCode}
										class="input input-bordered focus:input-primary"
										required
										pattern="\d{5}(-\d{4})?"
										on:input={validateZipCode}
										placeholder="12345"
									/>
								</div>
							</div>
						</form>
					</div>
				</div>
			</div>

			<div class="w-full lg:w-96">
				<div class="card bg-base-100 shadow-xl sticky top-4">
					<div class="card-body">
						<h2 class="card-title flex items-center">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="h-6 w-6 mr-2"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
								/>
							</svg>
							Order Summary
						</h2>
						<div class="divider"></div>
						{#each $cart as item}
							<div class="flex justify-between items-center mb-2">
								<div class="flex items-center">
									<span class="font-medium">{item.product.name}</span>
									<span class="text-base-content/70 ml-2">x {item.quantity}</span>
								</div>
								<span class="font-medium">${(item.product.price * item.quantity).toFixed(2)}</span>
							</div>
						{/each}
						<div class="divider"></div>
						<div class="text-lg font-bold flex justify-between">
							<span>Total:</span>
							<span>${cartTotal.toFixed(2)}</span>
						</div>
						<button
							class="btn btn-primary btn-lg mt-4 w-full"
							disabled={loading}
							on:click={handleSubmit}
						>
							{#if loading}
								<span class="loading loading-spinner loading-md"></span>
							{:else}
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-6 w-6 mr-2"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
									/>
								</svg>
							{/if}
							{loading ? 'Processing...' : 'Place Order'}
						</button>
						<p class="text-sm text-center mt-2 text-base-content/70">
							Secure checkout powered by our payment system
						</p>
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
