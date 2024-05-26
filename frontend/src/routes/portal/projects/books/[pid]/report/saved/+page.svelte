<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";

    const pid = $page.params["pid"];

    const store = getModalStore();
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    const load = async () => {


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
            <li>Reports</li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Saved</li>
        </ol>
    </svelte:fragment>
</AppBar>
