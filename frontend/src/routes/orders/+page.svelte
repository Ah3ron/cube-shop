<script>
	import { onMount } from 'svelte';

	let orders = [];
	let loading = true;
	let error = null;

	onMount(async () => {
		try {
			const token = localStorage.getItem('token');
			if (!token) {
				window.location.href = '/auth/login';
				return;
			}

			const response = await fetch('/api/orders', {
				method: 'GET',
				headers: {
					Authorization: `Bearer ${token}`
				}
			});

			if (!response.ok) {
				throw new Error('Failed to fetch orders');
			}

			orders = await response.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});

	function formatDate(dateString) {
		return new Date(dateString).toLocaleDateString('en-US', {
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}
</script>

<div class="container mx-auto px-4 py-8 mt-10">
	<h1 class="text-3xl font-bold mb-8">My Orders</h1>

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
					d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
				/>
			</svg>
			<span>{error}</span>
		</div>
	{:else if orders.length === 0}
		<div class="text-center py-8">
			<p class="text-xl mb-4">You haven't placed any orders yet</p>
			<a href="/catalog" class="btn btn-primary">Start Shopping</a>
		</div>
	{:else}
		<div class="space-y-6">
			{#each orders as order}
				<div class="card bg-base-100 shadow-xl">
					<div class="card-body">
						<div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
							<div>
								<h2 class="card-title">Order #{order.ID}</h2>
								<p class="text-base-content/70">Placed on {formatDate(order.CreatedAt)}</p>
							</div>
							<div
								class="badge badge-lg"
								class:badge-success={order.Status === 'completed'}
								class:badge-warning={order.Status === 'pending'}
								class:badge-error={order.Status === 'cancelled'}
							>
								{order.Status}
							</div>
						</div>

						<div class="divider"></div>

						<div class="space-y-4">
							{#each order.order_items as item}
								<div class="flex justify-between w-full items-center">
									<div class="flex items-center gap-4 w-full">
										<img
											src={item.product.image_url}
											alt={item.product.name}
											class="w-16 h-16 object-cover rounded-lg"
										/>
										<div>
											<h3 class="font-medium">{item.product.name}</h3>
											<p class="text-sm text-base-content/70">Quantity: {item.quantity}</p>
										</div>
									</div>
									<p class="font-medium">${(item.price * item.quantity).toFixed(2)}</p>
								</div>
							{/each}
						</div>

						<div class="divider"></div>

						<div class="flex flex-col md:flex-row justify-between gap-4">
							<div>
								<h3 class="font-medium mb-2">Shipping Details</h3>
								{#if order.shipping_details?.name}
									<p class="text-sm">{order.shipping_details.name}</p>
									<p class="text-sm">{order.shipping_details.email}</p>
									<p class="text-sm">{order.shipping_details.address}</p>
									<p class="text-sm">
										{order.shipping_details.city},
										{order.shipping_details.country}
										{order.shipping_details.zipCode}
									</p>
								{:else}
									<p class="text-sm text-base-content/70">No shipping details available</p>
								{/if}
							</div>
							<div class="text-right">
								<p class="text-sm mb-1">Order Total:</p>
								<p class="text-2xl font-bold">${order.total_price.toFixed(2)}</p>
							</div>
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>