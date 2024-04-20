<script lang="ts">
    import { gotoPorjects } from "$lib/nav";
    import { AutoForm, PageLayout, AutoTable } from "$lib/compo";

    import type { API } from "$lib/api";
    import { getContext } from "svelte";
    import Loader from "$lib/compo/loader/loader.svelte";

    import { getModalStore } from "@skeletonlabs/skeleton";
    import Icon from "@krowten/svelte-heroicons/Icon.svelte";

    const modalStore = getModalStore();

    let datas = [];
    let loading = true;

    const load = async () => {
        let api: API = $modalStore[0].meta["api"];

        const resp = await api.listProjectTypes();
        if (resp.status !== 200) {
            return;
        }

        datas = resp.data;
        loading = false;
    };

    load();

    let message = "";
</script>

{#if loading}
    <Loader />
{:else}
    <div class="card p-2 w-modal">
        <header class="header flex justify-center">
            <h4 class="h4">
                <span
                    class="bg-gradient-to-br from-pink-500 to-violet-500 bg-clip-text text-transparent box-decoration-clone uppercase"
                    >Select New Project type</span
                >
            </h4>
        </header>

        <div class="logo-cloud grid-cols-1 lg:!grid-cols-3 gap-1 border p-2">
            {#each datas as data}
                {@const icon = data["icon"]}
                <a
                    class="logo-item"
                    on:click={() => {
                        modalStore.close()
                    }}
                    href="/z/pages/portal/projects/new?ptype={data['ptype']}"
                >
                    {#if icon}
                        <Icon name={icon} class="w-6 h-6" />
                    {/if}
                    <span>{data["name"]}</span>
                </a>
            {/each}
        </div>
    </div>
{/if}
