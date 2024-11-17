/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {}
	},

	plugins: [require('daisyui')],
	daisyui: {
		themes: [
			{
				myTheme: {
					primary: '#dea19f',
					secondary: '#8a2521',
					accent: '#ff80b7',
					neutral: '#262626',
					'base-100': '#171717'
				}
			}
		]
	}
};
