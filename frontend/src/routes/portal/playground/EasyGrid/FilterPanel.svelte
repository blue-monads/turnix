<script lang="ts">
    import { OPERATORS } from "./easyTypes";
    import type { Column, FilterModel } from "./easyTypes";

    interface FilterPanelProps {
        closeFilter: () => void;
        clearFilter: () => void;
        column: Column;
        applyFilter: (filterModel: FilterModel[]) => void;
        currentFilterModel?: FilterModel[];
    }

    let {
        closeFilter,
        applyFilter,
        clearFilter,
        column,
        currentFilterModel,
    } = $props();

    const filters = $state(
        currentFilterModel || [
            {
                fiterType: column.type || "text",
                operator: "equal",
                value: [],
            },
        ]
    );

    // FilterModel
</script>

<div class="flex flex-col gap-2 text-xs p-1">
    {#each filters as filter}
        <div class="flex flex-col gap-2 border ">
            <select
                value={filter.operator}
                onchange={(ev) => {
                    filter.operator = ev.target.value;
                }}
                class="p-2"
            >
                {#each OPERATORS as op}
                    <option>{op}</option>
                {/each}
            </select>

            {#if filter.operator === "range"}
                <input
                    class="p-2 border"
                    type="text"
                    value={filter.value.at(0) || ""}
                    onchange={(ev) => {
                        filter.value[0] = ev.target.value;
                    }}
                />

                <input
                    class="p-2 border"
                    type="text"
                    value={filter.value.at(1) || ""}
                    onchange={(ev) => {
                        filter.value[1] = ev.target.value;
                    }}
                />
            {:else}
                <input
                    class="p-2 border"
                    type="text"
                    value={filter.value.at(0) || ""}
                    onchange={(ev) => {
                        filter.value[0] = ev.target.value;
                    }}
                />
            {/if}
        </div>
    {/each}

    <button 
        class="btn bg-gray-50 hover:bg-gray-100" 
        onclick={() => {
            filters.push({
                fiterType: column.type || "text",
                operator: "equal",
                value: [],
            });
        }}
        >add</button
    >

    <div class="flex justify-end gap-1">
        <button class="btn btn-sm variant-filled" onclick={closeFilter}
            >close</button
        >

        <button class="btn btn-sm variant-filled-tertiary" onclick={clearFilter}
            >clear</button
        >

        <button
            class="btn btn-sm variant-filled-primary"
            onclick={() => {
                applyFilter(column, filters);
            }}>apply</button
        >
    </div>
</div>
