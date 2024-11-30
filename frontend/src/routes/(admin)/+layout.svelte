<script>
	import '../../app.css';
	import { onMount } from 'svelte';

	import { userApi } from '$lib/api/user.js';

	let isAdmin = false;

	onMount(async () => {
		const token = localStorage.getItem('token');
		if (!token) {
			window.location.href = '/login';
			return;
		}

		try {
			const data = await userApi.getProfile();
			if (!data.role || data.role !== 'admin') {
				window.location.href = '/login';
			} else {
				isAdmin = true;
			}
		} catch (error) {
			console.log(error);

			window.location.href = '/login';
		}
	});
</script>

{#if isAdmin}
	<slot />
{/if}
