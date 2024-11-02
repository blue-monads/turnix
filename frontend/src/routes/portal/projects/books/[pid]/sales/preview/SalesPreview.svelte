<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import type {
        BooksAPI,
        Contact,
        Estimate,
        EstimateLine,
        Sale,
        SaleLine,
    } from "$lib/projects/books";
    import {
        getModalStore,
        RadioGroup,
        RadioItem,
    } from "@skeletonlabs/skeleton";
    import { formatCurrency } from "$lib";

    const store = getModalStore();

    interface Props {
        pid: number;
        lines: SaleLine[] | EstimateLine[];
        data: Sale | Estimate;
        name?: string;

        contactsNameIndex?: Record<number, string>;
    }

    let { pid, name="Sale", lines, data, contactsNameIndex = {} }: Props = $props();

    // export let api: BooksAPI;

    let overall_tax_percentage = 0;
    let overall_discount_percentage = 0;
</script>

<form class="p-2">
    <div class="card">
        <header class="card-header flex justify-between">
            <h3 class="h3">{name} #{data.id}</h3>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <div class="flex gap-2 justify-between">
                <div class="flex flex-col gap-2">
                    <div class="label">
                        <span class="font-semibold">Title : </span>
                        <span>{data.title || ""}</span>
                    </div>

                    <div class="label">
                        <span class="font-semibold">Billed To : </span>
                        <a
                            class="text-sm italic underline"
                            href={`/z/pages/portal/projects/books/${pid}/contacts/edit?cid=${data.client_id}`}
                        >
                            {contactsNameIndex[data.client_id] || ""}
                        </a>
                    </div>
                </div>

                <label class="label">
                    <span>Date of Sale</span>
                    <input
                        type="datetime-local"
                        class="input p-1"
                        placeholder="Date"
                        disabled
                        value={new Date(data.created_at)
                            .toISOString()
                            .slice(0, 16)}
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
                            <th class="px-4 py-2 text-right">Total</th>
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
                                    <span
                                        >{line.price} - ({line.discount_amount})</span
                                    >
                                    =
                                    <strong
                                        >{formatCurrency(
                                            (line.price || 0) -
                                                (line.discount_amount || 0) +
                                                (line.tax_amount || 0),
                                        )}</strong
                                    >
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
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>

            <div class="flex gap-2 w-full justify-between">
                <label class="label flex flex-col gap-2 w-full md:w-1/2">
                    <span>Notes</span>
                    <textarea
                        class="textarea p-1"
                        value={data.notes}
                        disabled={true}
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
                                        {formatCurrency(
                                            data.total_item_tax_amount,
                                        )}
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
                                            data.total_item_discount_amount,
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
                                    <strong
                                        >{formatCurrency(
                                            data.sub_total,
                                        )}</strong
                                    >
                                </td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border border-gray-800"
                                    >Overall TAX</td
                                >
                                <td
                                    class="px-2 py-2 border border-gray-800 text-right"
                                >
                                    <span>
                                        <strong>
                                            {formatCurrency(
                                                data.overall_tax_amount,
                                            )}
                                        </strong>
                                        {#if overall_tax_percentage}
                                            [{formatCurrency(
                                                overall_tax_percentage,
                                            )}%]
                                        {/if}
                                    </span>
                                </td>
                            </tr>

                            <tr>
                                <td class="px-2 py-2 border border-gray-800">
                                    Overall Discount
                                </td>

                                <td
                                    class="px-2 py-2 border border-gray-800 text-right"
                                >
                                    <span>
                                        <strong>
                                            {formatCurrency(
                                                data.overall_discount_amount,
                                            )}
                                        </strong>
                                        {#if overall_discount_percentage}
                                            [{formatCurrency(
                                                overall_discount_percentage,
                                            )}%]
                                        {/if}
                                    </span>
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
                                    {formatCurrency(data.total)}
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </section>
        <footer class="card-footer flex justify-end">
            <button class="btn variant-soft-error" onclick={() => {}}>
                delete
            </button>
        </footer>
    </div>
</form>
