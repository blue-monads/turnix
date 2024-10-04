<script>
	import "../styles.css";
	/** @type {{children?: import('svelte').Snippet}} */
	let { children } = $props();
</script>

<div class="p-2 md:p-12">
	{@render children?.()}
</div>
