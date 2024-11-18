import adapter from '@sveltejs/adapter-static';

export default {
	kit: {
		adapter: adapter({
			// default options are shown. On some platforms
			// these options are set automatically — see below
			pages: 'build',
			assets: 'build',
			fallback: undefined,
			precompress: false,
			strict: true
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
				if (
					[
						'/new',
						'/bestsellers',
						'/sale',
						'/cookie',
						'/contact',
						'/jobs',
						'/terms',
						'/privacy'
					].includes(path)
				) {
					return;
				}

				// otherwise fail the build
				throw new Error(message);
			}
		}
	}
};
