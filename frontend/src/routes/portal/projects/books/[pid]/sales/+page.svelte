<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI, type Sale } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoTable, Loader } from "$lib/compo";
    import { goto } from "$app/navigation";
    import SalesFilter from "./sub/SalesFilter.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let sales: Sale[] = [];
    let loading = true;

    const load = async () => {
        loading = true;
        const resp = await api.listSales(pid);
        if (resp.status !== 200) {
            return;
        }

        sales = (resp.data || []).map(item => {
            return {
                ...item,
                created_at: new Date(item.created_at).toISOString(),
                updated_at: new Date(item.updated_at).toISOString(),
            }
        });

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

            <li class="crumb">Sales</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <!-- Add Catagories -->

        <a
            href={`/z/pages/portal/projects/books/${pid}/sales/new`}
            class="btn variant-ghost-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Sale
        </a>

        <a
            href={`/z/pages/portal/projects/books/${pid}/sales/estimates`}
            class="btn variant-filled-secondary btn-sm"
        >
            <SvgIcon className="h-4 w-4 mr-1" name="document" />

            Estimates
        </a>

        <a
            href={`/z/pages/portal/projects/books/${pid}/sales/invoice`}
            class="btn variant-filled btn-sm"
        >
            <SvgIcon className="h-4 w-4 mr-1" name="document-text" />

            Invoices
        </a>
    </svelte:fragment>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <SalesFilter />

    <AutoTable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["title", "Title"],
            ["notes", "Notes"],
            ["client_name", "Client"],
            ["total", "Total"],
            ["created_at", "Created At"],
            ["updated_at", "Updated At"],
        ]}
        datas={sales}
        color={["ctype"]}
        actions={[
            {
                Name: "preview",
                Class: "variant-filled-primary",
                Action: async (id) => {
                    goto(
                        `/z/pages/portal/projects/books/${pid}/sales/preview?sid=${id}`,
                    );
                },
            },

            {
                Name: "options",
                Class: "variant-filled-secondary",
                Action: async (id) => {},
            },

            {
                Name: "delete",
                Class: "variant-filled-error",
                Action: async (id) => {
                    const res = await api.deleteSale(pid, id);
                    if (res.status !== 200) {
                        return;
                    }

                    load();
                },
            },
        ]}
    />
{/if}
