<script lang="ts">
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import Transactions from "./sub/Transactions.svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import type { RootAPI } from "$lib/api";
  import { Loader } from "$lib/compo";
  import type { TxnLine } from "./sub/txntype";
  import { formatResponse } from "./format";
  import { page } from "$app/stores";

  const pid = $page.params["pid"];
  const api = NewBookAPI(getContext("__api__") as RootAPI);
  const store = getModalStore();

  let loading = $state(true);
  let data: TxnLine[] = $state([]);
  let accountsIndex: Record<number, string> = $state({});
  let maxId: number = $state([]);

  const load = async () => {
    const rresp = await api.listAccount(pid);

    loading = true;

    await loadTxnLines(true);

    const accounts = (await rresp).data as Record<string, any>[];

    accounts.forEach((account) => {
      const id = account["id"];
      accountsIndex[id] = account["name"];
    });


    loading = false;
  };
  load();

  // prev_page

  const loadTxnLines = async (forward: boolean) => {


    const nextOffset = maxId.length > 0 ? maxId.at(-1) : 0;
    if (!forward && maxId.length > 0) {
      maxId = maxId.slice(0, -1);
    }

    const resp = await api.listTxnWithLines(pid, nextOffset);
    if (resp.status !== 200) {
      return;
    }

    const { maxId: nextMaxId, txns } = formatResponse(resp.data);

    data = txns;
    maxId.push(nextMaxId);



  }


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

      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li class="crumb">Transactions</li>
    </ol>
  </div>

  <div slot="trail" class="flex gap-2">
    <button class="inline-flex gap-1">
      <SvgIcon name="magnifying-glass" className="w-5 h-5" />
      Find
    </button>

    <a
      href={`/z/pages/portal/projects/books/${pid}/txn/new`}
      class="inline-flex gap-1"
    >
      <SvgIcon name="plus" className="w-5 h-5" />
      New
    </a>
  </div>
</AppBar>

{#if loading}
  <Loader />
{:else}
  {#key data}
    <Transactions 
      {accountsIndex} 
      lineData={data} 
      {pid} 
      on:prev_page={() => loadTxnLines(false)}
      on:next_page={() => loadTxnLines(true)}
      
      />
  {/key}
{/if}
