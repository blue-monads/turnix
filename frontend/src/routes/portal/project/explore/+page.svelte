<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import { autocompletion } from "@codemirror/autocomplete";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";
    import EasyGrid from "../../playground/EasyGrid/EasyGrid.svelte";

    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);
</script>

<div class="flex flex-col w-full h-[94vh] p-2">
    <EasyGrid
        columns={[]}
        enablePagination={true}
        enableSidebar={true}
        enableSort={true}
        onLoad={(params) => {
            return [];
        }}
    />
</div>
