<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoTable } from "$lib/compo";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let stocks: any[] = $state([]);

    const load = async () => {
        const resp = await api.listProductStockIn(pid);
        if (resp.status !== 200) {
            return;
        }

        stocks = resp.data;
    };

    load();
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>

            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li class="crumb">Stocks</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/stocks/new`}
            class="btn variant-filled btn-sm"
        >
            <SvgIcon className="h-4 w-4 mr-1" name="plus" />
            Stock In
        </a>
    </svelte:fragment>
</AppBar>

<AutoTable
    action_key={"id"}
    key_names={[
        ["id", "ID"],
        ["name", "Name"],
    ]}
    datas={stocks}
    color={[]}
    actions={[
        {
            Name: "edit",
            Class: "variant-filled-tertiary",
            Action: async (id) => {},
        },
        {
            Name: "delete",
            Class: "variant-filled-error",
            Action: async (id) => {},
        },
    ]}
/>
