<script lang="ts">
    import { WithVerySimpleModal } from "../../playground/VerySimpleModal";

    interface Props {
        children?: import("svelte").Snippet;
    }

    let { children }: Props = $props();
</script>

<WithVerySimpleModal> 
    {@render children?.()}

</WithVerySimpleModal>

