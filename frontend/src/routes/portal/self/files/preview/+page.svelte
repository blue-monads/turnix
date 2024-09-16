<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { params } from "$lib/params";
    import FilePreview from "./FilePreview.svelte";

    const api = getContext("__api__") as RootAPI;

    const fileId = $params["fid"];
    const path = $params["folder"] || "";
    const filename = $params["filename"];
</script>

<FilePreview {fileId} {filename} {api} />
