import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

export default {
	preprocess: [vitePreprocess()],
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: 'index.html',
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
		},
		prerender: {
			handleHttpError: ({ path, message }) => {
				// ignore deliberate link to shiny 404 page
				if (['/contact'].includes(path)) {
					return;
				}

				// otherwise fail the build
				throw new Error(message);
			}
		}
	}
};
