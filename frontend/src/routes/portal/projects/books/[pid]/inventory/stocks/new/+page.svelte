<script lang="ts">
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import StockInMaker from "./StockInMaker.svelte";
    import type { RootAPI } from "$lib";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

</script>

<StockInMaker 
    pid={Number(pid)} 
    api={api} 
    submit={async (data) => {
    return "OK";
}} />