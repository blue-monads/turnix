import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';



export default defineConfig({
    plugins: [svelte()],
    build: {
        outDir: '../../../../output/dist',
        minify: false,        
        lib: {
            formats: ["iife"],
            entry: './',
            name: "FilePicker",
            fileName: () => `filepicker.js`
        }
    }
});