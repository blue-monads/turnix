<script lang="ts">
    import { AutoTable } from "$lib/compo";
    import type { BooksAPI } from "$lib/projects/books";
    import { getModalStore } from "@skeletonlabs/skeleton";
    export let parent: any;

    const store = getModalStore();
    let mode: "loading" | "main";
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

        mode = "loading";

        const pid = $store[0].meta["pid"] as string;
        const resp = await api.listProducts(pid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        mode = "main";
    };

    loadData();

    console.log("@parnet", parent);

    const onSubmit = async () => {};
</script>

<div class="card rounded p-4 w-modal-wide min-h-96 {parent.backgroundColor}">
    <h4 class="h4">Pick Client</h4>

    {#if mode === "loading"}
        <div class="flex justify-center items-center">
            <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>
    {:else}
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
                    Action: async (id, data) => {},
                },
            ]}
        />
    {/if}
</div>
