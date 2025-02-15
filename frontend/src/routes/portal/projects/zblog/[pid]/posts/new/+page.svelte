<script lang="ts">
    import PostEditor from "../edit/PostEditor.svelte";
    import { page } from "$app/stores";
    import { NewZblogAPI } from "../../../zblogApi";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];

    const api = NewZblogAPI(getContext("__api__") as RootAPI);

    let content = "";
    let title = "";
    let slug = "";

    const save = async () => {
        await api.addPost(pid, { title, slug, content });
        goto(`/z/pages/portal/projects/zblog/${pid}`);
    };


</script>

<div class="p-4 flex flex-col gap-4">
    <h1 class="text-2xl font-bold">New Post</h1>

    <label class="label">
        <span>Title</span>
        <input type="text" class="input" bind:value={title} />
    </label>

    <label class="label">
        <span>Slug</span>
        <input type="text" class="input" bind:value={slug} />
    </label>

    <label class="label">
        <span>Content</span>
        <PostEditor bind:content />
    </label>

    <div class="flex justify-end">
        <button 
        class="btn btn-sm variant-filled"
        onclick={save}
        >Save
    </button>
    </div>
</div>
