<script lang="ts">
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
  import type { RootAPI } from "$lib/api";
  import { page } from "$app/stores";

  const pid = $page.params["pid"];

  const store = getModalStore();
  const api = NewBookAPI(getContext("__api__") as RootAPI);

  let accounts: any[] = [];

  const load = async () => {
    const resp = await api.listAccount(pid);
    if (resp.status !== 200) {
      return;
    }

    accounts = resp.data;
  };

  load();
</script>

<AppBar>
  <svelte:fragment slot="lead">
    <ol class="breadcrumb">
      <li class="crumb">
        <a class="anchor" href="/z/pages/portal/projects/books">Books</a>
      </li>
      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li>Reports</li>
    </ol>
  </svelte:fragment>

  <svelte:fragment slot="trail">
    <a
    href={`/z/pages/portal/projects/books/${pid}/report/editor`}
    class="btn variant-filled-primary btn-sm"
  >
    <SvgIcon className="h-4 w-4" name="pencil-square" />

    Editor 
  </a>
    <a
      href={`/z/pages/portal/projects/books/${pid}/report/live`}
      class="btn variant-filled btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="bolt" />

      Live Report
    </a>

    <a
      href={`/z/pages/portal/projects/books/${pid}/report/saved`}
      class="btn variant-filled-tertiary btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="cloud-arrow-up" />

      Saved Reports
    </a>

    

    <a
      href={`/z/pages/portal/projects/books/${pid}/report/template`}
      class="btn variant-filled-secondary btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="code-bracket-square" />

      Templates
    </a>

  </svelte:fragment>
</AppBar>
