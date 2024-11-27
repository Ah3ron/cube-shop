<script>
	import '../../app.css';
	let { children } = $props();
	import { page } from '$app/stores';
	import { cart } from '$lib/stores/cart';
	import { onMount } from 'svelte';

	onMount(async () => {
		await cart.fetch();
	});

	let cartCount = $derived($cart.items.reduce((acc, item) => acc + item.quantity, 0));

	let isAuthenticated = $state(false);

	function handleLogout() {
		localStorage.removeItem('token');
		window.location.href = '/login';
	}

	$effect(() => {
		if (typeof window !== 'undefined') {
			isAuthenticated = !!localStorage.getItem('token');
		}
	});
</script>

<svelte:head>
	<title>CubeShop - Your One-Stop Cube Store</title>
	<meta
		name="description"
		content="Browse our collection of cubes and puzzles for all skill levels. Fast worldwide shipping and expert advice from our community."
	/>
</svelte:head>

<div class="flex flex-col">
	<!-- Header -->
	<header class="navbar bg-base-100 shadow-lg fixed top-0 z-50">
		<div class="navbar-start">
			<div class="dropdown">
				<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
				<!-- svelte-ignore a11y_label_has_associated_control -->
				<label tabindex="0" class="btn btn-ghost lg:hidden">
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
							d="M4 6h16M4 12h8m-8 6h16"
						/>
					</svg>
				</label>
				<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
				<ul
					tabindex="0"
					class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					<li><a href="/" class:active={$page.url.pathname === '/'}>Home</a></li>
					<li><a href="/catalog" class:active={$page.url.pathname === '/catalog'}>Catalog</a></li>
					<li><a href="/about" class:active={$page.url.pathname === '/about'}>About</a></li>
				</ul>
			</div>
			<a href="/" class="btn btn-ghost normal-case text-xl">CubeShop</a>
		</div>

		<div class="navbar-center hidden lg:flex">
			<ul class="menu menu-horizontal px-1">
				<li><a href="/" class:active={$page.url.pathname === '/'}>Home</a></li>
				<li><a href="/catalog" class:active={$page.url.pathname === '/catalog'}>Catalog</a></li>
				<li><a href="/about" class:active={$page.url.pathname === '/about'}>About</a></li>
			</ul>
		</div>

		<div class="navbar-end">
			{#if isAuthenticated}
				<div class="dropdown dropdown-end mr-2">
					<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label tabindex="0" class="btn btn-ghost btn-circle">
						<div class="indicator">
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
									d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
								/>
							</svg>
							{#if cartCount > 0}
								<span class="badge badge-sm indicator-item badge-primary">
									{cartCount}
								</span>
							{/if}
						</div>
					</label>
					<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
					<div
						tabindex="0"
						class="mt-3 z-[1] card card-compact dropdown-content w-52 bg-base-100 shadow"
					>
						<div class="card-body">
							<span class="font-bold text-lg">{cartCount || 0} Items</span>
							<span class="text-info">Subtotal: ${$cart.total.toFixed(2)}</span>
							<div class="card-actions">
								<a href="/cart" class="btn btn-primary btn-block">View cart</a>
							</div>
						</div>
					</div>
				</div>

				<div class="dropdown dropdown-end">
					<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
					<!-- svelte-ignore a11y_label_has_associated_control -->
					<label tabindex="0" class="btn btn-ghost btn-circle avatar placeholder">
						<div class="w-10 rounded-full flex items-center justify-center">
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
									d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
								/>
							</svg>
						</div>
					</label>
					<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
					<ul
						tabindex="0"
						class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
					>
						<li><a href="/profile">Profile</a></li>
						<li><a href="/orders">Orders</a></li>
						<li><a href="/settings">Settings</a></li>
						<li><button onclick={handleLogout}>Logout</button></li>
					</ul>
				</div>
			{:else}
				<a href='/login" class="btn btn-primary ml-4">Login</a>
			{/if}
		</div>
	</header>

	<!-- Main content -->
	<main class="min-h-screen flex-grow pt-16 bg-base-200">
		{@render children()}
	</main>

	<!-- Footer -->
	<footer class="bg-base-200">
		<!-- Main Footer -->
		<div class="footer container mx-auto px-4 py-10">
			<div>
				<span class="footer-title">CubeShop</span>
				<p class="max-w-xs mt-2 text-base-content/70">
					Your one-stop destination for all types of cubes and puzzles. Providing reliable cube
					puzzles since 2024.
				</p>
				<div class="flex gap-4 mt-4">
					<a
						href="https://twitter.com"
						class="btn btn-ghost btn-sm btn-circle"
						aria-label="Twitter"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="20"
							height="20"
							viewBox="0 0 24 24"
							class="fill-current"
						>
							<path
								d="M24 4.557c-.883.392-1.832.656-2.828.775 1.017-.609 1.798-1.574 2.165-2.724-.951.564-2.005.974-3.127 1.195-.897-.957-2.178-1.555-3.594-1.555-3.179 0-5.515 2.966-4.797 6.045-4.091-.205-7.719-2.165-10.148-5.144-1.29 2.213-.669 5.108 1.523 6.574-.806-.026-1.566-.247-2.229-.616-.054 2.281 1.581 4.415 3.949 4.89-.693.188-1.452.232-2.224.084.626 1.956 2.444 3.379 4.6 3.419-2.07 1.623-4.678 2.348-7.29 2.04 2.179 1.397 4.768 2.212 7.548 2.212 9.142 0 14.307-7.721 13.995-14.646.962-.695 1.797-1.562 2.457-2.549z"
							></path>
						</svg>
					</a>
					<a
						href="https://youtube.com"
						class="btn btn-ghost btn-sm btn-circle"
						aria-label="YouTube"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="20"
							height="20"
							viewBox="0 0 24 24"
							class="fill-current"
						>
							<path
								d="M19.615 3.184c-3.604-.246-11.631-.245-15.23 0-3.897.266-4.356 2.62-4.385 8.816.029 6.185.484 8.549 4.385 8.816 3.6.245 11.626.246 15.23 0 3.897-.266 4.356-2.62 4.385-8.816-.029-6.185-.484-8.549-4.385-8.816zm-10.615 12.816v-8l8 3.993-8 4.007z"
							></path>
						</svg>
					</a>
					<a
						href="https://facebook.com"
						class="btn btn-ghost btn-sm btn-circle"
						aria-label="Facebook"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="20"
							height="20"
							viewBox="0 0 24 24"
							class="fill-current"
						>
							<path
								d="M9 8h-3v4h3v12h5v-12h3.642l.358-4h-4v-1.667c0-.955.192-1.333 1.115-1.333h2.885v-5h-3.808c-3.596 0-5.192 1.583-5.192 4.615v3.385z"
							></path>
						</svg>
					</a>
				</div>
			</div>

			<div>
				<span class="footer-title">Quick Links</span>
				<a href="/catalog" class="link link-hover">Catalog</a>
				<a href="/about" class="link link-hover">About Us</a>
				<a href="/contact" class="link link-hover">Contact</a>
			</div>

			<div class="w-full md:w-80">
				<span class="footer-title">Newsletter</span>
				<div class="form-control w-full">
					<div class="relative">
						<input
							type="email"
							placeholder="your@email.com"
							class="input input-bordered w-full pr-16"
						/>
						<button class="btn btn-primary absolute top-0 right-0 rounded-l-none">
							Subscribe
						</button>
					</div>
				</div>
			</div>
		</div>

		<!-- Copyright -->
		<div class="border-t border-base-300">
			<div class="container mx-auto px-4 py-4 text-sm text-base-content/70">
				© 2024 CubeShop. All rights reserved.
			</div>
		</div>
	</footer>
</div>
