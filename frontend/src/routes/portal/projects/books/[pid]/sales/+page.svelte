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
        const resp = await api.listCatagories(pid);
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
