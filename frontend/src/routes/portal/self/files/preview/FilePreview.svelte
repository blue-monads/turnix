<script lang="ts">
    import { RootAPI } from "$lib/api";


    let loading = false;

    export let fileId: string;
    export let filename: string;
    export let api: RootAPI;


    let elemRoot: HTMLElement;

    const imageExts = ["png", "jpg", "jpeg", "gif", "bmp", "webp", "svg"];
    const fileExt = filename.split(".").pop();

    let fileType: "image" | "video" | "audio" | "other" = imageExts.includes(
        fileExt as string,
    )
        ? "image"
        : fileExt === "mp4"
        ? "video"
        : fileExt === "mp3"
        ? "audio"
        : "other";




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
    {:else if fileType === "image"}
        <img
            class="h-fit w-auto overflow-auto"
            src={fileurlWithShortKey}
            alt={filename}
        />
    {:else if fileType === "video"}
        <!-- svelte-ignore a11y-media-has-caption -->
        <video controls>
            <source src={fileurlWithShortKey} type="video/mp4" />
        </video>
    {:else if fileType === "audio"}
        <!-- svelte-ignore a11y-media-has-caption -->
        <audio controls>
            <source src={fileurlWithShortKey} type="audio/mp3" />
        </audio>
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
