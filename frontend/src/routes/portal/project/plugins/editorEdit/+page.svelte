<script lang="ts">
    import { gotoPorjects } from "$lib/nav";
    import { AutoForm, PageLayout, AutoTable, Loader } from "$lib/compo";
    import { params } from "$lib/params";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";

    const api = getContext("__api__") as RootAPI;

    let pid = $params["pid"];
    let id = $params["id"];

    let data = $state({});
    let message = $state("");
    let loading = $state(true);

    const load = async () => {
        loading = true;
        const resp = await api.getProjectPlugin(pid, id);
        if (resp.status !== 200) {
            return;
        }
        data = resp.data;
        loading = false;
    };

    load();
</script>

{#if loading}
    <Loader />
{:else}
    <h1> Plugin Editor </h1>

    
{/if}
