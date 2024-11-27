<script>
	import { onMount } from 'svelte';
	import { ordersApi } from '$lib/api/orders.js';

	let orders = [];
	let loading = true;
	let error = null;

	onMount(async () => {
		try {
			orders = await ordersApi.getAll();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});

	async function updateStatus(orderId, newStatus) {
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
</script>

<div class="p-4 bg-base-100 rounded-box shadow-md">
	<h1 class="text-2xl font-bold mb-6">Управление заказами</h1>

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
									on:change={(e) => updateStatus(order.ID, e.target.value)}
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
