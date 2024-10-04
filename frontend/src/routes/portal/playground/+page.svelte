<script lang="ts">
  import AutoForm from "$lib/compo/autoform/auto_form.svelte";
  import { getModalStore } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const store = getModalStore();

  const items = [
    { name: "rename", icon: "pencil-square" },
    { name: "download", icon: "arrow-down-on-square" },
    { name: "delete", icon: "trash" },
  ];
</script>

<h4 class="h4">Playground</h4>

<button
  onclick={() => {
    store.trigger({
      type: "component",
      component: "file_picker",
      meta: {},
    });
  }}
  class="btn btn-sm variant-filled"
>
  click me
</button>

<button
  class="relative group transition-all duration-200 focus:overflow-visible w-max h-max p-1 overflow-hidden flex flex-row items-center justify-center bg-white gap-2 rounded-lg border border-zinc-200"
>
  <span> Dropdown </span>

  <div
    class="absolute shadow-lg top-10 left-0 w-32 h-max p-1 bg-white border border-zinc-200 rounded-lg flex flex-col gap-2"
  >
    {#each items as item}
      <span
        class="flex gap-1 justify-start items-center hover:bg-zinc-100 p-1 rounded-lg"
      >
        <SvgIcon name={item.icon} className="w-4 h-4" />

        <p>{item.name}</p>
      </span>
    {/each}
  </div>
</button>
