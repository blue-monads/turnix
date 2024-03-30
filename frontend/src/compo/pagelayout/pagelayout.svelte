<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    import Icon from "@krowten/svelte-heroicons/Icon.svelte";

    export let title: string;
    export let actions: { name: string; icon?: string, actionFn: Function }[];
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4 capitalize">{title}</h4>
    </svelte:fragment>

    <div slot="trail" class="">
        {#if actions}
            {#each actions as action}
                <button
                    on:click={() => {
                        action.actionFn();
                    }}
                    class="btn btn-sm variant-filled-primary text-secondary-50"
                    >

                    {#if action.icon}
                        <Icon name={action.icon} class="h-4 w-4" />
                    {/if}

                    
                    {action.name}
                    
                    
                    </button
                >
            {/each}
        {/if}
    </div>
</AppBar>

<div class="p-2">
    <slot />
</div>
