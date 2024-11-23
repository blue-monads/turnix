<script lang="ts">
  import { AppBar } from "@skeletonlabs/skeleton";
  import SvgIcon from "../icons/SvgIcon.svelte";

  interface Props {
    title: string;
    actions: { name: string; icon?: string; actionFn: Function }[];
    children?: import('svelte').Snippet;
  }

  let { title, actions, children }: Props = $props();
</script>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4 capitalize">{title}</h4>
  </svelte:fragment>

  <div slot="trail" class="">
    {#if actions}
      {#each actions as action}
        <button
          onclick={() => {
            action.actionFn();
          }}
          class="btn btn-sm variant-filled-primary text-secondary-50"
        >
          {#if action.icon}
            <SvgIcon name={action.icon} className="h-4 w-4" />
          {/if}

          {action.name}
        </button>
      {/each}
    {/if}
  </div>
</AppBar>

<div class="p-2">
  {@render children?.()}
</div>
