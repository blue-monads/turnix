<script lang="ts">
    import type { GridOptions, OperatorValue, FilterModel } from "./easyTypes";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import FilterPanel from "./FilterPanel.svelte";

    let {
        columns = [],
        onLoad,
        actions,
        key,
        enableSort,
        enableSidebar = true
    }: GridOptions = $props();

    let datas = $state([]);

    let sortMode: "asc" | "desc" = $state("asc");
    let sortKey: string = $state(undefined);
    let minId: number = $state(0);
    let maxId: number = $state(0);
    let loading = $state(false);
    let sidePanelMode: "none" | "columns" | "filters" | "charts" = $state("none");

    let filterModels: Record<string, FilterModel> = $state({});

    let activeFilter = $state(null);
    let filterPanelPosition = $state({ top: 0, left: 0 });

    let filterPanelRef: HTMLDivElement;

    const hashSeed: number = 74

    


    const hashCode = ( str: string) => {
        let hash = hashSeed;
        for (var i = 0; i < str.length; i++) {
            hash = str.charCodeAt(i) + ((hash << 5) - hash);
        }
        return hash;
    };
    const toggleFilterPanel = (column, event) => {
        if (activeFilter === column.key) {
            activeFilter = null;
        } else {
            activeFilter = column.key;
            setTimeout(() => {
                if (filterPanelRef) {
                    const rect = event.target.getBoundingClientRect();
                    const panelRect = filterPanelRef.getBoundingClientRect();

                    let left = rect.left + window.scrollX;
                    const rightEdge = left + panelRect.width;
                    const viewportWidth = window.innerWidth;

                    if (rightEdge > viewportWidth) {
                        left = Math.max(0, viewportWidth - panelRect.width);
                    }

                    filterPanelPosition = {
                        top: `${rect.bottom + window.scrollY}px`,
                        left: `${left}px`,
                    };
                }
            }, 0);
        }
    };

    const closeFilterPanel = () => {
        activeFilter = null;
        filterPanelPosition = { top: 0, left: 0 };
    };

    const loadData = async (loadType: "next" | "initial" | "previous") => {
        loading = true;

        const nextDatas = await onLoad({
            loadType,
            orderBy: sortKey,
            orderDirection: sortMode,
            minId: minId,
            maxId: maxId,
        });

        let nextMinId = 0;
        let nextMaxId = 0;

        if (nextDatas.length > 0) {
            let first = nextDatas[0][key];
            nextMaxId = first;
            nextMinId = first;

            nextDatas.forEach((item) => {
                const currItem = item[key];

                if (currItem < minId) {
                    nextMinId = item.id;
                } else if (currItem > maxId) {
                    nextMaxId = item.id;
                }
            });
        }

        minId = nextMinId;
        maxId = nextMaxId;
        datas = nextDatas;
        loading = false;
    };

    $effect(() => {
        loadData("initial");
    });

    $effect(() => {
        const handleClickOutside = (event) => {
            if (filterPanelRef && !filterPanelRef.contains(event.target)) {
                closeFilterPanel();
            }
        };

        document.addEventListener("click", handleClickOutside);

        return () => {
            document.removeEventListener("click", handleClickOutside);
        };
    });

    const sidebarToggle = (name: string) => () => {
        if (sidePanelMode === name) {
            sidePanelMode = "none";
        } else {
            sidePanelMode = name;
        }
    }


</script>

<div class="p-1 overflow-auto card">
    <div class="flex">
        <div class="flex-1">           
            <table class="table-auto border-collapse w-full relative">
                <thead>
                    <tr class="rounded-lg text-sm font-medium text-gray-700 text-left">
                        {#each columns as column}

                            <th class="px-2 py-1">
                                <div class="flex gap-2">
                                    <button
                                        class="flex gap-1"
                                        onclick={() => {
                                            if (!enableSort || column.disableSort) {
                                                return;
                                            }

                                            if (sortKey === column.key) {
                                                if (sortMode === "asc") {
                                                    sortMode = "desc";
                                                } else {
                                                    sortKey = "";
                                                }
                                            } else {
                                                sortKey = column.key;
                                                sortMode = "asc";
                                            }
                                        }}
                                    >
                                        <span>
                                            {column.title}
                                        </span>

                                        {#if column.key === sortKey}
                                            {#if sortMode === "asc"}
                                                <SvgIcon
                                                    name="chevron-down"
                                                    className="h-4 w-4"
                                                />
                                            {:else}
                                                <SvgIcon
                                                    name="chevron-up"
                                                    className="h-4 w-4"
                                                />
                                            {/if}
                                        {/if}
                                    </button>

                                    <div>
                                        <button
                                            class="rounded p-0.5 relative {filterModels[column.key] ? 'bg-blue-100' : ''}"
                                            onclick={(event) => {
                                                console.log("@col/ev", column);
                                                event.stopPropagation();
                                                toggleFilterPanel(column, event);
                                            }}
                                        >

                                        {#if filterModels[column.key]}
                                            <span class="rounded-full absolute top-1 -right-1 w-2 h-2 bg-blue-400 shadow"></span>
                                        {/if}


                                            <SvgIcon
                                                name="bars-3-bottom-right"
                                                className="h-4 w-4"
                                            />
                                        </button>
                                    </div>

                                    {#if activeFilter}
                                        <div
                                            bind:this={filterPanelRef}
                                            class="fixed rounded bg-white shadow-lg z-10 min-w-64 resize p-2 border-b"
                                            style="top: {filterPanelPosition.top}; left: {filterPanelPosition.left};"
                                        >
                                            <FilterPanel
                                                operator={filterModels[activeFilter]?.operator}
                                                value={filterModels[activeFilter]?.value}
                                                closeFilter={closeFilterPanel}
                                                clearFilter={() => {
                                                    filterModels[activeFilter] = null;
                                                    closeFilterPanel();
                                                }}
                                                column={columns.find(col => col.key === activeFilter)}
                                                applyFilter={(opval) => {
                                                    const col = opval.column;
                                                    console.log("@col", col);

                                                    const filterModel = filterModels[col.key] || {};
                                                    filterModels[col.key] = {
                                                        key: col.key,
                                                        fiterType:
                                                        col.type || "text",
                                                        ...filterModel,
                                                        ...opval,
                                                    };

                                                    loadData("initial");

                                                    console.log($state.snapshot(filterModels));

                                                    closeFilterPanel();
                                                }}
                                            />
                                        </div>
                                    {/if}
                                </div>
                            </th>
                        {/each}

                        {#if actions}
                            <th class="px-2 py-2"> Actions </th>
                        {/if}
                    </tr>
                </thead>
                <tbody class="text-sm font-normal text-gray-700">
                    {#if loading}
                        <tr
                            class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                        >
                            <td colspan={columns.length} class="px-2 py-2">
                                <span class="p-1 rounded-lg">Loading...</span>
                            </td>
                        </tr>
                    {:else}
                        {#each datas as data}
                            <tr
                                class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                            >
                                {#each columns as column}
                                {@const dataValue = data[column.key]}
                                    <td class="px-2 py-2">
                                        
                                        {#if dataValue }
                                            <span 
                                            class="p-1 rounded-lg {column.cssClasses || ""}"
                                            style={`background-color: ${column.rendererOptions?.autoColor ? `hsl(${hashCode(dataValue) % 360}, 100%, 80%)` : ""}`}
                                        
                                        >
                                                {#if column.type === "date" && dataValue }
                                                    {new Date(dataValue).toLocaleDateString()}
                                                {:else if column.type === "number"}
                                                    <span class="pill">{dataValue}</span>                                        
                                                {:else if columns.type === "lookup" && columns.rendererOptions?.mappings} 
                                                    {columns.rendererOptions.mappings[dataValue] || ""}
                                                {:else}
                                                    {data[column.key] || ""}
                                                {/if}

                                            </span>
                                            {/if}
                                            
                                




                                    </td>

                                    {#if actions}
                                        <td class="px-2 py-2 flex flex-row">
                                            {#each actions as action}
                                                <button
                                                    onclick={() =>
                                                        action.action(
                                                            data[action.name],
                                                            data
                                                        )}
                                                    class="flex m-1 text-white transform hover:scale-110 btn btn-sm {action.Class ||
                                                        'bg-blue-400'}"
                                                >
                                                    {#if action.icon}
                                                        <SvgIcon
                                                            name={action.icon}
                                                            className="h-5 w-5"
                                                        />
                                                    {/if}

                                                    {action.name}</button
                                                >
                                            {/each}
                                        </td>
                                    {/if}
                                {/each}
                            </tr>
                        {/each}
                    {/if}
                </tbody>
            </table>
        </div>

        {#if enableSidebar} 
            {#if sidePanelMode === "columns"}
                <div class="flex min-w-60 flex-col border-l border-gray-200 p-1 resize-x overflow-auto">
                    i am sidebar 
                </div>
            {/if}

                <div class="flex w-6 flex-col border-l border-gray-200 p-1">
                    
                    <div class="flex justify-start gap-4 rotate-90 ">
                        <button class="hover:bg-gray-50 px-1"
                        onclick={sidebarToggle("columns")}
                        
                        >
                            Columns
                        </button>

                        <button
                        class="hover:bg-gray-50 px-1"
                        onclick={sidebarToggle("columns")}
                    
                        >
                            Filters
                        </button>
                        <button
                        class="hover:bg-gray-50 px-1"
                        onclick={sidebarToggle("columns")}
                    
                        >
                            Charts
                        </button>
                        
                    </div>
                </div>
        {/if}

</div>





    {#if !loading}
        <div class="flex justify-end gap-2">
            <select class="rounded w-20 border-gray-300 border pl-2">
                <option>10</option>
                <option>20</option>
                <option>50</option>
            </select>

            <div class="flex gap-2">
                <button
                    class="btn btn-sm bg-gray-100"
                    onCLick={async () => {
                        loadData("previous");
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="chevron-left" />
                </button>

                <button
                    class="btn btn-sm bg-gray-100"
                    onCLick={async () => {
                        loadData("next");
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="chevron-right" />
                </button>
            </div>
        </div>
    {/if}
</div>
