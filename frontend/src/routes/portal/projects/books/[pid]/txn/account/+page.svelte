<script lang="ts">
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  //  import Transactions from "$lib/container/books/Transactions.svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";

  import { params } from "$lib/params";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import type { RootAPI } from "$lib/api";
  import { Loader } from "$lib/compo";
  import { formatResponse } from "../format";
  import { page } from "$app/stores";
  import type { TxnLine } from "../sub/txntype";
  import Transactions from "../sub/Transactions.svelte";

  const pid = $page.params["pid"];
  const aid = $params["aid"];
  const api = NewBookAPI(getContext("__api__") as RootAPI);
  const store = getModalStore();

  let loading = $state(true);
  let data: TxnLine[] = $state([]);
  let accountsIndex: Record<number, string> = $state({});

  const load = async () => {
    const rresp = api.listAccount(pid);

    loading = true;
    const resp = await api.listAccTxnWithLines(pid, aid);
    if (resp.status !== 200) {
      return;
    }

    console.log("@data", resp.data);

    const formateddata = formatResponse(resp.data);
    console.log("@data_____", data);

    const accounts = (await rresp).data as Record<string, any>[];

    accounts.forEach((account) => {
      const id = account["id"];
      accountsIndex[id] = account["name"];
    });

    data = formateddata.txns;

    loading = false;
  };
  load();
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
    <Transactions {accountsIndex} lineData={data} {pid} />
  {/key}
{/if}
