<script lang="ts">
  import type { RootAPI } from "$lib/api";
  import Account from "../edit/Account.svelte";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import { params } from "$lib/params";
  import { Loader } from "$lib/compo";
  import { goto } from "$app/navigation";
  import { AppBar } from "@skeletonlabs/skeleton";
  import { page } from "$app/stores";

  const pid = $page.params["pid"];

  const api = NewBookAPI(getContext("__api__") as RootAPI);

  let loading = $state(false);
  const onChange = async (data: Record<string, any>) => {
    loading = true;

    await api.addAccount(pid, data);

    goto(location.pathname.replace("/new", ""));

    loading = false;
  };
</script>

<AppBar>
  <div slot="lead" class="flex gap-2">
    <ol class="breadcrumb">
      <li class="crumb">
        <a class="anchor" href="/z/pages/portal/projects/books">Books</a>
      </li>
      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li class="crumb">Add Account</li>
    </ol>
  </div>
</AppBar>

{#if loading}
  <Loader />
{:else}
  <Account {onChange} />
{/if}
