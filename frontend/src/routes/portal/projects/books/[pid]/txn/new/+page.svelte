<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import type { RootAPI } from "$lib/api";
  import { Loader } from "$lib/compo";
  import Transaction from "$lib/container/books/Transaction.svelte";
  import { NewBookAPI } from "$lib/projects/books";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
  import { getContext } from "svelte";

  import { params } from "$lib/params";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  const pid = $page.params["pid"];
  const api = NewBookAPI(getContext("__api__") as RootAPI);
  const store = getModalStore();

  let loading = true;
  let accountsIndex: Record<number, string> = {};
  let accounts: never[] = [];

  const onSubmit = async (data: Record<string, any>) => {
    const resp = await api.addTxn(pid, data);
    if (resp.status !== 200) {
      return;
    }

    goto(location.pathname.replace("/new", ""));
  };

  const load = async () => {
    loading = true;
    const resp = await api.listAccount(pid);
    if (resp.status !== 200) {
      return;
    }

    (resp.data || []).forEach((acc: Record<string, any>) => {
      accountsIndex[Number(acc["id"])] = acc["name"];
    });

    accounts = resp.data;
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
      <li class="crumb">Add Transaction</li>
    </ol>
  </div>

 
</AppBar>

{#if loading}
  <Loader />
{:else}
  <Transaction
    {onSubmit}
    {accountsIndex}
    onPickAccount={async () => {
      return new Promise((res, onerr) => {
        store.trigger({
          type: "component",
          component: "books_account_picker",
          // @ts-ignore
          meta: {
            data: accounts,
            onClick: (id) => {
              res(id);
            },
          },
        });
      });
    }}
  />
{/if}
