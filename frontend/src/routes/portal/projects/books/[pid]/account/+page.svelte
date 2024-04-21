<script lang="ts">
  import { goto } from "$app/navigation";
  import { FloatyButton } from "$lib/compo";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
  import type { RootAPI } from "$lib/api";
  import { params } from "$lib/params";
  import { page } from "$app/stores";

  const pid = $page.params["pid"];

  const store = getModalStore();
  const api = NewBookAPI(getContext("__api__") as RootAPI);

  let accounts = [];

  const load = async () => {};
</script>

<AppBar>
  <svelte:fragment slot="lead">
    <ol class="breadcrumb">
      <li class="crumb">
        <a class="anchor" href="/z/pages/portal/projects/books">Books</a>
      </li>
      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li>Account</li>
    </ol>
  </svelte:fragment>

  <svelte:fragment slot="trail">
    <a
      href={`/z/pages/portal/projects/books/${pid}/account/new`}
      class="btn variant-filled btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="plus" />

      account
    </a>
  </svelte:fragment>
</AppBar>
