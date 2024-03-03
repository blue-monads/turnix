import { skeleton } from '@skeletonlabs/tw-plugin';


/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [skeleton({
		themes: { preset: [ "skeleton" ] }
	})],
}

