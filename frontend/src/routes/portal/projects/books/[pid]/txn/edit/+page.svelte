<script lang="ts">
    import { goto } from "$app/navigation";
    import type { RootAPI } from "$lib/api";
    import { Loader } from "$lib/compo";

    import { NewBookAPI, type Account } from "$lib/projects/books";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";
    import { page } from "$app/stores";
    import TxnForm from "../TxnForm.svelte";
    import { params } from "$lib/params";

    const pid = $page.params["pid"];
    const tid = $params["tid"]


    const api = NewBookAPI(getContext("__api__") as RootAPI);
    const store = getModalStore();

    let loading = true;

    const onSubmit = async (data: Record<string, any>) => {};

    const load = async () => {
        loading = false

        const resp = await api.getTxnWithLines(pid, tid)


    };

    load();
</script>

<AppBar>
    <div slot="lead" class="flex gap-2">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li class="crumb">Edit Transaction</li>
        </ol>
    </div>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <TxnForm {onSubmit} edit={true} />
{/if}
