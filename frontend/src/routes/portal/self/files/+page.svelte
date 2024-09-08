<script lang="ts">
    import { getContext } from "svelte";
    import type { RootAPI, File } from "$lib/api";
    import { params } from "$lib/params";
    import FileListings from "./FileListings.svelte";

    export let files: File[] = []; //sampleFiles

    export let selected: string;

    const api = getContext("__api__") as RootAPI;

    $: _path = $params["folder"] || "";

    let loading = false;

    const load = async (lpath?: string) => {
        loading = true;
        const resp = await api.listSelfFiles(lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
        loading = false;
    };

    $: load(_path);

</script>

<FileListings
    bind:selected={selected}
    baseUrl="/z/pages/portal/self/files"
    path={_path}
    {files}
    {loading}
/>

