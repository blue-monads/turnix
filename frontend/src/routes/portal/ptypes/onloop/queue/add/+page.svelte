<script lang="ts">
    import AutoForm from "../../../../../../compo/autoform/auto_form.svelte";
    import { gotoProjectOnloopQueues, gotoProjectOnloopTemplates } from "$lib/nav";
    import { params } from "$lib/params";
    import type { API } from "$lib/api";
    import { getContext } from "svelte";

    let pid = $params["pid"];
    let tid = $params["tid"];

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
                ftype: "SELECT",
                key_name: "ttype",
                name: "type",
                options: ["bexec"],
            },
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "consts",
                name: "Constants",
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
        const resp = await  api.addQueueMessage(pid, {
            ...newdata,
            projectId: Number(pid),
            templateId: Number(tid),
        });
        if (resp.status !== 200) {
            return
        }

        gotoProjectOnloopQueues(pid, tid)
    }}
/>
