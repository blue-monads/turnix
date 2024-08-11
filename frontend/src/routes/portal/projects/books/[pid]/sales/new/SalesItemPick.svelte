<script lang="ts">
    import { AutoTable } from "$lib/compo";
    import type { BooksAPI } from "$lib/projects/books";
    import type { SaleLine } from "./sales";

    import { getModalStore } from "@skeletonlabs/skeleton";
    export let parent: any;

    const store = getModalStore();
    let mode: "loading" | "pick_product" | "set_quantity" = "pick_product";
    let info = "";
    let qty = 1;
    let amount = 0;
    let product_id: number;

    let data: object[] = $store[0].meta["data"] || [];

    const loadData = async () => {
        const api = $store[0].meta["api"] as BooksAPI;
        if (!api) {
            return;
        }

        const pid = $store[0].meta["pid"] as string;
        const resp = await api.listProducts(pid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
    };

    loadData();

    console.log("@parnet", parent);

    const onSubmit = async () => {
        const onSave = $store[0].meta["onSave"] as (data: SaleLine) => void;
        if (onSave) {
            onSave({
                product_id,
                qty,
                amount,
                info,
            });
        }
        store.close();
    };
</script>

<div class="card rounded p-4 w-modal-wide min-h-96 {parent.backgroundColor}">
    {#if $store[0]}
        {#if mode === "pick_product"}
            <h4 class="h4">Pick Product</h4>
            <!-- Enable for debugging: -->
            <AutoTable
                action_key={"id"}
                key_names={[
                    ["id", "ID"],
                    ["name", "Name"],
                    ["info", "Info"],
                    ["price", "Price"],
                    ["variant_id", "Variant"],
                ]}
                datas={data}
                color={[]}
                actions={[
                    {
                        Name: "select",

                        icon: "plus",
                        // @ts-ignore
                        Action: async (id, data) => {
                            product_id = id;
                            amount = data["price"] || 0;
                            info = (data["name"] || "") + "  " + (data["variant_id"] || "");
                            mode = "set_quantity";
                        },
                    },
                ]}
            />
        {:else if mode === "loading"}
            <div class="flex justify-center items-center">
                <div class="spinner-border text-primary" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>
            </div>
        {:else}
            <section class="p-4 flex flex-col gap-4">
                <label class="label">
                    <span>Quantity</span>
                    <input
                        bind:value={qty}
                        class="input p-1"
                        type="number"
                        placeholder="Input"
                    />
                </label>

                <label class="label">
                    <span>Per Unit Amount</span>
                    <input
                        bind:value={amount}
                        class="input p-1"
                        type="number"
                        placeholder="Input"
                    />
                </label>

                <label class="label flex flex-col gap-2 w-full">
                    <span>Notes</span>
                    <textarea
                        bind:value={info}
                        class="textarea p-1 w-full"
                        rows="4"
                        placeholder={"information about account"}
                    />
                </label>
            </section>

            <footer class="card-footer flex justify-end">
                <button class="btn variant-filled" on:click={onSubmit}>
                    save
                </button>
            </footer>
        {/if}
    {/if}
</div>
