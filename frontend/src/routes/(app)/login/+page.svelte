<script>
	import { authApi } from '$lib/api/auth.js';

	let email = '';
	let password = '';
	let error = '';

	async function handleLogin(e) {
		e.preventDefault();
		try {
			let data = await authApi.login(email, password);
			localStorage.setItem('token', data.token);
			window.location.href = '/';
		} catch (err) {
			error = err.message;
		}
	}
</script>

<div class="min-h-screen flex items-center justify-center bg-base-200">
	<div class="card w-96 bg-base-100 shadow-xl">
		<div class="card-body">
			<h2 class="card-title justify-center text-2xl font-bold mb-6">Login</h2>

			{#if error}
				<div class="alert alert-error mb-4">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="stroke-current shrink-0 h-6 w-6"
						fill="none"
						viewBox="0 0 24 24"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
						/></svg
					>
					<span>{error}</span>
				</div>
			{/if}

			<form on:submit={handleLogin} class="space-y-4">
				<div class="form-control">
					<label class="label" for="email">
						<span class="label-text">Email</span>
					</label>
					<input type="email" id="email" bind:value={email} class="input input-bordered" required />
				</div>

				<div class="form-control">
					<label class="label" for="password">
						<span class="label-text">Password</span>
					</label>
					<input
						type="password"
						id="password"
						bind:value={password}
						class="input input-bordered"
						required
					/>
				</div>

				<div class="form-control mt-6">
					<button type="submit" class="btn btn-primary">Login</button>
				</div>
			</form>

			<div class="divider">OR</div>

			<div class="text-center">
				<a href="/register" class="link link-hover">Create new account</a>
			</div>
		</div>
	</div>
</div>
