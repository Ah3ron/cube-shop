import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/api': {
				target: 'https://cube-shop.up.railway.app/',
				changeOrigin: true
			}
		}
	},
	build: {
		minify: 'terser',
		terserOptions: {
			compress: {
				drop_console: true,
				drop_debugger: true
			},
			format: {
				comments: false
			}
		},
		rollupOptions: {
			output: {
				manualChunks: undefined
			}
		}
	}
});
