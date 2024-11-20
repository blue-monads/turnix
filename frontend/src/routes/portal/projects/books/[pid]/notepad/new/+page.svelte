<script lang="ts">
    import NoteEditor from "./NoteEditor.svelte";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);




    const submit = async () => {

    }

</script>


<div class="px-2 py-10 md:px-20">
    <NoteEditor />
</div>

