<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import { autocompletion } from "@codemirror/autocomplete";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";
    import EasyGrid from "../../playground/EasyGrid/EasyGrid.svelte";
    import DataExplore from "./DataExplore.svelte";

    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;

    let tables: string[] = $state([]);
    let selectedTable: string | null = $state(null);


    const load = async () => {
        const resp = await rootApi.listProjectTables(pid);
        if (resp.status !== 200) {
            return;
        }

        tables = resp.data;
        selectedTable = tables.at(0);
    };

    load();
</script>

<div class="flex flex-col w-full h-[94vh] p-2">
    {#if tables.length === 0}
        <div class="flex justify-center items-center h-full">
            <span class="text-gray-500">Empty</span>
        </div>
    {:else}
        <TabGroup>
            {#each tables as table}
                <Tab bind:group={selectedTable} name="SQL Query" value={table}>
                    <span>{table}</span>
                </Tab>
            {/each}

            <svelte:fragment slot="panel">
                {#if selectedTable}
                    {#key selectedTable}
                        <DataExplore table={selectedTable} />
                    {/key}
                {/if}
            </svelte:fragment>
        </TabGroup>
    {/if}
</div>
