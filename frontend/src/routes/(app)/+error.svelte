<script>
	import { page } from '$app/stores';
	import { fade, fly } from 'svelte/transition';
</script>

<div
	class="min-h-screen flex items-center justify-center bg-base-200 px-4"
	in:fade={{ duration: 300 }}
>
	<div class="text-center" in:fly={{ y: 50, duration: 500, delay: 300 }}>
		<h1 class="text-9xl font-bold text-primary mb-4">404</h1>
		<p class="text-2xl mb-8">Страница не найдена</p>
		<p class="text-base-content/70 mb-8">
			Извините, страница "{$page.url.pathname}" не существует.
		</p>
		<a href="/" class="btn btn-primary" in:fly={{ y: 20, duration: 500, delay: 500 }}>
			Вернуться на главную
		</a>
	</div>
</div>
