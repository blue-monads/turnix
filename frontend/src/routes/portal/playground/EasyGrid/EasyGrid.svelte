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
        headerActions,
        handle = $bindable(),
        enableStartAutoLoad = true,
        enableSidebar = false,
        enablePagination = true,
        enableFilter = true,
        initialData
    }: GridOptions = $props();

    let datas = $state([]);

    let sortMode: "asc" | "desc" = $state("asc");
    let sortKey: string | undefined = $state(undefined);
    let minId: number = $state(0);
    let maxId: number = $state(0);
    let loading = $state(false);
    let sidePanelMode: "none" | "columns" | "filters" | "charts" = $state("none");
    let recordLoadCount = $state(25);

    let filterModels: Record<string, FilterModel[]> = $state({});
    let enabledColumns: string[] = $state(columns.map(col => col.key));
    let loadedColumns: string[] = $state([]);

    let activeFilter = $state(null);
    let filterPanelPosition = $state({ top: 0, left: 0 });
    let needsReload = $state(false);

    let filterPanelRef: HTMLDivElement;

    const hashSeed: number = 74

    handle = {
        "reload": () => loadData("initial"),
        "next": () => loadData("next"),
        "previous": () => loadData("previous"),
    }


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
        needsReload = false;
        let activeColumns = $state.snapshot(enabledColumns)

        let nextDatas = await onLoad({
            loadType,
            orderBy: sortKey,
            orderDirection: sortMode,
            filterModels: $state.snapshot(filterModels),
            activeColumns,
            minId: minId,            
            maxId: maxId,
            pageSize: recordLoadCount,
        });

        if (!nextDatas) {
            nextDatas = [];
        }

        setData(nextDatas);
        loading = false;
        loadedColumns = activeColumns;

    };

    const setData = (nextDatas: Record<string, any>[]) => {
        let nextMinId = 0;
        let nextMaxId = 0;

        if (nextDatas.length > 0) {
            let first = nextDatas[0][key];
            nextMaxId = first;
            nextMinId = first;

            nextDatas.forEach((item) => {
                const currItem = item[key];

                if (currItem < minId) {
                    nextMinId = currItem;
                } 
                
                if (currItem > maxId) {
                    nextMaxId = currItem;
                }
            });
        }

        minId = nextMinId;
        maxId = nextMaxId;
        datas = nextDatas;
    }

    if (enableStartAutoLoad) {
        loadData("initial");
    } else {
        loadedColumns = $state.snapshot(enabledColumns);
        setData(initialData || []);
    }

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

<div class="p-1 card">
    <div class="flex">
        <div class="flex-1 min-h-32 overflow-auto"> 
            
        {#if columns.length === 0}
            <div class="flex justify-center items-center h-full">
                <span class="text-gray-500">Empty</span>
            </div>
        {:else if loading}
            <div class="flex justify-center items-center h-full">
                <span class="text-gray-500">Loading...</span>
            </div>
        {:else}            
            <table class="table-auto border-collapse relative w-full">
                <thead>
                    <tr class="rounded-lg text-sm font-medium text-gray-700 text-left">
                        {#each columns as column}

                            {#if loadedColumns.includes(column.key) && enabledColumns.includes(column.key)}
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


                                        {#if enableFilter}
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
                                        {/if}



                                        {#if activeFilter}
                                            <div
                                                bind:this={filterPanelRef}
                                                class="fixed rounded bg-white shadow-lg z-10 min-w-64 resize p-2 border-b"
                                                style="top: {filterPanelPosition.top}; left: {filterPanelPosition.left};"
                                            >
                                                <FilterPanel
                                                    currentFilterModel={filterModels[activeFilter]}
                                                    closeFilter={closeFilterPanel}
                                                    clearFilter={() => {
                                                        filterModels[activeFilter] = null;
                                                        needsReload = true;
                                                        closeFilterPanel();
                                                    }}
                                                    column={columns.find(col => col.key === activeFilter)}
                                                    applyFilter={(column, filters) => {

                                                        filterModels[column.key] = filters;

                                                        needsReload = true;

                                                        console.log($state.snapshot(filterModels));

                                                        closeFilterPanel();
                                                    }}
                                                />
                                            </div>
                                        {/if}
                                    </div>
                                </th>
                            {/if}
                        {/each}

                        {#if actions || headerActions}
                            <th class="px-2 py-2 flex"> 
                                {#if headerActions}
                                    {#each headerActions as ha }
                                        <button
                                        class="flex m-1 btn btn-sm bg-gray-50 {ha.Class }"
                                            onclick={() => {
                                                ha.action();
                                            }}
                                        >
                                            {#if ha.icon}
                                                <SvgIcon
                                                    name={ha.icon}
                                                    className="h-5 w-5"
                                                />
                                            {/if}

                                            {ha.name}
                                            </button>
                                        
                                    {/each}
                                {:else}
                                    Actions

                                {/if}
                            </th>
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
                    {:else if datas.length === 0}
                        <tr
                            class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                        >
                            <td colspan={columns.length} class="px-2 py-2">
                                <span class="p-1 rounded-lg">Empty</span>
                            </td>
                        </tr>
                    {:else}
                        {#each datas as data}
                            <tr
                                class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                            >
                                {#each columns as column}
                                {@const dataValue = data[column.key]}


                                {#if loadedColumns.includes(column.key) && enabledColumns.includes(column.key)}

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

                                    {/if}
                                {/each}
                            </tr>
                        {/each}
                    {/if}
                </tbody>
            </table>
            {/if}

        </div>

        {#if enableSidebar} 
            {#if sidePanelMode != "none"}
                <div class="flex min-w-48 flex-col border-l border-gray-200 p-1 resize-x overflow-auto">

                    {#if sidePanelMode === "columns"}
                        <div class="flex flex-col gap-2 p-2">
                            {#each columns as column}
                                <div class="flex">
                                    <label class="flex items-center gap-2">
                                        <input 
                                        type="checkbox" 
                                        checked={enabledColumns.includes(column.key)} 
                                        class="h-4 w-4 border-gray-300 rounded"
                                        onchange={() => {
                                            if (enabledColumns.includes(column.key)) {
                                                enabledColumns = enabledColumns.filter(col => col !== column.key);
                                            } else {
                                                enabledColumns = [...enabledColumns, column.key];
                                                if (!loadedColumns.includes(column.key)) {
                                                    needsReload = true;
                                                }
                                            }
                                        }}                                        
                                        >
                                        <span>{column.title}</span>
                                    </label>                                    
                                </div>
                            {/each}
                        </div>
                    {/if}

                </div>
            {/if}

                <div class="flex w-6 flex-col border-l border-gray-200 p-1">
                    
                    <div class="flex justify-start gap-4 rotate-90 ">
                        <button class="hover:bg-gray-50 px-1"
                        onclick={sidebarToggle("columns")}
                        
                        >
                            Columns
                        </button>

                        <!-- <button
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
                        </button> -->
                        
                    </div>
                </div>
        {/if}



</div>



    {#if !loading && enablePagination}
        <div class="flex justify-end gap-2">

            <button 
                class="btn btn-sm  { needsReload ? 'bg-blue-200' : 'bg-gray-50'}"
                onclick={async () => {
                    loadData("initial");
                }}
                >
                <SvgIcon className="h-4 w-4" name="arrow-path" />
            </button>

            <select 
                class="rounded w-20 border-gray-300 border pl-2"
                value={recordLoadCount + ""}
                onchange={(ev) => {
                    recordLoadCount = Number(ev.target.value);
                    needsReload = true;
                    console.log("@ev", ev.target.value);
                    console.log($state.snapshot(recordLoadCount));
                }}
                
                >
                <option value="25">25</option>
                <option value="50">50</option>
                <option value="100">100</option>
                <option value="250">250</option>
            </select>

            <div class="flex gap-2">
                <button
                    class="btn btn-sm bg-gray-100"
                    onclick={async () => {
                        loadData("previous");
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="chevron-left" />
                </button>

                <button
                    class="btn btn-sm bg-gray-100"
                    onclick={async () => {
                        loadData("next");
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="chevron-right" />
                </button>
            </div>
        </div>
    {/if}
</div>
