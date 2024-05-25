<script lang="ts">
    import { goto } from "$app/navigation";

    import type { RootAPI } from "$lib/api";
    import { Loader } from "$lib/compo";
    import Transaction from "$lib/container/books/Transaction.svelte";
    import { NewBookAPI, type Account } from "$lib/projects/books";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";

    import { page } from "$app/stores";

    export let onSubmit: Function;

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
    const store = getModalStore();

    let loading = true;
    let accountsIndex: Record<number, string> = {};
    let accounts: Account[] = [];

    const load = async () => {
        loading = true;
        const resp = await api.listAccount(pid);
        if (resp.status !== 200) {
            return;
        }

        (resp.data || []).forEach((acc: Record<string, any>) => {
            accountsIndex[Number(acc["id"])] = acc["name"];
        });

        accounts = resp.data;
        loading = false;
    };

    load();
</script>

{#if loading}
    <Loader />
{:else}
    <Transaction
        {onSubmit}
        {accountsIndex}
        onPickAccount={async () => {
            return new Promise((res, onerr) => {
                store.trigger({
                    type: "component",
                    component: "books_account_picker",
                    // @ts-ignore
                    meta: {
                        data: accounts,
                        onClick: (id) => {
                            res(id);
                        },
                    },
                });
            });
        }}
        {...$$restProps}
    />
{/if}
