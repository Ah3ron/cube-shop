<script>
	import { fade, fly, scale } from 'svelte/transition';
	import { onMount } from 'svelte';

	let visible = $state(false);

	onMount(() => {
		visible = true;
	});
</script>

{#if visible}
	<div class="hero min-h-screen bg-base-200 relative" in:fade={{ duration: 600 }}>
		<div class="hero-content text-center w-full max-w-[1200px]">
			<div class="max-w-md h-[300px]">
				<h1 class="text-5xl font-bold" in:fly={{ y: 30, duration: 600 }}>CubeShop</h1>
				<p class="py-6" in:fly={{ y: 30, duration: 600, delay: 100 }}>
					Your one-stop destination for all types of cubes and puzzles
				</p>
				<a
					href="/catalog"
					class="btn btn-primary hover:scale-105 transition-transform duration-200"
					in:fly={{ y: 30, duration: 600, delay: 200 }}
				>
					Browse Catalog
				</a>
			</div>
		</div>
	</div>
	<div
		class="absolute bottom-8 left-1/2 -translate-x-1/2 text-center"
		in:fade={{ duration: 600, delay: 1500 }}
	>
		<div class="flex flex-col items-center gap-2">
			<span class="text-sm opacity-70">Scroll to explore</span>
			<div class="scroll-arrow">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="24"
					height="24"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					class="animate-bounce"
				>
					<path d="M12 5v14M19 12l-7 7-7-7" />
				</svg>
			</div>
		</div>
	</div>

	<div class="container mx-auto px-4 py-16" in:fade={{ duration: 600, delay: 300 }}>
		<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
			{#each Array(3) as _, i}
				<div
					class="card bg-base-100 shadow-xl hover:shadow-2xl transition-all duration-300 hover:-translate-y-2"
					in:fly={{ x: -30, duration: 500, delay: 400 + i * 100 }}
				>
					<div class="card-body">
						<h2 class="card-title">
							{['Wide Selection', 'Expert Advice', 'Fast Shipping'][i]}
						</h2>
						<p>
							{[
								'Explore our vast collection of puzzles for all skill levels',
								'Get guidance from our experienced cubing community',
								'Quick delivery to your doorstep worldwide'
							][i]}
						</p>
					</div>
				</div>
			{/each}
		</div>
	</div>

	<div class="bg-base-200 py-16" in:fade={{ duration: 600, delay: 700 }}>
		<div class="container mx-auto px-4">
			<h2 class="text-3xl font-bold text-center mb-8">Featured Categories</h2>
			<div class="grid grid-cols-2 md:grid-cols-4 gap-4">
				{#each ['3x3 Cubes', 'Speed Cubes', 'Puzzle Sets', 'Accessories'] as category, i}
					<div
						class="card bg-base-100 hover:scale-105 transition-all duration-300"
						in:scale={{ duration: 400, delay: 800 + i * 100 }}
					>
						<figure class="px-4 pt-4">
							<div class="w-32 h-32 bg-primary/10 rounded-xl flex items-center justify-center">
								<span class="text-4xl">{['🎲', '🔄', '🧩', '🎯'][i]}</span>
							</div>
						</figure>
						<div class="card-body items-center text-center">
							<h3 class="card-title">{category}</h3>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<div class="bg-base-100 py-16" in:fade={{ duration: 600, delay: 1000 }}>
		<div class="container mx-auto px-4">
			<div class="flex flex-col md:flex-row items-center place-content-between gap-8">
				<div class="flex-1" in:fly={{ x: -30, duration: 600, delay: 1100 }}>
					<h2 class="text-3xl font-bold mb-4">Join Our Community</h2>
					<p class="mb-4">
						Connect with fellow cube enthusiasts, share your achievements, and learn new techniques.
					</p>
					<button class="btn btn-primary hover:scale-105 transition-transform duration-200">
						Sign Up Now
					</button>
				</div>
				<div class="flex-1 max-w-min" in:fly={{ x: 30, duration: 600, delay: 1200 }}>
					<div class="bg-base-300 stats shadow hover:shadow-xl transition-all duration-300">
						<div class="stat">
							<div class="stat-title">Happy Customers</div>
							<div class="stat-value">10K+</div>
						</div>
						<div class="stat">
							<div class="stat-title">Products</div>
							<div class="stat-value">500+</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}

<style>
</style>
