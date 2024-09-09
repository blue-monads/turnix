<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import type { Contact } from "$lib/projects/books";
    import { createEventDispatcher } from "svelte";
    import type { NewSaleLine } from "./sub/sales";
    import {
        getModalStore,
        RadioGroup,
        RadioItem,
    } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";

    const store = getModalStore();
    const dispatch = createEventDispatcher();

    export let pid = 0;

    // data fields
    export let title = "";
    export let notes = "";
    export let client_id = 0;
    export let total_item_price = 0;
    export let total_item_tax_amount = 0;
    export let total_item_discount_amount = 0;
    export let sub_total = 0;
    export let overall_discount_amount = 0;
    export let overall_tax_amount = 0;
    export let total = 0;
    export let sales_date = new Date().toISOString().slice(0, 16);

    // extra data
    export let lines: NewSaleLine[] = [];
    export let contacts: Contact[] = [];
    export let contactsNameIndex: Record<number, string> = {};
    export let submit = async () => {};
    export let api: RootAPI;

    let mode = "invoice";

    const clientPicker = () => {
        store.trigger({
            type: "component",
            component: "books_client_picker",
            meta: {
                api,
                pid,
                onSelect: (data: any) => {
                    client_id = data["id"] || 0;
                },
            },
        });
    };

    const salesPicker = () => {
        store.trigger({
            type: "component",
            component: "books_sales_item_pick",
            meta: {
                api,
                pid,

                onSave: (data: any) => {
                    lines.push(data);
                    lines = lines;
                },
            },
        });
    };
</script>

<form class="p-2" on:submit|preventDefault={submit}>
    <div class="card">
        <header class="card-header flex justify-between">
            <h3 class="h3">New Sale</h3>
            <div class="w-64">
                <RadioGroup>
                    <RadioItem
                        bind:group={mode}
                        name="justify"
                        value={"direct_sale"}>Direct Sale</RadioItem
                    >
                    <RadioItem
                        bind:group={mode}
                        name="justify"
                        value={"invoice"}>Invoice</RadioItem
                    >
                </RadioGroup>
            </div>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <div class="flex gap-2 justify-between">
                <label class="label">
                    <span>Billed To</span>

                    <div class="flex gap-2">
                        <input
                            type="text"
                            disabled
                            value={contactsNameIndex[client_id] || ""}
                            class="input p-1"
                            placeholder="Client"
                        />

                        <button on:click={clientPicker}>
                            <SvgIcon
                                name="plus"
                                className="w-4 h-4 inline-block align-middle"
                            />
                        </button>
                    </div>
                </label>

                <label class="label">
                    <span>Date of Sale</span>
                    <input
                        type="datetime-local"
                        class="input p-1"
                        placeholder="Date"
                        bind:value={sales_date}
                    />
                </label>
            </div>

            <div class="max-h-96 overflow-auto">
                <table
                    class="w-full text-left table-auto print:text-sm"
                    id="table-items"
                >
                    <thead>
                        <tr
                            class="variant-filled print:bg-gray-300 print:text-black"
                        >
                            <th class="px-4 py-2">Item</th>
                            <th class="px-4 py-2 text-right">Qty</th>
                            <th class="px-4 py-2 text-right">Unit Price</th>
                            <th class="px-4 py-2 text-right">Subtotal</th>
                            <th class="px-4 py-2 text-right"></th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each lines as line, index}
                            <tr>
                                <td class="px-4 py-2 border">{line.info}</td>
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                    >{line.qty}</td
                                >
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                    >{line.amount}</td
                                >
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                    >{line.amount * line.qty}</td
                                >
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                >
                                    <button
                                        class="hover:underline text-warning-800 text-xs"
                                        on:click={() => {
                                            lines.splice(index, 1);
                                            lines = lines;
                                        }}
                                    >
                                        remove
                                    </button>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="flex justify-start p-2">
                <button
                    class="btn btn-sm variant-filled"
                    on:click={salesPicker}
                >
                    <SvgIcon
                        name="plus"
                        className="w-4 h-4 inline-block align-middle"
                    />
                </button>
            </div>

            <div class="flex gap-2 w-full justify-between">
                <label class="label flex flex-col gap-2 w-full md:w-1/2">
                    <span>Notes</span>
                    <textarea
                        bind:value={notes}
                        class="textarea p-1"
                        rows="4"
                        placeholder="Notes about the sale.."
                    />
                </label>

                <div class="flex flex-col gap-2 text-sm">
                    <table>
                        <tbody>
                            <tr>
                                <td class="px-2 py-2 border"> Sub Total </td>
                                <td class="px-2 py-2 border text-right">
                                    10000
                                </td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border">TAX</td>
                                <td class="px-2 py-2 border text-right">12%</td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border">Discount</td>
                                <td class="px-2 py-2 border text-right">0%</td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border font-semibold">
                                    Total
                                </td>
                                <td
                                    class="px-2 py-2 border text-right font-semibold"
                                >
                                    1120
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </section>
        <footer class="card-footer flex justify-end">
            <button class="btn variant-filled" on:click={submit}> save </button>
        </footer>
    </div>
</form>
