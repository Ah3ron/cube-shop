<script>
	import '../app.css';
	let { children } = $props();
	import { page } from '$app/stores';

	let isAuthenticated = $state(false);

	function handleLogout() {
		localStorage.removeItem('token');
		window.location.href = '/auth/login';
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

<div class="min-h-screen flex flex-col">
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
				<a href="/auth/login" class="btn btn-primary ml-4">Login</a>
			{/if}
		</div>
	</header>

	<!-- Main content -->
	<main class="flex-grow pt-16">
		{@render children()}
	</main>

	<!-- Footer -->
	<footer class="footer px-12 py-16 bg-base-200 text-base-content">
		<div>
			<span class="footer-title">Shop</span>
			<a href="/catalog" class="link link-hover">Catalog</a>
			<a href="/new" class="link link-hover">New Arrivals</a>
			<a href="/bestsellers" class="link link-hover">Bestsellers</a>
			<a href="/sale" class="link link-hover">Sale</a>
		</div>
		<div>
			<span class="footer-title">Company</span>
			<a href="/about" class="link link-hover">About us</a>
			<a href="/contact" class="link link-hover">Contact</a>
			<a href="/jobs" class="link link-hover">Jobs</a>
		</div>
		<div>
			<span class="footer-title">Legal</span>
			<a href="/terms" class="link link-hover">Terms of use</a>
			<a href="/privacy" class="link link-hover">Privacy policy</a>
			<a href="/cookie" class="link link-hover">Cookie policy</a>
		</div>
		<div>
			<span class="footer-title">Newsletter</span>
			<div class="form-control w-80">
				<!-- svelte-ignore a11y_label_has_associated_control -->
				<label class="label">
					<span class="label-text">Enter your email address</span>
				</label>
				<div class="relative">
					<input
						type="text"
						placeholder="username@site.com"
						class="input input-bordered w-full pr-16"
					/>
					<button class="btn btn-primary absolute top-0 right-0 rounded-l-none">Subscribe</button>
				</div>
			</div>
		</div>
	</footer>
	<footer class="footer px-14 py-4 border-t bg-base-200 text-base-content border-base-300">
		<div class="items-center grid-flow-col">
			<svg
				width="24"
				height="24"
				viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg"
				fill-rule="evenodd"
				clip-rule="evenodd"
				class="fill-current"
			>
				<path
					d="M22.672 15.226l-2.432.811.841 2.515c.33 1.019-.209 2.127-1.23 2.456-1.15.325-2.148-.321-2.463-1.226l-.84-2.518-5.013 1.677.84 2.517c.391 1.203-.434 2.542-1.831 2.542-.88 0-1.601-.564-1.86-1.314l-.842-2.516-2.431.809c-1.135.328-2.145-.317-2.463-1.229-.329-1.018.211-2.127 1.231-2.456l2.432-.809-1.621-4.823-2.432.808c-1.355.384-2.558-.59-2.558-1.839 0-.817.509-1.582 1.327-1.846l2.433-.809-.842-2.515c-.33-1.02.211-2.129 1.232-2.458 1.02-.329 2.13.209 2.461 1.229l.842 2.515 5.011-1.677-.839-2.517c-.403-1.238.484-2.553 1.843-2.553.819 0 1.585.509 1.85 1.326l.841 2.517 2.431-.81c1.02-.33 2.131.211 2.461 1.229.332 1.018-.21 2.126-1.23 2.456l-2.433.809 1.622 4.823 2.433-.809c1.242-.401 2.557.484 2.557 1.838 0 .819-.51 1.583-1.328 1.847m-8.992-6.428l-5.01 1.675 1.619 4.828 5.011-1.674-1.62-4.829z"
				></path>
			</svg>
			<p>CubeShop Ltd. <br />Providing reliable cube puzzles since 2024</p>
		</div>
		<div class="md:place-self-center md:justify-self-end">
			<div class="grid grid-flow-col gap-4">
				<a class="link link-hover" href="https://twitter.com" aria-label="Twitter">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						class="fill-current"
					>
						<path
							d="M24 4.557c-.883.392-1.832.656-2.828.775 1.017-.609 1.798-1.574 2.165-2.724-.951.564-2.005.974-3.127 1.195-.897-.957-2.178-1.555-3.594-1.555-3.179 0-5.515 2.966-4.797 6.045-4.091-.205-7.719-2.165-10.148-5.144-1.29 2.213-.669 5.108 1.523 6.574-.806-.026-1.566-.247-2.229-.616-.054 2.281 1.581 4.415 3.949 4.89-.693.188-1.452.232-2.224.084.626 1.956 2.444 3.379 4.6 3.419-2.07 1.623-4.678 2.348-7.29 2.04 2.179 1.397 4.768 2.212 7.548 2.212 9.142 0 14.307-7.721 13.995-14.646.962-.695 1.797-1.562 2.457-2.549z"
						></path>
					</svg>
				</a>
				<a class="link link-hover" href="https://youtube.com" aria-label="YouTube">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						class="fill-current"
					>
						<path
							d="M19.615 3.184c-3.604-.246-11.631-.245-15.23 0-3.897.266-4.356 2.62-4.385 8.816.029 6.185.484 8.549 4.385 8.816 3.6.245 11.626.246 15.23 0 3.897-.266 4.356-2.62 4.385-8.816-.029-6.185-.484-8.549-4.385-8.816zm-10.615 12.816v-8l8 3.993-8 4.007z"
						></path>
					</svg>
				</a>
				<a class="link link-hover" href="https://facebook.com" aria-label="Facebook">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
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
	</footer>
</div>
