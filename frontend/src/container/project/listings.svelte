<script lang="ts">
    import { AutoTable, PageLayout } from "$lib/compo";

    import { getContext, onMount } from "svelte";
    import * as nav from "$lib/nav";
    import type {  RootAPI } from "$lib/api";
    export let ptype = "";
    export let title = "projects";

    const api = getContext("__api__") as RootAPI;

    let datas: any[] = [];

    const load = async () => {
        const resp = await api.listProjects(ptype);
        if (resp.status !== 200) {
            return;
        }
        datas = resp.data;
    };

    load();
</script>

<PageLayout {title} actions={[{ name: "add", actionFn: nav.gotoAddProject }]}>
    <AutoTable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["ptype", "Project type"],
            ["owner", "Owner"],
        ]}
        {datas}
        color={["ptype"]}
        actions={[
            {
                Name: "explore",
                Class: "bg-green-400",
                Action: async (id, data) => {
                    const ptype = data["ptype"];
                    if (ptype === "onloop") {
                        nav.gotoProjectOnloopTemplates(id);
                    }
                },
            },

            {
                Name: "edit",

                Action: async (id) => {
                    nav.gotoEditProject(id);
                },
            },
            {
                Name: "delete",
                Class: "bg-red-400",
                Action: async (id) => {
                    await api.removeProject(id);

                    load();
                },
            },
        ]}
    />
</PageLayout>
