import { skeleton } from '@skeletonlabs/tw-plugin';
import { join } from 'path';


/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		join(require.resolve(
			'@skeletonlabs/skeleton'),
			'../**/*.{html,js,svelte,ts}'
		)
	],

  theme: {
    extend: {},
  },
  plugins: [skeleton({
		themes: { preset: [ "skeleton", "crimson", 'gold-nouveau', "hamlindigo", "modern", "rocket", "sahara", "seafoam", "skeleton" ] }
	})],
}

