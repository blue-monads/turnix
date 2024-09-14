<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { params } from "$lib/params";

    const api = getContext("__api__") as RootAPI;

    let loading = false;

    const fileId = $params["fid"];
    const path = $params["folder"] || "";
    const filename = $params["filename"];

    let elemRoot: HTMLElement;

    const imageExts = ["png", "jpg", "jpeg", "gif", "bmp", "webp", "svg"];
    const fileExt = filename.split(".").pop();
    const isimage = imageExts.includes(fileExt as string);
    const isVideo = fileExt === "mp4";

    let fileurlWithShortKey = "";

    const load = async () => {
        loading = true;
        const res = await api.getFileShortKey(fileId);
        if (res.status !== 200) {
            return;
        }
        fileurlWithShortKey = api.getFileWithShortKeyURL(res.data);
        loading = false;
    };

    load();

    const download = async () => {
        const a = document.createElement("a");
        a.href = fileurlWithShortKey;
        a.download = filename;

        document.body.appendChild(a);
        a.click();
    };
</script>

<div
    class="h-inherit w-inherit overflow-auto"
    bind:this={elemRoot}
>
    {#if loading}
        <div>loading...</div>
    {:else if isimage}
        <img
            class="h-fit w-auto overflow-auto"
            src={fileurlWithShortKey}
            alt={filename}
        />
    {:else if isVideo}
        <!-- svelte-ignore a11y-media-has-caption -->
        <video controls>
            <source src={fileurlWithShortKey} type="video/mp4" />
        </video>
    {:else}
        <button class="btn btn-sm variant-filled" on:click={download}>
            download
        </button>
    {/if}
</div>

<style>
    div {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>
