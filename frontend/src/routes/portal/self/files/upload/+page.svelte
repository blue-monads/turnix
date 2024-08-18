<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import {
        FileDropzone,
    } from "@skeletonlabs/skeleton";
    import { params } from "$lib/params";

    const api = getContext("__api__") as RootAPI;

    let loading = false;
    let file: File;

    const onDrop = async (e: any) => {
        console.log("onDrop", e);
        file = e.target.files[0];
    };
</script>

<FileDropzone on:change={onDrop} name="files">
    <svelte:fragment slot="lead">
        <span>{file?.name || ""}</span>
    </svelte:fragment>
    <svelte:fragment slot="message">Upload files</svelte:fragment>
</FileDropzone>

<div class="flex justify-end py-2">
    <button
        class="btn btn-sm variant-filled"
        disabled={!file || loading}
        on:click={async () => {
            loading = true;
            const path = $params["folder"] || "";
            await api.addSelfFile(file.name, path, file);

            loading = false;
        }}
    >
        {#if loading}
            Uploading...
        {:else}
            Upload
        {/if}
    </button>
</div>
