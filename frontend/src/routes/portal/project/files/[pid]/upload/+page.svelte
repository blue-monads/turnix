<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import {
        AppBar,
        FileDropzone,
        getModalStore,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { page } from "$app/stores";
    import Loader from "$lib/compo/loader/loader.svelte";
    import { params } from "$lib/params";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];

    const api = getContext("__api__") as RootAPI;

    let loading = false;
    let file: File;

    const onDrop = async (e: any) => {
        console.log("onDrop", e);
        file = e.target.files[0];
    };
</script>

<FileDropzone on:change={onDrop} name="files">
    <div slot="lead" class="flex justify-center">
        <SvgIcon name="arrow-up-tray" className="w-6 h-6" />
    </div>

    <svelte:fragment slot="message">
        {#if file}
            {file.name}
        {:else}
            <strong>Upload a file</strong> or drag and drop
        {/if}
    </svelte:fragment>
</FileDropzone>

<div class="flex justify-end py-2">
    <button
        class="btn btn-sm variant-filled"
        disabled={!file || loading}
        on:click={async () => {
            loading = true;
            const path = $params["folder"] || "";
            await api.addProjectFile(pid, file.name, path, file);

            loading = false;
            goto(`/z/pages/portal/project/files/${pid}?folder=${path}`);
        }}
    >
        {#if loading}
            Uploading...
        {:else}
            Upload
        {/if}
    </button>
</div>
