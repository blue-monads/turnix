<script lang="ts">
    import NoteEditor from "./NoteEditor.svelte";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
</script>

<div class="px-2 py-10 md:px-20">
    <NoteEditor
        onSave={async (data) => {
            await api.addNotepad(pid, data);

            goto(`/z/pages/portal/projects/books/${pid}/notepad`);
            
        }}
    />
</div>
