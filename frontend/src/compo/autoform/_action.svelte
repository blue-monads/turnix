<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let onClick: Function;
  export let name = "";
  let loading = false;

  const handler = async () => {
    loading = true;
    const r = onClick();
    if (r instanceof Promise) {
      await r;
    }
    loading = false;
  };
</script>

<button
  on:click={handler}
  class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-blue-400"
>
  {#if loading}
    <Icon name="arrow-path" class="h-5 w-5 animate-spin" />
  {:else}
    <Icon name="arrow-down-on-square" class="h-5 w-5" />
  {/if}

  {name}
</button>
