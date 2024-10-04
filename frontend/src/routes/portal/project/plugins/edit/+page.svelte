<script lang="ts">
    import { gotoPorjects } from "$lib/nav";
    import { AutoForm, PageLayout, AutoTable, Loader } from "$lib/compo";
    import { params } from "$lib/params";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";

    const api = getContext("__api__") as RootAPI;

    let pid = $params["pid"];
    let id = $params["id"];

    let data = $state({});
    let message = $state("");
    let loading = $state(true);

    const load = async () => {
        loading = true;
        const resp = await api.getProjectPlugin(pid, id);
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
        {message}
        schema={{
            fields: [
                { ftype: "TEXT", key_name: "name", name: "Name" },
                {
                    ftype: "SELECT",
                    key_name: "ptype",
                    name: "type",
                    options: ["sideapp_v1"],
                },
                {
                    ftype: "LONG_TEXT",
                    key_name: "server_code",
                    name: "Server Code",
                },
                {
                    ftype: "LONG_TEXT",
                    key_name: "client_code",
                    name: "Client Code",
                },
            ],
            name: "Edit Project Plugin",
            required_fields: [],
        }}
        onSave={async (newdata) => {
            const resp = await api.updateProjectPlugin(pid, id, newdata);
            if (resp.status !== 200) {
                message = resp.data;
                return;
            }

            gotoPorjects();
        }}
    />
{/if}
