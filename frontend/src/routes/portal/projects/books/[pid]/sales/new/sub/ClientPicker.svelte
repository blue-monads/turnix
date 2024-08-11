<script lang="ts">
    import { AutoTable, Loader } from "$lib/compo";
    import type { BooksAPI } from "$lib/projects/books";
    import { getModalStore } from "@skeletonlabs/skeleton";
    export let parent: any;

    const store = getModalStore();
    let mode: "loading" | "main";

    let data: object[] = $store[0].meta["data"] || [];

    const loadData = async () => {
        if (data.length) {
            return;
        }

        const api = $store[0].meta["api"] as BooksAPI;
        if (!api) {
            return;
        }

        mode = "loading";

        const pid = $store[0].meta["pid"] as string;
        const resp = await api.listContacts(pid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        mode = "main";
    };

    loadData();
</script>

<div class="card rounded p-4 w-modal-wide min-h-96 {parent.backgroundColor}">
    <h4 class="h4">Pick Contact</h4>

    {#if mode === "loading"}
        <Loader />
    {:else}
        <AutoTable
            action_key={"id"}
            key_names={[
                ["id", "ID"],
                ["name", "Name"],
                ["ctype", "Type"],
                ["email", "Email"],
                ["phone", "Phone"],
            ]}
            datas={data}
            color={["ctype"]}
            actions={[
                {
                    Name: "select",
                    icon: "plus",
                    // @ts-ignore
                    Action: async (id, data) => {
                        const onSelect = $store[0].meta["onSelect"];
                        onSelect?.(data);
                        store.close();
                    },
                },
            ]}
        />
    {/if}
</div>
