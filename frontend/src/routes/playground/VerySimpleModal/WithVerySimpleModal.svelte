<script lang="ts">
    import { setContext } from "svelte";
    import VerySimpleModal from "./VerySimpleModal.svelte";

    let showModal = $state(false);
    let innerContent: any = $state();

    const modalctx = {
        close: () => {
            innerContent = undefined;
            showModal = false;
        },
        show: (inner: any) => {
            innerContent = inner;
            showModal = true;
        },
    }

    setContext("__simple_modal", modalctx);
</script>

<VerySimpleModal bind:showModal>
    {@render innerContent?.(modalctx)}
</VerySimpleModal>
