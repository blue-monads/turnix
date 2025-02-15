<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    interface Props {
        handler: Function;
    }

    let { handler }: Props = $props();
</script>

<div class="fixed bottom-5 right-5">
    <button
        class="btn-icon variant-filled bg-primary-hover-token"
        onclick={() => {
            handler && handler()
        }}
    >
        <SvgIcon name="squares-plus" className="w-6 h-6" />
    </button>
</div>
