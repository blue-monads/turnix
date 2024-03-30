<script lang="ts">
    import Autotable from "../../../lib/compo/autotable/autotable.svelte";
    import PageLayout from "../../../lib/compo/pagelayout/pagelayout.svelte";
    import { getContext, onMount } from "svelte";
    import * as nav from "$lib/nav";
    import type { API } from "$lib/api";

    const api = getContext("__api__") as API;

    let datas: any[] = [];

    const load = async () => {
        const resp = await api.listProjects();
        if (resp.status !== 200) {
            return;
        }
        datas = resp.data;
    };

    load();
</script>

<PageLayout
    title="projects"
    actions={[{ name: "add", actionFn: nav.gotoAddProject }]}
>
    <Autotable
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
