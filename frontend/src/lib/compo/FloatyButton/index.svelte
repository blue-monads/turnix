<script lang="ts">
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
        <svg viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
            <path
                fill-rule="evenodd"
                d="M12 3.75a.75.75 0 0 1 .75.75v6.75h6.75a.75.75 0 0 1 0 1.5h-6.75v6.75a.75.75 0 0 1-1.5 0v-6.75H4.5a.75.75 0 0 1 0-1.5h6.75V4.5a.75.75 0 0 1 .75-.75Z"
                clip-rule="evenodd"
            />
        </svg>
    </button>
</div>
