<!-- src/routes/+page.svelte -->
<script>
	import { fade, fly, scale } from 'svelte/transition';
	import { onMount } from 'svelte';

	let visible = false;
	onMount(() => {
		visible = true;
	});
</script>

{#if visible}
	<div class="hero min-h-screen bg-base-200" in:fade={{ duration: 1000 }}>
		<div class="hero-content text-center">
			<div class="max-w-md">
				<h1 class="text-5xl font-bold" in:fly={{ y: 50, duration: 1000 }}>CubeShop</h1>
				<p class="py-6" in:fly={{ y: 50, duration: 1000, delay: 200 }}>
					Your one-stop destination for all types of cubes and puzzles
				</p>
				<a
					href="/catalog"
					class="btn btn-primary hover:scale-105 transition-transform duration-200"
					in:fly={{ y: 50, duration: 1000, delay: 400 }}
				>
					Browse Catalog
				</a>
			</div>
		</div>
	</div>

	<div class="container mx-auto px-4 py-16" in:fade={{ duration: 1000, delay: 600 }}>
		<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
			{#each Array(3) as _, i}
				<div
					class="card bg-base-100 shadow-xl hover:shadow-2xl transition-all duration-300 hover:-translate-y-2"
					in:fly={{ x: -50, duration: 800, delay: 800 + i * 200 }}
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

	<div class="bg-base-200 py-16" in:fade={{ duration: 1000, delay: 1400 }}>
		<div class="container mx-auto px-4">
			<h2 class="text-3xl font-bold text-center mb-8">Featured Categories</h2>
			<div class="grid grid-cols-2 md:grid-cols-4 gap-4">
				{#each ['3x3 Cubes', 'Speed Cubes', 'Puzzle Sets', 'Accessories'] as category, i}
					<div
						class="card bg-base-100 hover:scale-105 transition-all duration-300"
						in:scale={{ duration: 800, delay: 1600 + i * 200 }}
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

	<div class="container mx-auto px-4 py-16" in:fade={{ duration: 1000, delay: 2000 }}>
		<div class="flex flex-col md:flex-row items-center place-content-between gap-8">
			<div class="flex-1 max-w-[50%]" in:fly={{ x: -50, duration: 1000, delay: 2200 }}>
				<h2 class="text-3xl font-bold mb-4">Join Our Community</h2>
				<p class="mb-4">
					Connect with fellow cube enthusiasts, share your achievements, and learn new techniques.
				</p>
				<button class="btn btn-primary hover:scale-105 transition-transform duration-200">
					Sign Up Now
				</button>
			</div>
			<div class="flex-1 max-w-min" in:fly={{ x: 50, duration: 1000, delay: 2400 }}>
				<div class="stats shadow hover:shadow-xl transition-all duration-300">
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
{/if}

<style>
	:global(.card) {
		backface-visibility: hidden;
	}
</style>
