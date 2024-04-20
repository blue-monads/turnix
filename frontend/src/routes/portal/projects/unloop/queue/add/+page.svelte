<script lang="ts">
    import { AutoForm } from "$lib/compo";
    import {
        gotoProjectOnloopQueues,
        gotoProjectOnloopTemplates,
    } from "$lib/modules/unloop/route";
    import { params } from "$lib/params";
    import type { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { stringify } from "postcss";

    let pid = $params["pid"];
    let tid = $params["tid"];

    const api = getContext("__api__") as RootAPI;

    let message = "";
</script>

<AutoForm
    data={{}}
    {message}
    schema={{
        fields: [
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "contents",
                name: "Contents",
            },
            {
                ftype: "SELECT",
                key_name: "status",
                name: "Status",
                options: ["submited", "processed", "denied"],
            },
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "extrameta",
                name: "Extra Meta",
            },
        ],
        name: "Add Queue message",
        required_fields: ["name"],
    }}
    onSave={async (newdata) => {
        const resp = await api.addQueueMessage(pid, {
            ...newdata,
            projectId: Number(pid),
            templateId: Number(tid),
        });
        if (resp.status !== 200) {
            return;
        }

        gotoProjectOnloopQueues(pid, tid);
    }}
/>
