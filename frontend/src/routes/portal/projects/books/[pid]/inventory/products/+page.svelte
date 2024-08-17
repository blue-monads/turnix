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

    let cats: any[] = [];

    const load = async () => {
        const resp = await api.listProducts(pid);
        if (resp.status !== 200) {
            return;
        }

        cats = resp.data;
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
            <li>
                <a
                    class="anchor"
                    href={`/z/pages/portal/projects/books/${pid}/inventory`}
                    >Inventory</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>

            <li class="crumb">Products</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">

        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/catagories/new`}
            class="btn variant-ghost-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Catagory
        </a>

        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/products/new`}
            class="btn variant-ghost-secondary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Product
        </a>

        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/catagories`}
            class="btn variant-filled btn-sm"
        >
            <SvgIcon className="h-4 w-4 mr-1" name="squares-2x2" />

            Catagories
        </a>
    </svelte:fragment>
</AppBar>

<AutoTable
    action_key={"id"}
    key_names={[
        ["id", "ID"],
        ["name", "Name"],
        ["variant_id", "Variant"],
        ["stock_count", "Stock"],
        ["info", "Info"],
        ["price", "Price"],
    ]}
    datas={cats}
    color={["variant_id"]}
    actions={[
        {
            Name: "Sales",
            Class: "variant-filled-primary",
            icon: "plus",
            Action: async (id) => {
            },
        },
        {
            Name: "Stocks", 
            Class: "variant-filled-secondary",
            icon: "plus",
            Action: async (id) => {
            },
        },

        {
            Name: "edit",
            Class: "variant-filled-tertiary",
            Action: async (id) => {
                goto(
                    `/z/pages/portal/projects/books/${pid}/inventory/products/edit?pid=${pid}&prid=${id}`,
                );
            },
        },
        {
            Name: "delete",
            Class: "variant-filled-error",
            Action: async (id) => {
                await api.deleteProduct(pid, id);

                load();
            },
        },
    ]}
/>
