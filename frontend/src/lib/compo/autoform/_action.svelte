<script lang="ts">
 
  import SvgIcon from "../icons/SvgIcon.svelte";

  interface Props {
    onClick: Function;
    name?: string;
  }

  let { onClick, name = "" }: Props = $props();
  let loading = $state(false);

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
  onclick={handler}
  class="btn btn-sm variant-filled"
>
  {#if loading}
    <SvgIcon name="arrow-path" className="h-5 w-5 animate-spin" />
  {:else}
    <SvgIcon name="arrow-down-on-square" className="h-5 w-5" />
  {/if}

  {name}
</button>
