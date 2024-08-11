<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";

    const pid = $page.params["pid"];
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

            <li class="crumb">Catagories</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <!-- Add Catagories -->

        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/catagories/new`}
            class="btn variant-ghost-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Add Catagory
        </a>

        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/products/new`}
            class="btn variant-ghost-secondary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Add Product
        </a>

        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/products`}
            class="btn variant-filled btn-sm"
        >
            <SvgIcon className="h-4 w-4 mr-1" name="rectangle-group" />

            Products
        </a>
    </svelte:fragment>
</AppBar>
