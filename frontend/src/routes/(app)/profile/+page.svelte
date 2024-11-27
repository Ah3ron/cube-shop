<script>
	import { onMount } from 'svelte';
	import { userApi } from '$lib/api/user.js';

	let profile = {
		created_at: null,
		email: null,
		id: null,
		name: null
	};

	let isLoading = true;
	let error = null;

	onMount(async () => {
		try {
			profile = await userApi.getProfile();
		} catch (err) {
			error = err;
		} finally {
			isLoading = false;
		}
	});
</script>

<div class="container mx-auto px-4 py-8">
	<h1 class="text-3xl font-bold mb-8">Profile</h1>

	{#if isLoading}
		<div class="flex justify-center">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else if error}
		<div class="alert alert-error">
			<span>{error}</span>
		</div>
	{:else}
		<div class="card bg-base-100 shadow-xl">
			<div class="card-body">
				<div class="flex items-center gap-6">
					<div class="avatar placeholder">
						<div class="bg-neutral text-neutral-content rounded-full w-24">
							<span class="text-3xl">{profile.name?.[0]?.toUpperCase() ?? '?'}</span>
						</div>
					</div>

					<div class="space-y-2">
						<h2 class="text-2xl font-semibold">{profile.name}</h2>
						<p class="text-base-content/70">{profile.email}</p>
						<p class="text-sm">Member since: {new Date(profile.created_at).toLocaleDateString()}</p>
					</div>
				</div>

				<div class="divider"></div>

				<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
					<div class="card bg-base-200">
						<div class="card-body">
							<h3 class="card-title">Recent Orders</h3>
							<p>View your order history in the Orders section</p>
							<div class="card-actions justify-end">
								<a href="/orders" class="btn btn-primary">View Orders</a>
							</div>
						</div>
					</div>

					<div class="card bg-base-200">
						<div class="card-body">
							<h3 class="card-title">Account Settings</h3>
							<p>Update your profile and preferences</p>
							<div class="card-actions justify-end">
								<a href="/settings" class="btn btn-primary">Settings</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
