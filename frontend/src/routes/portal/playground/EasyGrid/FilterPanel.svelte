<script lang="ts">
    import { OPERATORS } from "./easyTypes";

    let {
        closeFilter,
        applyFilter,
        clearFilter,
        operator,
        column,
        value,
    } = $props();

    // let stateOperator = $state( operator ||"equal");
    // let stateValue = $state(value || "");
</script>

<div class="flex flex-col gap-2 text-xs">
    <select bind:value={operator} class="p-2">
        {#each OPERATORS as op}
            <option>{op}</option>
        {/each}
    </select>

    <input bind:value class="p-2 border" type="text" />

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
                applyFilter({
                    operator,
                    value,
                    column,
                });
            }}>apply</button
        >
    </div>
</div>
