<script lang="ts">
    import { AppBar } from '@skeletonlabs/skeleton';
    import {  onMount } from "svelte";
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


<AppBar> 
    <svelte:fragment slot="lead">
        <h2 class="h3">Projects</h2>
    </svelte:fragment>
    
    <svelte:fragment slot="trail">
        <button type="button" class="btn variant-filled-primary text-secondary-50">Add</button>
    </svelte:fragment>
</AppBar>


<Autotable
    key_names={[
        ["name", "Name"],
        ["ptype", "Project type"],
        ["owner", "Owner"],
    ]}
    {datas}
    actions={[]}
/>
