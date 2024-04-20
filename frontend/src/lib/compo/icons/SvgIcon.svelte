<script lang="ts">
  // @ts-ignore
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  export let name: string;
  export let className: string;
</script>

<Icon {name} class={className} />
