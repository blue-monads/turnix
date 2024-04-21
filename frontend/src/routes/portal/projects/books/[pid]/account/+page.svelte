<script lang="ts">
  import { goto } from "$app/navigation";
  import { AutoTable, FloatyButton } from "$lib/compo";
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

    <a
      href={`/z/pages/portal/projects/books/${pid}/txn/new`}
      class="btn variant-filled-secondary btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="plus" />

      transaction
    </a>

    <a
      href={`/z/pages/portal/projects/books/${pid}/txn`}
      class="btn variant-glass-secondary btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="clipboard-document" />

      transactions
    </a>

    <a
      href={`/z/pages/portal/projects/books/${pid}/report`}
      class="btn variant-glass-tertiary btn-sm"
    >
      <SvgIcon className="h-4 w-4" name="document-chart-bar" />

      reports
    </a>
  </svelte:fragment>
</AppBar>

<AutoTable
  action_key={"id"}
  key_names={[
    ["id", "ID"],
    ["name", "Name"],
    ["info", "Info"],
    ["acc_type", "Account type"],
  ]}
  datas={accounts}
  color={["acc_type"]}
  actions={[
    {
      Name: "explore txns",
      Class: "variant-filled-primary",

      icon: "plus",
      Action: async (id) => {
        const next = `/txn/account?pid=${pid}&aid=${id}`;

        if (location.pathname.endsWith("/account")) {
          goto(location.pathname.replace("/account", next));
        } else {
          goto(location.pathname + next);
        }
      },
    },

    {
      Name: "edit",
      Class: "variant-filled-secondary",
      Action: async (id) => {
        const next = `/account/edit?pid=${pid}&aid=${id}`;

        if (location.pathname.endsWith("/account")) {
          goto(location.pathname.replace("/account", next));
        } else {
          goto(location.pathname + next);
        }
      },
    },
  ]}
/>
