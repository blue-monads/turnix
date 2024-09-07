<script lang="ts">
  import AutoForm from "$lib/compo/autoform/auto_form.svelte";
  import { getModalStore } from "@skeletonlabs/skeleton";
  import NewSale from "./new/new.svelte";

  const store = getModalStore();

</script>

<h4 class="h4">Playground</h4>

<button
  on:click={() => {

    store.trigger({
      type: "component",
      component: "file_picker",
      meta: { },
    });

  }}
  class="btn btn-sm variant-filled"
>
  click me
</button>
