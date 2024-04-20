<script>
  import { goto } from "$app/navigation";
  import { FloatyButton } from "$lib/compo";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import ListProject from "../../listProject.svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";

  const store = getModalStore();
</script>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4">Books</h4>
  </svelte:fragment>

  <svelte:fragment slot="trail">
    <a
      href="/z/pages/portal/projects/books/new"
      class="btn variant-filled btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="plus" />

      account
    </a>
  </svelte:fragment>
</AppBar>

<ListProject ptype={"books"} />
