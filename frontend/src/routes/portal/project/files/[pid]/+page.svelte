<script lang="ts">
    import { run } from 'svelte/legacy';

    import FileListings from "../../../self/files/FileListings.svelte";
    import { getContext } from "svelte";
    import type { RootAPI, File } from "$lib/api";
    import { page } from "$app/stores";
    import { params } from "$lib/params";


    interface Props {
        files?: File[];
        selected: string;
    }

    let { files = $bindable([]), selected = $bindable() }: Props = $props();

    const api = getContext("__api__") as RootAPI;
    const pid = $page.params["pid"];

    let _path = $derived($params["folder"] || "");

    let loading = $state(false);

    const load = async (lpath?: string) => {
        loading = true;
        const resp = await api.listProjectFiles(pid, lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
        loading = false;
    };

    run(() => {
        load(_path);
    });

</script>

<FileListings
    bind:selected
    baseUrl={`/z/pages/portal/project/files/${pid}`}
    path={_path}
    {files}
    {loading}
/>
