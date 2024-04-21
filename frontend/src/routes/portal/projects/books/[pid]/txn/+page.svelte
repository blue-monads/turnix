<script lang="ts">
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import Transactions from "$lib/container/books/Transactions.svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";

  import { params } from "$lib/params";
  import { page } from "$app/stores";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import type { RootAPI } from "$lib/api";
  const pid = $page.params["pid"];
  const api = NewBookAPI(getContext("__api__") as RootAPI);
  const store = getModalStore();
</script>

<AppBar>
  <div slot="lead" class="flex gap-2">
    <ol class="breadcrumb">
      <li class="crumb">
        <a class="anchor" href="/z/pages/portal/projects/books">Books</a>
      </li>
      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li class="crumb">
        <a class="anchor" href={`/z/pages/portal/projects/books/${pid}`}
          >Account</a
        >
      </li>
    </ol>
  </div>

  <div slot="trail" class="flex gap-2">
    <button class="inline-flex gap-1">
      <SvgIcon name="magnifying-glass" className="w-5 h-5" />
      Find
    </button>

    <button class="inline-flex gap-1">
      <SvgIcon name="plus" className="w-5 h-5" />
      New
    </button>
  </div>
</AppBar>

<Transactions />
