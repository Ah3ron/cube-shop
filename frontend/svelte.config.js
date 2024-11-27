import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

export default {
	preprocess: [vitePreprocess()],
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: '404.html',
			precompress: true,
			strict: false
		}),
		csp: {
			mode: 'auto',
			directives: {
				'script-src': ['self'],
				'upgrade-insecure-requests': false,
				'block-all-mixed-content': false
			}
		}
	}
};
