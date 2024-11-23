<script lang="ts">
    import Image from "./Image.svelte";
    import Music from "./Music.svelte";
    import Other from "./Other.svelte";
    import Video from "./Video.svelte";
    import Zip from "./Zip.svelte";

    interface Props {
        name?: string;
        size: string;
    }

    let { name = "", size }: Props = $props();

    const icons = {
        image: Image,
        video: Video,
        music: Music,
        zip: Zip,
        other: Other,
    };

    const exts = {
        mp4: "video",
        jpg: "image",
        png: "image",
        webp: "image",
        mp3: "music",
        zip: "zip",
    };

    let itype = $state("other");
    if (name.includes(".")) {
        itype = exts[name.split(".").pop()] || "other";
    }
</script>

<svelte:component this={icons[itype]} height={size} width={size} />
