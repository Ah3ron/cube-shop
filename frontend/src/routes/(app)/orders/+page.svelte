<script>
	import { onMount } from 'svelte';
	import { ordersApi } from '$lib/api/orders.js';

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

			orders = await ordersApi.getAll();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});

	function formatDate(dateString) {
		return new Date(dateString).toLocaleDateString('ru-RU', {
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}
</script>

<div class="container mx-auto px-4 py-8 mt-10">
	<h1 class="text-3xl font-bold mb-8">Мои заказы</h1>

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
			<p class="text-xl mb-4">У вас пока нет заказов</p>
			<a href="/catalog" class="btn btn-primary">Перейти в каталог</a>
		</div>
	{:else}
		<div class="space-y-6">
			{#each orders as order}
				<div class="card bg-base-100 shadow-xl">
					<div class="card-body">
						<div
							class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4"
						>
							<div>
								<h2 class="card-title">Заказ №{order.ID}</h2>
								<p class="text-base-content/70">Оформлен {formatDate(order.CreatedAt)}</p>
							</div>
							<div
								class="badge badge-lg"
								class:badge-success={order.status === 'completed'}
								class:badge-warning={order.status === 'pending'}
								class:badge-error={order.status === 'cancelled'}
							>
								{order.status === 'completed'
									? 'Выполнен'
									: order.status === 'pending'
										? 'В обработке'
										: 'Отменён'}
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
											<p class="text-sm text-base-content/70">Количество: {item.quantity}</p>
										</div>
									</div>
									<p class="font-medium">{(item.price * item.quantity).toLocaleString()}$</p>
								</div>
							{/each}
						</div>

						<div class="divider"></div>

						<div class="flex flex-col md:flex-row justify-between gap-4">
							<div>
								<h3 class="font-medium mb-2">Данные доставки</h3>
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
									<p class="text-sm text-base-content/70">Данные о доставке отсутствуют</p>
								{/if}
							</div>
							<div class="text-right">
								<p class="text-sm mb-1">Итого:</p>
								<p class="text-2xl font-bold">{order.total_price.toLocaleString()} $</p>
							</div>
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
