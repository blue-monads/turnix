<script lang="ts">
    import { run } from 'svelte/legacy';

    import { getContext } from "svelte";
    import type { RootAPI, File } from "$lib/api";
    import { params } from "$lib/params";
    import FileListings from "./FileListings.svelte";


    interface Props {
        files?: File[];
        selected: string;
    }

    let { files = $bindable([]), selected = $bindable() }: Props = $props();

    const api = getContext("__api__") as RootAPI;

    let _path = $derived($params["folder"] || "");

    let loading = $state(false);

    const load = async (lpath?: string) => {
        loading = true;
        const resp = await api.listSelfFiles(lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
        loading = false;
    };

    $effect.pre(() => {
        load(_path);
    })

</script>

<FileListings
    bind:selected={selected}
    baseUrl="/z/pages/portal/self/files"
    path={_path}
    {files}
    {loading}
/>

