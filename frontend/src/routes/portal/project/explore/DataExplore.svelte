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

    let { table } = $props();

    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;

    let columns = $state([]);

    const load = async () => {
        const resp = await rootApi.listProjectTableColumns(pid, table);
        if (resp.status !== 200) {
            return;
        }

        columns = resp.data.map((col) => {
            return {
                title: col.name,
                key: col.name,
                type: "text",
            };
        });
    };

    load();

</script>

<EasyGrid
    {columns}
    enablePagination={true}
    enableSidebar={true}
    enableSort={true}
    onLoad={(params) => {
        return [];
    }}
/>
