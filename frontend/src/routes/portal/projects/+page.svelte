<script lang="ts">
    import Autotable from "../../../compo/autotable/autotable.svelte";
    import PageLayout from "../../../compo/pagelayout/pagelayout.svelte";
    import { onMount } from "svelte";
    import * as nav from "$lib/nav";

    import { api } from "$lib/api";

    let datas: any[] = [];

    const load = async () => {
        const resp = await api.listProjects();
        if (resp.status !== 200) {
            return;
        }
        datas = resp.data;
    };

    onMount(() => {
        load();
    });
</script>

<PageLayout
    title="projects"
    actions={[{ name: "add", actionFn: nav.gotoAddProject }]}
>
    <Autotable
        key_names={[
            ["name", "Name"],
            ["ptype", "Project type"],
            ["owner", "Owner"],
        ]}
        {datas}
        actions={[]}
    />
</PageLayout>
