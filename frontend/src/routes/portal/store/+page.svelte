<script lang="ts">
    import type { RootAPI } from "$lib";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import Loader from "$lib/compo/loader/loader.svelte";
    import axios from "axios";
    import { getContext } from "svelte";

    const storeLink =
        "https://raw.githubusercontent.com/blue-monads/store/refs/heads/master/store.json";

    let storeData: Record<string, any> | undefined = $state(undefined);
    let installingPtype: string | undefined = $state(undefined);
    const api = getContext("__api__") as RootAPI;

    const load = async () => {
        const resp = await axios.get(storeLink);
        if (resp.status !== 200) {
            return;
        }

        storeData = resp.data;
    };

    const install = async (data: Record<string, any>) => {
        installingPtype = data.slug;
        const resp = await api.projectTypeInstall(data.url);
        if (resp.status !== 200) {
            return;
        }


        installingPtype = undefined
    };

    load();
</script>

<div class="flex flex-col gap-2 md:p-5">
    {#if storeData}
        <div class="flex items-center max-w-md mx-auto">
            <h2 class="h2">{storeData.name}</h2>
        </div>
        <div class="flex min-h-18 p-4">
            <div class="flex items-center max-w-md mx-auto bg-white rounded-lg">
                <div class="w-full">
                    <input
                        type="search"
                        class="w-full p-2 text-gray-800 rounded-full focus:outline-none"
                        placeholder="search"
                    />
                </div>
                <div>
                    <button
                        class="flex items-center justify-center w-12 h-12 text-white"
                    >
                        <SvgIcon
                            name="magnifying-glass"
                            className="w-8 h-8 text-gray-500"
                        />
                    </button>
                </div>
            </div>
        </div>

        <div class="flex flex-col md:flex-row p-2 justify-center gap-2">
            {#each storeData.listings as list}
                <div
                    class="min-h-12 card flex flex-col p-4 gap-2 md:flex-row sm:min-w-64 md:min-w-96 justify-between"
                >
                    <div>
                        <SvgIcon
                            className="w-16 h-16 text-gray-500"
                            name="squares-2x2"
                        />
                    </div>

                    <div class="flex flex-col min-w-32">
                        <h5 class="h5 text-gray-800">{list.name}</h5>
                        <p class="p">{list.info}</p>
                        <div class="inline-flex gap-1">
                            {#each list.tags as tag}
                                <span
                                    class="chip variant-ringed hover:underline"
                                    >{tag}</span
                                >
                            {/each}
                        </div>
                    </div>

                    <div class="flex md:justify-center items-center">
                        <button
                            disabled={!!installingPtype}
                            onclick={() => {
                                install(list)
                            }}
                            class="btn btn-md variant-filled text-xs hover:opacity-45 disabled:opacity-40"
                        >
                            {#if list.slug === installingPtype}
                                <svg aria-hidden="true" class="w-4 h-4 text-gray-200 animate-spin dark:text-gray-600 fill-slate-500" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                                    <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
                                </svg>
                            {:else}
                                Install
                            {/if}
                        </button>
                    </div>
                </div>
            {/each}
        </div>
    {:else}
        <Loader />
    {/if}
</div>
