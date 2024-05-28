<script lang="ts">
    import { AutoTable, PageLayout } from "$lib/compo";

    import { getContext } from "svelte";
    import * as nav from "$lib/nav";
    import type { RootAPI } from "$lib/api";
    import { params } from "$lib/params";

    const ptype = $params["ptype"];
    const pid = $params["pid"];

    const api = getContext("__api__") as RootAPI;

    let datas: any[] = [];

    const load = async () => {
        const resp = await api.listProjectHooks(pid);
        if (resp.status !== 200) {
            return;
        }
        datas = resp.data;
    };

    load();
</script>

<PageLayout
    title={"Project Hooks"}
    actions={[{ name: "add", actionFn: () => nav.gotoAddProjectHook(ptype, pid) }]}
>
    <AutoTable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["hook_type", "Hook Type"],
            ["event", "Event"]

        ]}
        {datas}
        color={["hook_type"]}
        actions={[
            {
                Name: "explore",
                Class: "bg-green-400",
                Action: async (id, data) => {},
            },

            {
                Name: "edit",
                Action: async (id, data) => {

                    const resp = await api.getProject(pid)

                    nav.gotoEditProjectHook(resp.data["ptype"], pid, id)

                },
            },
            {
                Name: "delete",
                Class: "bg-red-400",
                Action: async (id) => {


                    await api.removeProjectHook(pid, id)


                    load();
                },
            },
        ]}
    />
</PageLayout>
