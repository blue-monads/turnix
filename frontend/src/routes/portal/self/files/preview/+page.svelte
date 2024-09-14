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


    const loadAsImage = async () => {
        const res = await api.getFileShortKey(fileId);
        if (res.status !== 200) {
            return;
        }


        if (!fileExt) {
            return;
        }


        const img = new Image();
            img.src = api.getFileWithShortKeyURL(res.data);
            img.onload = () => {
                console.log("img loaded");
            };
            img.onerror = (ev) => {
                console.log("img error", ev);
            };

            elemRoot.appendChild(img);
    };

    if (isimage) {
        loadAsImage();
    }

    const download = async () => {
        const res = await api.getFileShortKey(fileId);
        if (res.status !== 200) {
            return;
        }

        const a = document.createElement("a");
        a.href = api.getFileWithShortKeyURL(res.data);
        a.download = filename;

        document.body.appendChild(a);
        a.click();
    }




</script>




<div class="overflow-auto" bind:this={elemRoot}>

    {#if !isimage}
        <button 
        class="btn btn-sm variant-filled" 
        on:click={download}>
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
