<script lang="ts">
    import type { RootAPI } from "$lib/api";

    import { NewBookAPI, type Account } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { Loader } from "$lib/compo";
    import { goto } from "$app/navigation";
    import { AppBar } from "@skeletonlabs/skeleton";
    import AccountPage from "$lib/container/books/Account.svelte";
    import { params } from "$lib/params";
    import { page } from "$app/stores";
   
    const pid = $page.params["pid"];
    const aid = $params["aid"];
  
    const api = NewBookAPI(getContext("__api__") as RootAPI);
  
    let loading = true;
    let data: Account | null = null;

    const onChange = async (data: Record<string, any>) => {
        loading = true;
        const resp = await api.updateAccount(pid, aid, data)
        if (resp.status !== 200) {
            return
        }

        await load()
    };

    const load = async  () => {
        loading = true
        const resp = await api.getAccount(pid, aid)
        if (resp.status !== 200) {
            return
        }
        data = resp.data
        loading = false
    }

    load()


    $: {
        console.log("@aid", $params)
    }
     

  </script>
  
  <AppBar>
    <div slot="lead" class="flex gap-2">
      <ol class="breadcrumb">
        <li class="crumb">
          <a class="anchor" href="/z/pages/portal/projects/books">Books</a>
        </li>
        <li class="crumb-separator" aria-hidden>&rsaquo;</li>
        <li class="crumb">Accounts</li>
        <li class="crumb-separator" aria-hidden>&rsaquo;</li>
        <li class="crumb">{aid || ""}</li>
      </ol>
    </div>
  </AppBar>
  
  {#if loading}
    <Loader />
  {:else}
    <AccountPage {onChange} acc_type={data?.acc_type} info={data?.info} name={data?.name} edit={true} />
  {/if}
  