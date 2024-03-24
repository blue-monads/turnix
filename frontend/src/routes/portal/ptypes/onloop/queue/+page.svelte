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
        const resp = await api.listQueueMessages(pid)
        if (resp.status !== 200) {
            return;
        }
        datas = resp.data;
    };

    load();
</script>

<PageLayout
    title="Onloop Queue Messages"
    actions={[{ name: "add", actionFn: nav.gotoProjectOnloopQueueAdd }]}
>
    <Autotable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["submitter", "Submitter"],
            ["status", "Status"],
            ["createdAt", "Created At"]
        ]}
        {datas}
        color={["ttype"]}
        actions={[
            {
                Name: "explore",
                Class: "bg-green-400",
                Action: async (id, data) => {},
            },
            {
                Name: "delete",
                Class: "bg-red-400",
                Action: async (id) => {},
            },
        ]}
    />
</PageLayout>
