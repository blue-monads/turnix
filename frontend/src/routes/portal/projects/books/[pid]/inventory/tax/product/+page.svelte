<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI, type ProductTax } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoTable } from "$lib/compo";
    import { params } from "$lib/params";

    const pid = $page.params["pid"];
    const tid = $params["tid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let taxes: ProductTax[] = $state([]);
    let loading = true;

    const load = async () => {
        loading = true;

        const resp = await api.listTaxProduct(pid);
        if (resp.status !== 200) {
            return;
        }

        taxes = resp.data;
        loading = false;
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

            <li class="crumb">Taxes</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <button class="btn variant-ghost-primary btn-sm">
            <SvgIcon className="h-4 w-4" name="plus" />
            Product Tax
        </button>
    </svelte:fragment>
</AppBar>

<AutoTable
    action_key={"id"}
    key_names={[
        ["id", "ID"],
        ["catagory_id", "Catagory"],
        ["product_id", "Product"],
        ["tax_id", "Tax"],
    ]}
    datas={taxes}
    color={["ttype"]}
    actions={[
        {
            Name: "delete",
            Class: "variant-filled-error",
            Action: async (id) => {
                await api.deleteTax(pid, id);

                load();
            },
        },
    ]}
/>
