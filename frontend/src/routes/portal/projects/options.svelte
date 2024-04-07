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
                    >Options</span
                >
            </h4>
        </header>

        <div class="logo-cloud flex flex-col gap-1 border p-2">
            <a
                class="logo-item"
                on:click={() => {
                    modalStore.close();
                }}
                href="/z/pages/portal/ptypes/{$modalStore[0].meta["data"]["ptype"]}?pid={$modalStore[0].meta["data"]["id"]}"
            >
                <Icon name="document-text" class="w-6 h-6" />
                <span>Project Home</span>
            </a>

            <a
                class="logo-item"
                on:click={() => {
                    modalStore.close();
                }}
                href="/z/pages/portal/filestore"
            >
                <Icon name="folder" class="w-6 h-6" />
                <span>Project Files</span>
            </a>

            <a
            class="logo-item"
            on:click={() => {
                modalStore.close();
            }}
            href="/z/pages/portal/projects"
        >
            <Icon name="pencil-square" class="w-6 h-6" />
            <span>Edit Project</span>
        </a>
        </div>
    </div>
{/if}
