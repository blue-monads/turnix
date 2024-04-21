<script lang="ts">
  import type { RootAPI } from "$lib/api";
  import Account from "$lib/container/books/Account.svelte";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";

  import { page } from "$app/stores";
  import { Loader } from "$lib/compo";
  import { goto } from "$app/navigation";

  const pid = $page.params["pid"];

  const api = NewBookAPI(getContext("__api__") as RootAPI);

  let loading = false;
  const onChange = async (data: Record<string, any>) => {
    loading = true

    await api.addAccount(pid, data)

    goto(location.pathname.replace("/new", ""))

    loading = false


  };
</script>

{#if loading}
  <Loader />
{:else}
  <Account {onChange} />
{/if}
