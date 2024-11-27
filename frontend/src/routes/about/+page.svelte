<script>
	import { fade, fly } from 'svelte/transition';
</script>

<div class="container mx-auto px-4 py-10">
	<div class="max-w-4xl mx-auto">
		<div in:fade={{ duration: 400 }}>
			<h1 class="text-4xl font-bold mb-8 text-center">About CubeShop</h1>

			<div class="prose lg:prose-xl mx-auto">
				<div class="card bg-base-100 shadow-xl mb-8" in:fly={{ y: 50, duration: 400 }}>
					<div class="card-body">
						<h2 class="card-title">Our Story</h2>
						<p>
							Founded in 2024, CubeShop has grown from a small puzzle enthusiast's dream into a
							leading online destination for cube puzzles and brain teasers. Our mission is to bring
							joy and challenge to puzzle lovers worldwide.
						</p>
					</div>
				</div>

				<div class="grid md:grid-cols-3 gap-6 mb-8">
					<div class="card bg-base-100 shadow-xl" in:fly={{ y: 50, duration: 400, delay: 100 }}>
						<div class="card-body items-center text-center">
							<div class="text-4xl mb-4">🎯</div>
							<h3 class="card-title">Mission</h3>
							<p>Providing high-quality puzzles for enthusiasts of all skill levels</p>
						</div>
					</div>

					<div class="card bg-base-100 shadow-xl" in:fly={{ y: 50, duration: 400, delay: 200 }}>
						<div class="card-body items-center text-center">
							<div class="text-4xl mb-4">👥</div>
							<h3 class="card-title">Community</h3>
							<p>Building a global community of puzzle enthusiasts</p>
						</div>
					</div>

					<div class="card bg-base-100 shadow-xl" in:fly={{ y: 50, duration: 400, delay: 300 }}>
						<div class="card-body items-center text-center">
							<div class="text-4xl mb-4">💡</div>
							<h3 class="card-title">Innovation</h3>
							<p>Constantly seeking new and exciting puzzle designs</p>
						</div>
					</div>
				</div>

				<div class="card bg-base-100 shadow-xl mb-8" in:fly={{ y: 50, duration: 400, delay: 400 }}>
					<div class="card-body">
						<h2 class="card-title">Our Team</h2>
						<p>
							Our team consists of passionate puzzle enthusiasts who are dedicated to providing the
							best possible experience for our customers. From customer service to product curation,
							we put our heart into everything we do.
						</p>
					</div>
				</div>

				<div class="card bg-base-100 shadow-xl" in:fly={{ y: 50, duration: 400, delay: 500 }}>
					<div class="card-body">
						<h2 class="card-title">Contact Us</h2>
						<div class="space-y-4">
							<div class="flex items-center gap-4">
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
										d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
									/>
								</svg>
								<span>support@cubeshop.com</span>
							</div>
							<div class="flex items-center gap-4">
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
										d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
									/>
								</svg>
								<span>+1 (555) 123-4567</span>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
