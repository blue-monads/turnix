<script lang="ts">
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { AutoForm } from "$lib/compo";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
    const cid = $params["cid"];

    let loading = true;
    let data: any;

    const load = async () => {
        loading = true;
        const resp = await api.getContact(pid, cid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    load();
</script>

{#if loading}
    <Loader />
{:else}
    <AutoForm
        {data}
        schema={{
            name: "Edit Contact",
            fields: [
                {
                    name: "Name",
                    ftype: "TEXT",
                    key_name: "name",
                },
                {
                    name: "Info",
                    ftype: "LONG_TEXT",
                    key_name: "info",
                },
                {
                    name: "Type",
                    ftype: "SELECT",
                    key_name: "ctype",
                    options: ["vendor", "client"],
                },
                {
                    name: "Email",
                    ftype: "EMAIL",
                    key_name: "email",
                },
                {
                    name: "Phone",
                    ftype: "TEXT",
                    key_name: "phone",
                },
                {
                    name: "Second phone",
                    ftype: "TEXT",
                    key_name: "phone2",
                },
                {
                    name: "Third phone",
                    ftype: "TEXT",
                    key_name: "phone3",
                },
                {
                    name: "Address",
                    ftype: "LONG_TEXT",
                    key_name: "address",
                },
                {
                    name: "Second Address",
                    ftype: "LONG_TEXT",
                    key_name: "address2",
                },
            ],
            required_fields: ["name"],
        }}
        onSave={async (data) => {
            const res = await api.addContact(pid, data);
            if (res.status !== 200) {
                return;
            }

            goto(`/z/pages/portal/projects/books/${pid}/contacts`);
        }}
    />
{/if}
