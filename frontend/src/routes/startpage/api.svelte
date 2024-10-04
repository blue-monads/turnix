<script lang="ts">
    import { setContext } from "svelte";

    interface Props {
        api: any;
        children?: import('svelte').Snippet;
    }

    let { api, children }: Props = $props();

    setContext("__start_api__", api);
</script>

{#if api}
    {@render children?.()}
{/if}
