<script lang="ts">
    import NoteEditor from "../new/NoteEditor.svelte";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { goto } from "$app/navigation";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let loading = $state(true);
    let data: any = $state(null);

    const load = async () => {
        loading = true;
        const resp = await api.getNotepad(pid, $params["id"]);
        if (resp.status !== 200) {
            return;
        }
        data = resp.data;
        loading = false;
    };

    load();
</script>

<div class="px-2 py-10 md:px-20">
    {#if loading}
        <Loader />
    {:else}
        <NoteEditor
            {data}
            onSave={async (data) => {
                await api.updateNotepad(pid, $params["id"], data);
                goto(`/z/pages/portal/projects/books/${pid}/notepad`);
            }}
        />
    {/if}
</div>
