<script lang="ts">
    import AutoForm from "../../../../../../compo/autoform/auto_form.svelte";
    import { gotoProjectOnloopTemplates } from "$lib/nav";
    import { params } from "$lib/params";
    import type { API } from "$lib/api";
    import { getContext } from "svelte";

    let pid = $params["pid"];

    const api = getContext("__api__") as API;

    let message = "";
</script>

<AutoForm
    data={{}}
    {message}
    schema={{
        fields: [
            { ftype: "TEXT", key_name: "name", name: "Name" },
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "contents",
                name: "Contents",
            },
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "result",
                name: "Result",
            },
            {
                ftype: "LONG_TEXT",
                key_name: "contentScript",
                name: "Content script",
            },
        ],
        name: "Add Template",
        required_fields: ["name"],
    }}
    onSave={async (newdata) => {
        const resp = await api.addTemplate(pid, newdata)
        if (resp.status !== 200) {
            message = resp.data;
            return;
        }

        gotoProjectOnloopTemplates(pid)
    }}
/>
