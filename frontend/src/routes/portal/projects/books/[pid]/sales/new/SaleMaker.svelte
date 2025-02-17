<script lang="ts">
    import { run } from "svelte/legacy";

    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import type { BooksAPI, Contact } from "$lib/projects/books";
    import type { NewSaleLine, SaleData } from "./sub/sales";
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
        name?: string;
        title?: string;
        notes?: string;
        client_id?: number;
        client_name?: string;
        total_item_price?: number;
        total_item_tax_amount?: number;
        total_item_discount_amount?: number;
        sub_total?: number;
        overall_discount_amount?: number;
        overall_tax_amount?: number;
        total?: number;
        sales_date?: any;
        lines?: NewSaleLine[];
        contactsNameIndex?: Record<number, string>;
        submit: (data: SaleData) => Promise<string>;
        modeOptions?: [string, string][];
        api: BooksAPI;
    }

    let {
        pid,
        name = "New Sale",
        title = $bindable(""),
        notes = $bindable(""),
        client_id = $bindable(0),
        client_name = "",
        total_item_price = $bindable(0),
        total_item_tax_amount = $bindable(0),
        total_item_discount_amount = $bindable(0),
        sub_total = $bindable(0),
        overall_discount_amount = $bindable(0),
        overall_tax_amount = $bindable(0),
        total = $bindable(0),
        sales_date = $bindable(new Date().toISOString().slice(0, 16)),
        modeOptions = [
            ["invoice", "Invoice"],
            ["direct_sale", "Direct Sale"],
        ],
        lines = $bindable([
            // {
            //     info: "test",
            //     product_id: 1,
            //     qty: 4,
            //     amount: 20,
            //     price: 24,
            //     tax_amount: 1,
            //     discount_amount: 4,
            //     total_amount: 84,
            // },
        ]),
        contactsNameIndex = {},
        submit,
        api,
    }: Props = $props();

    let mode = $state(modeOptions.length > 0 ? modeOptions[0][0] : "");
    let message = $state("");

    let overall_tax_percentage = $state(0);
    let overall_discount_percentage = $state(0);

    run(() => {
        total_item_price = lines.reduce((acc, item) => {
            return acc + item.amount * item.qty;
        }, 0);

        total_item_tax_amount = lines.reduce((acc, item) => {
            return acc + item.tax_amount * item.qty;
        }, 0);

        total_item_discount_amount = lines.reduce((acc, item) => {
            return acc + item.discount_amount * item.qty;
        }, 0);

        sub_total = lines.reduce((acc, item) => {
            return acc + item.total_amount;
        }, 0);

        overall_tax_percentage = (overall_tax_amount / sub_total) * 100;
        overall_discount_percentage =
            (overall_discount_amount / sub_total) * 100;

        total = sub_total + overall_tax_amount - overall_discount_amount;
    });

    const clientPicker = () => {
        store.trigger({
            type: "component",
            component: "books_client_picker",
            meta: {
                api,
                pid,
                onSelect: (data: any) => {
                    if (!title) {
                        title = data["name"] || "";
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

    const onSetOverAllTax = () => {
        store.trigger({
            type: "component",
            component: "books_sales_overall_tax_picker",
            meta: {
                sub_total,
                overall_tax_amount,
                onSet: (value: number) => {
                    overall_tax_amount = value;
                },
            },
        });
    };

    const onSetOverAllDiscount = () => {
        store.trigger({
            type: "component",
            component: "books_sales_overall_discount_picker",
            meta: {
                sub_total,
                overall_discount_amount,
                onSet: (value: number) => {
                    overall_discount_amount = value;
                },
            },
        });
    };
</script>

<form class="p-2">
    <div class="card">
        <header class="card-header flex justify-between">
            <h3 class="h3">{name}</h3>

            <div class="w-64">
                {#if modeOptions.length > 0}
                    <RadioGroup>
                        {#each modeOptions as [value, name]}
                            <RadioItem bind:group={mode} name="justify" {value}
                                >{name}</RadioItem
                            >
                        {/each}
                    </RadioGroup>
                {/if}
            </div>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <div class="flex gap-2 justify-between">
                <div class="flex flex-col gap-2">
                    <label class="label">
                        <span>Title</span>

                        <div class="flex gap-2">
                            <input
                                type="text"
                                bind:value={title}
                                class="input p-1"
                                placeholder="title"
                            />
                        </div>
                    </label>

                    <label class="label">
                        <span>Billed To</span>

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
                            <th class="px-4 py-2 text-right">Tax</th>
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
                                >
                                    {#if line.amount === line.price}
                                        {line.amount}
                                    {:else}
                                        <span
                                            >{line.price} - ({line.discount_amount})</span
                                        >
                                        =
                                        <strong>{line.amount}</strong>
                                    {/if}
                                </td>
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                >
                                    {#if line.tax_amount}
                                        {line.tax_amount}
                                    {:else}
                                        -
                                    {/if}
                                </td>
                                <td
                                    class="px-4 py-2 text-right border tabular-nums slashed-zero"
                                >
                                    (
                                    <span>
                                        {line.price}
                                        {#if line.discount_amount}
                                            - {line.discount_amount}
                                        {/if}

                                        {#if line.tax_amount}
                                            + {line.tax_amount}
                                        {/if}
                                    </span>

                                    ) * {line.qty} =

                                    <strong>
                                        {line.total_amount}
                                    </strong>
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
                                <td class="px-2 py-2 border border-gray-400"
                                    >Total Items TAX</td
                                >
                                <td
                                    class="px-2 py-2 border border-gray-400 text-right"
                                >
                                    <span>
                                        {formatCurrency(total_item_tax_amount)}
                                    </span>
                                </td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border border-gray-400"
                                    >Total Items Discount</td
                                >
                                <td
                                    class="px-2 py-2 border border-gray-400 text-right"
                                >
                                    <span>
                                        {formatCurrency(
                                            total_item_discount_amount,
                                        )}
                                    </span>
                                </td>
                            </tr>

                            <tr>
                                <td
                                    class="px-2 py-2 border border-gray-400 border-b-gray-800"
                                >
                                    Sub Total
                                </td>
                                <td
                                    class="px-2 py-2 border border-gray-400 border-b-gray-800 text-right"
                                >
                                    <strong>{formatCurrency(sub_total)}</strong>
                                </td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border border-gray-800"
                                    >Overall TAX</td
                                >
                                <td
                                    class="px-2 py-2 border border-gray-800 text-right"
                                >
                                    <button
                                        class="underline"
                                        onclick={onSetOverAllTax}
                                    >
                                        <span>
                                            <strong>
                                                {formatCurrency(
                                                    overall_tax_amount,
                                                )}
                                            </strong>
                                            {#if overall_tax_percentage}
                                                [{formatCurrency(
                                                    overall_tax_percentage,
                                                )}%]
                                            {/if}
                                        </span>
                                    </button>
                                </td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border border-gray-800">
                                    Overall Discount
                                </td>

                                <td
                                    class="px-2 py-2 border border-gray-800 text-right"
                                >
                                    <button
                                        class="underline"
                                        onclick={onSetOverAllDiscount}
                                    >
                                        <span>
                                            <strong>
                                                {formatCurrency(
                                                    overall_discount_amount,
                                                )}
                                            </strong>
                                            {#if overall_discount_percentage}
                                                [{formatCurrency(
                                                    overall_discount_percentage,
                                                )}%]
                                            {/if}
                                        </span>
                                    </button>
                                </td>
                            </tr>

                            <tr>
                                <td
                                    class="px-2 py-2 border border-gray-800 font-semibold"
                                >
                                    Total
                                </td>
                                <td
                                    class="px-2 py-2 border border-gray-800 text-right font-semibold"
                                >
                                    {formatCurrency(total)}
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div>
                <p class="text-red-500">{message}</p>
            </div>
        </section>
        <footer class="card-footer flex justify-end">
            <button
                disabled={lines.length === 0 || !client_id}
                class="btn variant-filled"
                onclick={async () => {
                    const resp = await submit({
                        sale: {
                            title,
                            notes,
                            client_id,
                            client_name,
                            total_item_price,
                            total_item_tax_amount,
                            total_item_discount_amount,
                            sub_total,
                            overall_discount_amount,
                            overall_tax_amount,
                            total,
                            sales_date,
                        },
                        lines: lines,
                    });

                    message = resp;
                }}
            >
                save
            </button>
        </footer>
    </div>
</form>
