<script lang="ts">
    import { formatCurrency } from "$lib";
    import { getModalStore } from "@skeletonlabs/skeleton";
    export let parent: any;

    const store = getModalStore();

    const meta = $store[0].meta || {};
    let sub_total = meta["sub_total"] || 0;
    let overall_discount_amount = meta["overall_discount_amount"] || 0;
    let onSet: (data: any) => void = meta["onSet"] || (() => {});

    let new_discouned_amount = sub_total - overall_discount_amount;

    let overall_discount_percentage = 0;

    $: {
        overall_discount_percentage =
            ((sub_total - new_discouned_amount) / sub_total) * 100;
    }
</script>

<div class="card rounded w-modal-slim {parent.backgroundColor}">
    <section class="p-4 flex flex-col gap-4">
        <label class="label">
            <span>Overall Discounted Amount</span>
            <input
                bind:value={new_discouned_amount}
                class="input p-1"
                type="number"
                placeholder="Input"
            />
        </label>
    </section>

    <footer class="card-footer flex justify-between">
        <div class="text-xs">
            <span>Overall Discount Percentage</span>
            <span>{formatCurrency(overall_discount_percentage)}%</span>
        </div>

        <button
            class="btn variant-filled"
            on:click={() => {
                onSet(sub_total - new_discouned_amount);
                store.close();
            }}
        >
            set
        </button>
    </footer>
</div>
