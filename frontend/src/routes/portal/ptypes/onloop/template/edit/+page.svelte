<script lang="ts">
    import AutoForm from "../../../../../../compo/autoform/auto_form.svelte";
    import Loader from "../../../../../../compo/loader/loader.svelte";

    import { gotoProjectOnloopTemplates } from "$lib/nav";
    import { params } from "$lib/params";
    import type { API } from "$lib/api";
    import { getContext } from "svelte";

    const api = getContext("__api__") as API;

    let pid = $params["pid"];
    let tid = $params["tid"];

    let message = "";
    let loading = true;
    let data = {}

    const load = async  () => {
        loading = true
        const resp = await api.getTemplate(pid, tid )
        if (resp.status !== 200) {
            return
        }
        data = resp.data;
        loading = false
    }

    load()


</script>

{#if loading}
    <Loader />
{:else}
    <AutoForm
        data={data}
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
            name: "Edit Template",
            required_fields: ["name"],
        }}
        onSave={async (newdata) => {
            const resp = await api.updateTemplate(pid, tid, newdata);
            if (resp.status !== 200) {
                message = resp.data;
                return;
            }

            gotoProjectOnloopTemplates(pid);
        }}
    />
{/if}
