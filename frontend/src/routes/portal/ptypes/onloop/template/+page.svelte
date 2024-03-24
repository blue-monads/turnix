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
    actions={[{ name: "add", actionFn: nav.gotoAddProjectLactions }]}
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
                Name: "explore",
                Class: "bg-green-400",
                Action: async (id, data) => {},
            },

            {
                Name: "edit",
                Action: async (id) => {},
            },
            {
                Name: "delete",
                Class: "bg-red-400",
                Action: async (id) => {},
            },
        ]}
    />
</PageLayout>
