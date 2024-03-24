<script lang="ts">
    import { setContext } from "svelte";

    export let api;

    setContext("__api__", api);
</script>

{#if api}
    <slot />
{/if}
