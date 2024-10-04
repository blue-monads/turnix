<script lang="ts">
    import { AutoForm, PageLayout, AutoTable } from "$lib/compo";
    import type { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { params } from "$lib/params";
    import { goto } from "$app/navigation";

    let pid = $params["pid"];

    const api = getContext("__api__") as RootAPI;
    let message = $state("");
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
        name: "New Plugin",
        required_fields: ["name"],
    }}
    onSave={async (newdata) => {
        const resp = await api.addProjectPlugin(pid, newdata);
        if (resp.status !== 200) {
            message = resp.data;
            return;
        }

        goto(`/z/pages/portal/project/plugins?pid=${pid}`);

    }}
/>
