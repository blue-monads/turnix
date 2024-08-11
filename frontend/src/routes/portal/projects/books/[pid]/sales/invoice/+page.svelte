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


</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>

            <li class="crumb-separator" aria-hidden>&rsaquo;</li>

            <li class="crumb">Invoices</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <!-- Add Catagories -->

        <a
            href={`/z/pages/portal/projects/books/${pid}/sales/invoice/new`}
            class="btn variant-ghost-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Invoice
        </a>


        <a
            href={`/z/pages/portal/projects/books/${pid}/sales`}
            class="btn variant-filled btn-sm"
        >
            <SvgIcon className="h-4 w-4 mr-1" name="banknotes" />

            Sales
        </a>
    </svelte:fragment>
</AppBar>