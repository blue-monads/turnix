<script lang="ts">
    import {
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import CodeMirror from "svelte-codemirror-editor";
    import {
        sql, SQLite
    } from "@codemirror/lang-sql";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";
    import EasyGrid from "../../playground/EasyGrid/EasyGrid.svelte";
    import type { GridHandle } from "../../playground/EasyGrid/easyTypes";

    let tabSet: number = $state(0);
    let serverCode = $state("select * from __project__Accounts;");

    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;

    interface ColumnMapping {
        sqlName: string;
        gridName: string;
        rendererType: string;
    }

    let columnMappings: ColumnMapping[] = $state([]);
    let lastData = $state([]);
    let mainGridHandle: GridHandle | undefined = $state();
    let mappingGridHandle: GridHandle | undefined = $state();


    // export type RendererType = "text" | "date" | "number" | "currency" | "autocolor" | "relativedate" | "lookup" | string

    const inferColumnMapping = (data: Record<string, any>[]) => {
        const first = data.at(0);
        if (!first) {
            return;
        }

        columnMappings = Object.keys(first).map((key) => ({
            sqlName: key,
            gridName: key,
            rendererType: "text",
        }));
    };

    const finalCoumns = $derived(
        columnMappings.map((cm) => {
            return {
                title: cm.gridName,
                key: cm.sqlName,
                
                rendererType: cm.rendererType,
            };
        })
    );

    $effect(() => {
        console.log("@columnMappings", $state.snapshot(columnMappings));
        console.log("@finalCoumns", $state.snapshot(finalCoumns));
    });
</script>

<div class="flex flex-col w-full h-[94vh] p-2">
    <div class="w-full border overflow-auto resize-y p-2 h-24 bg-white rounded">
        <CodeMirror
            bind:value={serverCode}
            extensions={[]}
            lang={sql({
                dialect: SQLite,
            })}
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
                            console.log("@handle", mainGridHandle);
                            mainGridHandle?.reload();
                        }}
                    >
                        <SvgIcon className="w-4 h-4 mt-1" name="play" />
                        Run
                    </button>
                </div>

                <svelte:fragment slot="panel">
                    <div class="min-h-20 overflow-auto">
                        {#if tabSet === 0}
                            {#key finalCoumns}
                                <EasyGrid
                                    key="id"
                                    bind:handle={mainGridHandle}
                                    columns={finalCoumns}
                                    enableStartAutoLoad={false}
                                    enablePagination={false}
                                    enableSidebar={false}
                                    enableSort={false}
                                    enableFilter={false}
                                    initialData={lastData}
                                    onLoad={async (params) => {
                                        if (!serverCode) {
                                            return [];
                                        }

                                        const resp = await rootApi.runProjectSQL2(
                                            pid,
                                            {
                                                qstr: serverCode,
                                                data: [],
                                            }
                                        );

                                        if (resp.status !== 200) {
                                            return [];
                                        }

                                        if (finalCoumns.length === 0) {
                                            inferColumnMapping(resp.data);
                                        }


                                        lastData = resp.data;
                                        return lastData;
                                    }}
                                />
                            {/key}
                        {:else if tabSet === 1}
                            <EasyGrid
                                key=""
                                bind:handle={mappingGridHandle}
                                columns={[
                                    {
                                        title: "SQL Name",
                                        key: "sqlName",
                                    },
                                    {
                                        title: "Grid Name",
                                        key: "gridName",
                                    },
                                    {
                                        title: "Renderer Type",
                                        key: "rendererType",
                                        rendererOptions: {
                                            autoColor: true,
                                        },
                                    },
                                ]}
                                enablePagination={false}
                                enableSidebar={false}
                                enableSort={false}
                                onLoad={async (params) => {
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
                                        action: () => {},
                                    },
                                ]}
                            />
                        {/if}
                    </div>
                </svelte:fragment>
            </TabGroup>
        </div>
    </div>
</div>
