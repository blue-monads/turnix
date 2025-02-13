<script lang="ts">
    import { run } from "svelte/legacy";

    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import type { BooksAPI, Contact } from "$lib/projects/books";
    import type { NewStockIn, NewStockInLine, StockInData } from "./sub/stocks";
    import {
        getModalStore,
        RadioGroup,
        RadioItem,
    } from "@skeletonlabs/skeleton";
    import { formatCurrency } from "$lib";

    const store = getModalStore();

    // data fields

    // extra data

    interface Props {
        pid: number;
        info?: string;
        
        client_id?: number;
        client_name?: string;
     
        sales_date?: any;
        lines?: NewStockInLine[];
        contactsNameIndex?: Record<number, string>;
        submit: (data: StockInData) => Promise<string>;
        api: BooksAPI;
    }

    let {
        pid,
        info = "",
        client_id = $bindable(0),
      
        sales_date = $bindable(new Date().toISOString().slice(0, 16)),
        lines = $bindable([
           
        ]),
        contactsNameIndex = {},
        submit,
        api,
    }: Props = $props();

   

    const clientPicker = () => {
        store.trigger({
            type: "component",
            component: "books_client_picker",
            meta: {
                api,
                pid,
                onSelect: (data: any) => {
                    if (!info) {
                        info = data["name"] || "";
                    }

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

<form class="p-2">
    <div class="card">
        <header class="card-header flex justify-between">
            <h3 class="h3">New stock in</h3>

            
        </header>

        <section class="p-4 flex flex-col gap-4">
            <div class="flex gap-2 justify-between">
                <div class="flex flex-col gap-2">
                    
                    <div></div>

                    <label class="label">
                        <span>From</span>

                        <div class="flex gap-2">
                            <span class="text-sm italic"
                                >{contactsNameIndex[client_id] || ""}
                            </span>

                            <button onclick={clientPicker}>
                                <SvgIcon
                                    name="plus"
                                    className="w-4 h-4 inline-block align-middle"
                                />
                            </button>
                        </div>
                    </label>
                </div>

                <label class="label">
                    <span>Date</span>
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
                                >
                                <strong>{line.amount}</strong>
                                </td>
                               
                                
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                >
                                    <button
                                        class="hover:underline text-warning-800 text-xs"
                                        onclick={() => {
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
                <button class="btn btn-sm variant-filled" onclick={salesPicker}>
                    <SvgIcon
                        name="plus"
                        className="w-4 h-4 inline-block align-middle"
                    />
                </button>
            </div>


            <div class="flex justify-start">
                <label class="label md:w-1/2">
                    <span>Notes</span>

                    <div class="flex gap-2">
                        

                        <textarea
                            bind:value={info}
                            class="textarea p-1"
                            rows="4"
                            placeholder="Notes about the sale.."
                            > </textarea>
                    </div>
                </label>
            </div>

        </section>

    </div>
</form>
