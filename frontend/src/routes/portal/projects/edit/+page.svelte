<script lang="ts">
    import { gotoPorjects } from "$lib/nav";
    import { AutoForm, PageLayout, AutoTable, Loader } from "$lib/compo";
    import { params } from "$lib/params";
    import type { API } from "$lib/api";
    import { getContext } from "svelte";

    const api = getContext("__api__") as API;

    let pid = $params["pid"];

    let data = {};
    let message = "";
    let loading = true;

    const load = async () => {
        loading = true;
        const resp = await api.getProject(pid);
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
                    options: ["onloop"],
                },
            ],
            name: "Add Project",
            required_fields: ["name"],
        }}
        onSave={async (newdata) => {
            const resp = await api.updateProject(pid, newdata);
            if (resp.status !== 200) {
                message = resp.data;
                return;
            }

            gotoPorjects();
        }}
    />
{/if}
