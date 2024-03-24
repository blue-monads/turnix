<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    export let title: string;
    export let actions: { name: string; actionFn: Function }[];
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h2 class="h3 capitalize">{title}</h2>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        {#if actions}
            {#each actions as action}
                <button
                    type="button"
                    on:click={() => {
                        action.actionFn();
                    }}
                    class="btn variant-filled-primary text-secondary-50"
                    >{action.name}</button
                >
            {/each}
        {/if}
    </svelte:fragment>
</AppBar>

<slot />
