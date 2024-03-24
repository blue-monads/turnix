<script lang="ts">
    import type { API } from "$lib/api/api";
    import { getContext, onMount } from "svelte";
    import Autotable from "../../../compo/autotable/autotable.svelte";
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

<Autotable key_names={[]} {datas} actions={[]} />

<div>Projects</div>
