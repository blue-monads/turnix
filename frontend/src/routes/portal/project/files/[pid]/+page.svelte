<script lang="ts">
    import FileListings from "../../../self/files/FileListings.svelte";
    import { createEventDispatcher, getContext } from "svelte";
    import type { RootAPI, File } from "$lib/api";
    import { page } from "$app/stores";
    import { params } from "$lib/params";

    export let files: File[] = []; //sampleFiles

    export let selected: string;

    const api = getContext("__api__") as RootAPI;
    const pid = $page.params["pid"];

    $: _path = $params["folder"] || "";

    let loading = false;

    const load = async (lpath?: string) => {
        loading = true;
        const resp = await api.listProjectFiles(pid, lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
        loading = false;
    };

    $: load(_path);

</script>

<FileListings
    bind:selected
    baseUrl={`/z/pages/portal/project/files/${pid}`}
    path={_path}
    {files}
    {loading}
/>
