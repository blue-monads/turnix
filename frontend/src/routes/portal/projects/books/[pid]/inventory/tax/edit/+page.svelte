<script lang="ts">
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoForm } from "$lib/compo";
    import { goto } from "$app/navigation";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $page.params["pid"];
    const tid = $params["tid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let message = $state("");
    let data: any = $state();
    let loading = $state(true);

    const load = async () => {
        loading = true;
        const resp = await api.getTax(pid, tid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    load();

    const submit = async (data: Record<string, any>) => {
        const resp = await api.updateTax(pid, tid, data);
        if (resp.status !== 200) {
            message = resp.data.message;
            return;
        }

        goto(`/z/pages/portal/projects/books/${pid}/inventory/tax`);
    };
</script>

{#if loading}
    <Loader />
{:else}
    <AutoForm
        {message}
        {data}
        schema={{
            fields: [
                {
                    name: "Name",
                    ftype: "TEXT",
                    key_name: "name",
                },
                {
                    name: "Type",
                    ftype: "SELECT",
                    key_name: "ttype",
                    options: ["item_percent"],
                },
                {
                    name: "Rate",
                    ftype: "INT",
                    key_name: "rate",
                },
                {
                    name: "Info",
                    ftype: "LONG_TEXT",
                    key_name: "info",
                },
            ],
            name: "Edit Catagory",
            required_fields: [],
        }}
        onSave={submit}
    />
{/if}
