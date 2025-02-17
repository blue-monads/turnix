import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	optimizeDeps: {
		exclude: [
			// "@codemirror/lang-javascript", 
			// "@codemirror/lang-html", 
			// "@codemirror/lang-sql",
			// "@codemirror/lang-markdown",
		],
    },
});
