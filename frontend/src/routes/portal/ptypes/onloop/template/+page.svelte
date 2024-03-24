<script lang="ts">
    import Autotable from "../../../../../compo/autotable/autotable.svelte";
    import PageLayout from "../../../../../compo/pagelayout/pagelayout.svelte";
    import { getContext, onMount } from "svelte";
    import * as nav from "$lib/nav";
    import type { API } from "$lib/api";
    import { params } from "$lib/params";

    let pid = $params["pid"];

    const api = getContext("__api__") as API;

    let datas: any[] = [];

    const load = async () => {
        const resp = await api.listTemplates(pid);
        if (resp.status !== 200) {
            return;
        }
        datas = resp.data;
    };

    load();
</script>

<PageLayout
    title="OnLoop Templates"
    actions={[
        { name: "add", actionFn: () => nav.gotoProjectOnloopTemplateAdd(pid) },
    ]}
>
    <Autotable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["ttype", "Template type"],
        ]}
        {datas}
        color={["ttype"]}
        actions={[
            {
                Name: "watch",
                Class: "bg-green-400",
                Action: async (id) => {
                    nav.gotoProjectOnloopTemplateWatch(pid, id);
                },
            },

            {
                Name: "push",
                Class: "bg-orange-400",
                Action: async (id) => {
                    nav.gotoProjectOnloopQueueAdd(pid, id)
                },
            },

            {
                Name: "edit",
                Action: async (id) => {
                    nav.gotoProjectOnloopTemplateEdit(pid, id);
                },
            },
            {
                Name: "delete",
                Class: "bg-red-400",
                Action: async (id) => {
                    await api.removeTemplate(pid, id)

                    load()
                },
            },
        ]}
    />
</PageLayout>
