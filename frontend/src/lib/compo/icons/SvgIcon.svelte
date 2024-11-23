<script lang="ts">
  // @ts-ignore
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  interface Props {
    name: string;
    className: string;
  }

  let { name, className }: Props = $props();
</script>

<Icon {name} class={className} />
