<script lang="ts">
    import { gotoPorjects } from "$lib/nav";
    import { AutoForm, PageLayout, AutoTable } from "$lib/compo";
    import type {  RootAPI } from "$lib/api";
    import { getContext } from "svelte";

    const api = getContext("__api__") as RootAPI;

    let message = "";
</script>

<AutoForm
    data={{}}
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
        const resp = await api.addProject(newdata);
        if (resp.status !== 200) {
            message = resp.data;
            return;
        }

        gotoPorjects();
    }}
/>
