<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { getDrawerStore } from "@skeletonlabs/skeleton";

    const drawer = getDrawerStore();

    const getRootAPI = $drawer["meta"].getRootAPI as () => RootAPI;

    let messages: Record<string, any>[] = [];

    const load = async () => {
        const api = getRootAPI();

        const resp = await api.listUserMessages();
        if (resp.status !== 200) {
            return;
        }

        messages = resp.data;
    };

    load();
</script>

<div class="w-full h-full relative">
    <button
        on:click={() => {
            drawer.close();
        }}
        class="absolute top-2 right-6 h-4 w-4 rounded border"
    >
        <SvgIcon name="x-mark" className="h-4 w-4" />
    </button>

    <h5 class="h5 p-2">Notifications</h5>

    <div class="flex flex-col gap-2 border-t border-gray-200 p-2 max-h-[90%] overflow-scroll">
        {#each messages as message}
            <div class="card p-2 shadow">
                <div class="flex justify-between">
                    <div class="flex flex-col gap-1">
                        <h6 class="h6">{message.title || ""}</h6>
                        <p class="text-sm text-gray-500">
                            {message.contents || ""}
                        </p>
                    </div>
                </div>
            </div>
        {/each}
    </div>

    <div class="flex justify-end p-2 gap-2">
        <a 
        
        on:click={() => {
            drawer.close();
        }}
        href="/z/pages/portal/self/inbox" 
            class="underline inline-flex items-center gap-1">
            <SvgIcon className="w-4 h-4" name="inbox" />
            inbox
        </a>
    </div>

</div>
