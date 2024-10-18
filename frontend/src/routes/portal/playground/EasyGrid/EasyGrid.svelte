<script lang="ts">
    import type { GridOptions } from "./easyTypes";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    let { columns = [], onLoad, actions, key }: GridOptions = $props();

    let datas = $state([]);

    let sortMode: "asc" | "desc" = $state("asc");
    let sortKey: string = $state("");
    let minId: number = $state(0);
    let maxId: number = $state(0);

    let filterModels: Record<string, any> = $state({});

    const loadData = async (loadType: "next" | "initial" | "previous") => {
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
    };

    $effect(() => {
        loadData("initial");
    });
</script>

<div class="p-1 overflow-auto card">
    <table class="table-auto border-collapse w-full">
        <thead>
            <tr class="rounded-lg text-sm font-medium text-gray-700 text-left">
                {#each columns as column}
                    <th class="px-2 py-1">
                        <div class="flex gap-2">
                            <button
                                class="flex gap-1"
                                onclick={() => {
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

                                    sortKey = column.key;
                                    sortMode = "asc";
                                }}
                            >
                                <span>
                                    {column.title}
                                </span>

                                {#if column.key === sortKey}
                                    {#if sortMode === "asc"}
                                        <SvgIcon
                                            name="chevron-up"
                                            className="h-4 w-4"
                                        />
                                    {:else}
                                        <SvgIcon
                                            name="chevron-down"
                                            className="h-4 w-4"
                                        />
                                    {/if}
                                {/if}
                            </button>

                            <div>
                                <button onclick={() => {}}>
                                    <SvgIcon
                                        name="bars-3-bottom-right"
                                        className="h-4 w-4"
                                    />
                                </button>
                            </div>
                        </div>
                    </th>
                {/each}

                {#if actions}
                    <th class="px-2 py-2"> Actions </th>
                {/if}
            </tr>
        </thead>
        <tbody class="text-sm font-normal text-gray-700">
            {#each datas as data}
                <tr
                    class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                >
                    {#each columns as column}
                        <td class="px-2 py-2">
                            <span class="p-1 rounded-lg"
                                >{data[column.key] || ""}
                            </span>
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
        </tbody>
    </table>
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
</div>
