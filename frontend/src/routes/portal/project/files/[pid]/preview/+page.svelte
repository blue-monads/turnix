<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { params } from "$lib/params";
    import { page } from "$app/stores";

    const api = getContext("__api__") as RootAPI;

    const fileId = $params["fid"];
    const path = $params["folder"] || "";
    const filename = $params["filename"];
    const pid = $page.params["pid"];

    let elemRoot: HTMLElement;

    const imageExts = ["png", "jpg", "jpeg", "gif", "bmp", "webp", "svg"];

    const load = async () => {
        const res = await api.getProjectFile(pid, fileId);
        if (res.status !== 200) {
            return;
        }

        const fileExt = filename.split(".").pop();

        if (!fileExt) {
            return;
        }

        if (imageExts.includes(fileExt)) {
            console.log("image");

            const img = new Image();
            img.src = URL.createObjectURL(res.data);
            img.onload = () => {
                console.log("img loaded");
            };
            img.onerror = (ev) => {
                console.log("img error", ev);
            };

            elemRoot.appendChild(img);
        } else {
            // download as file button

            const a = document.createElement("a");
            a.href = URL.createObjectURL(res.data);
            a.download = filename;
            a.innerText = "download";
            a.style.color = "while";
            a.style.padding = "5%";
            a.style.paddingLeft = "10%";
            a.style.textDecoration = "underline";

            elemRoot.appendChild(a);
        }
    };

    load();
</script>

<div class="overflow-auto" bind:this={elemRoot}>
    <a href="/" class="hidden">a</a>
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
