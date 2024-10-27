<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import CodeMirror from "svelte-codemirror-editor";
    import { html } from "@codemirror/lang-html";
    import {
        javascript,
        localCompletionSource,
        scopeCompletionSource,
    } from "@codemirror/lang-javascript";
    import { autocompletion } from "@codemirror/autocomplete";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";
    import EasyGrid from "../../playground/EasyGrid/EasyGrid.svelte";
    import type { GridHandle } from "../../playground/EasyGrid/easyTypes";

    let tabSet: number = $state(0);
    let serverCode = $state("");


    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;

    interface ColumnMapping {
        sqlName: string;
        gridName: string;
        rendererType: string;
    }

    let columnMappings: ColumnMapping[] = $state([]);
    let lastData = $state([]);
    let handle: GridHandle = $state();

    // export type RendererType = "text" | "date" | "number" | "currency" | "autocolor" | "relativedate" | "lookup" | string
</script>

<div class="flex flex-col w-full h-[94vh] p-2">
    <div class="w-full border overflow-auto resize-y p-2 h-24 bg-white rounded">
        <CodeMirror
            bind:value={serverCode}
            extensions={[]}
            lang={javascript()}
        />
    </div>

    <div class="w-full border  p-2">
        <div class="flex-1 w-full">
            <TabGroup>
                <Tab bind:group={tabSet} name="SQL Query" value={0}>
                    <span>Grid</span>
                </Tab>
                <Tab bind:group={tabSet} name="UI" value={1}>
                    <span>Mappings</span>
                </Tab>

                <div class="flex-1" />

                <div class="flex justify-end p-2">
                    <button
                        class="inline-flex rounded p.05 variant-filled self-center"
                        onclick={() => {
                            console.log("@handle", handle);

                            handle.reload();

                        }}
                    >
                        <SvgIcon className="w-4 h-4 mt-1" name="play" />
                        Run
                    </button>
                </div>

                <svelte:fragment slot="panel">
                    <div class="max-h-[40vh] md:max-h-[90vh] overflow-auto">
                        {#if tabSet === 0}
                            <EasyGrid
                                bind:handle
                                columns={[]}
                                enableStartAutoLoad={false}
                                enablePagination={false}
                                enableSidebar={false}
                                enableSort={false}
                                onLoad={async (params) => {
                                    if (!serverCode) {
                                        return [];
                                    }

                                    const resp = await rootApi.runProjectSQL(pid, serverCode);
                                    if (resp.status !== 200) {
                                        return [];
                                    }

                                    lastData = resp.data;
                                    return lastData;
                                }}
                            />
                        {:else if tabSet === 1}
                            <EasyGrid
                                enableStartAutoLoad={false}
                                columns={[
                                    {
                                        title: "SQL Name",
                                        key: "sqlName",
                                        type: "text",
                                    },
                                    {
                                        title: "Grid Name",
                                        key: "gridName",
                                        type: "text",
                                    },
                                    {
                                        title: "Renderer Type",
                                        key: "rendererType",
                                        type: "text",
                                        rendererOptions: {
                                            autoColor: true,
                                        },
                                    },
                                ]}
                                enablePagination={false}
                                enableSidebar={false}
                                enableSort={false}
                                onLoad={(params) => {
                                    return columnMappings;
                                }}
                                headerActions={[
                                    {
                                        name: "Add",
                                        action: () => {
                                            columnMappings = [
                                                ...columnMappings,
                                                {
                                                    sqlName: "",
                                                    gridName: "",
                                                    rendererType: "text",
                                                },
                                            ];
                                        },
                                    },
                                    {
                                        name: "Infer",
                                        action: () => {


                                        },
                                    },
                                ]}


                            />
                        {/if}
                    </div>
                </svelte:fragment>
            </TabGroup>

            <!-- <div>AAA</div> -->
        </div>
    </div>
</div>
