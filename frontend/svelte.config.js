import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: [vitePreprocess({})],

  kit: {
    adapter: adapter({
      pages: 'output/build',
      assets: 'output/build',
      fallback: 'index.html', // undefined
      precompress: false,
      strict: true
    }),
    paths: {
      base: "/z/pages"
    },
    // version: {
    //   name: "0.1.12",
    // }
  },
};

export default config;
