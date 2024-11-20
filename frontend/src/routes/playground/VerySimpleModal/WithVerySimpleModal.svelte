<script lang="ts">
    import { setContext } from "svelte";
    import VerySimpleModal from "./VerySimpleModal.svelte";

    interface Props {
        children?: import("svelte").Snippet;
    }

    let { children }: Props = $props();

    let showModal = $state(false);
    let epoch = $state(0);
    // svelte-ignore non_reactive_update
    let innerContent: any;
    // svelte-ignore non_reactive_update
    let pp: any;

    const modalctx = {
        close: () => {
            innerContent = undefined;
            showModal = false;
            epoch = epoch + 1;
        },
        show: (inner: any, ppp: any) => {
            epoch = epoch + 1;
            innerContent = inner;
            pp = ppp;
            showModal = true;
            epoch = epoch + 1;
        },
    };

    setContext("__vs_modal__", modalctx);
</script>

<VerySimpleModal bind:showModal>
    {#key epoch}
        {#if innerContent}
            {@render innerContent(pp)}
        {/if}
    {/key}
</VerySimpleModal>

{@render children?.()}
