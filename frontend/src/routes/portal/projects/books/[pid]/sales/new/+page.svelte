<script lang="ts">
    import { NewBookAPI, type Catagory } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoForm } from "$lib/compo";
    import { goto } from "$app/navigation";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import type { SaleLine } from "./sales";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import SalesItemPick from "./SalesItemPick.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let message = "";

    const store = getModalStore();

    let name = "Sale #1234";
    let info = "";
    let totalPrice = 0;
    let client_id = 1;
    let sales_date = new Date().toISOString().slice(0, 16)

    let lines: SaleLine[] = [
        {
            info: "Shared Hosting - Simple Plan (Monthly)",
            product_id: 1,
            qty: 1,
            amount: 2.45,
        },
    ];

    const clientIndex: Record<number, string> = {
        1: "Philip moros",
        2: "John Doe",
    };

    const submit = async () => {};
</script>

<form class="p-2" on:submit|preventDefault={submit}>
    <div class="card">
        <header class="card-header">
            <h4 class="h4">New Sale</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <div class="flex gap-2 justify-between">
                <label class="label">
                    <span>Billed To</span>
                    <input
                        type="text"
                        value={clientIndex[client_id] || ""}
                        class="input p-1"
                        placeholder="Client"
                    />
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

            <label class="label flex flex-col gap-2">
                <span>Title</span>
                <input
                    bind:value={name}
                    class="input p-1 w-36"
                    type="text"
                    placeholder="Input"
                />
            </label>

            <div class="max-h-96 overflow-auto">
                <table
                    class="w-full text-left table-auto print:text-sm"
                    id="table-items"
                >
                    <thead>
                        <tr
                            class="variant-filled  print:bg-gray-300 print:text-black"
                        >
                            <th class="px-4 py-2">Item</th>
                            <th class="px-4 py-2 text-right">Qty</th>
                            <th class="px-4 py-2 text-right">Unit Price</th>
                            <th class="px-4 py-2 text-right">Subtotal</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each lines as line}
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
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="flex justify-start p-2">
                <button
                    class="btn btn-sm variant-filled"
                    on:click={() => {
                        store.trigger({
                            type: "component",
                            component: "books_sales_item_pick",
                            meta: {
                                api,
                                pid,
                                
                                onSave: (data) => {
                                    lines.push(data);
                                    lines = lines;
                                },
                            },
                        });

                        // lines.push({
                        //     info: "Shared Hosting - Simple Plan (Monthly)",
                        //     product_id: 1,
                        //     qty: 1,
                        //     amount: 2.45,
                        // });
                        // lines = lines;
                    }}
                >
                    <SvgIcon
                        name="plus"
                        className="w-4 h-4 inline-block align-middle"
                    />
                </button>
            </div>

            <div class="flex gap-2 w-full justify-between">
                <label class="label flex flex-col gap-2 w-full md:w-1/3">
                    <span>Notes</span>
                    <textarea
                        bind:value={info}
                        class="textarea p-1 w-full"
                        rows="4"
                        placeholder={"information about account"}
                    />
                </label>

                <label class="label flex flex-col gap-2 w-full md:w-1/3">
                    <span>Total</span>
                    <input
                        type="number"
                        bind:value={totalPrice}
                        class="input p-1 w-full"
                        placeholder="Input"
                    />
                </label>
            </div>
        </section>
        <footer class="card-footer flex justify-end">
            <button type="submit" class="btn variant-filled"> save </button>
        </footer>
    </div>
</form>
