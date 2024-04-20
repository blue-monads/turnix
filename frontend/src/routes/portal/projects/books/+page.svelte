<script>
  import { goto } from "$app/navigation";
  import { FloatyButton } from "$lib/compo";
  import ListProject from "../../listProject.svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";

  const store = getModalStore();
</script>

<ListProject ptype={"books"} />

<FloatyButton
  handler={() => {
    goto("/z/pages/portal/projects/books/new");
  }}
/>
