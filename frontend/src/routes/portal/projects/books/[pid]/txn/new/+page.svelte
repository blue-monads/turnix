<script lang="ts">
  import { goto } from "$app/navigation";

  import type { RootAPI } from "$lib/api";

  import { NewBookAPI, type Account } from "$lib/projects/books";
  import { AppBar } from "@skeletonlabs/skeleton";
  import { getContext } from "svelte";

  import { page } from "$app/stores";
  import TxnForm from "../TxnForm.svelte";

  const pid = $page.params["pid"];
  const api = NewBookAPI(getContext("__api__") as RootAPI);

  const onSubmit = async (data: Record<string, any>) => {
    const resp = await api.addTxn(pid, data as any);
    if (resp.status !== 200) {
      return;
    }

    goto(location.pathname.replace("/new", ""));
  };
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

<TxnForm {onSubmit} />
