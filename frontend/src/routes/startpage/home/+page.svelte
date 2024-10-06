<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    import { buildApi, type StartAPI } from "../startAPI";
    import { getContext } from "svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    interface Props {
        isHomePage?: boolean;
    }

    let { isHomePage = true }: Props = $props();

    const api = getContext("__start_api__") as StartAPI;

    let isInstanceRunning = $state(false);
    let workingDir = $state("");
    let loading = $state(true);
    let port = "";

    const load = async () => {
        loading = true;

        const resp = await api.getStatus();
        if (resp.status !== 200) {
            return;
        }

        isInstanceRunning = resp.data["is_running"];
        workingDir = resp.data["working_dir"];
        port = resp.data["port"] as string;
        loading = false;
    };

    load();

    const localNavigate = () => {
        // local-navigate
        const rpcFunc = (window as any)["__ebrowser_rpc__"];

        rpcFunc("local-navigate", {
            url: `http://localhost${port}/z/pages`,
        });
    };
</script>

<div class="flex items-center justify-center min-h-screen flex-col gap-4 pt-20">
    {#if loading}
        <div
            class="flex items-center justify-center h-96 flex-col gap-4"
        >
            <div
                class="flex items-center justify-center flex-col gap-4"
            >
                <h2 class="h2">Loading...</h2>
            </div>
        </div>
    {:else if isInstanceRunning}
        <h2 class="h2">Turnix is running.</h2>
        <span class="text-indigo-500 chip">{workingDir} </span>
        <p class="p text-xl">
            click explore to start using it. Or click on
            {#if !isHomePage}
                <span>options</span>
            {:else}
                <span>sidebar</span>
            {/if}
        </p>

        <div class="flex gap-2">
            <button class="btn variant-soft-primary" onclick={localNavigate}>
                Explore
            </button>
            {#if !isHomePage}
                <a
                    href="/z/pages/startpage/home"
                    class="btn variant-soft-secondary">Show Options</a
                >
            {/if}
        </div>
    {:else}
        <h2 class="h2">Turnix is not running</h2>
        <p class="p text-xl">
            Click start to run a instance in current directory.
        </p>
        <p class="p text-xl">
            Or click on
            {#if !isHomePage}
                <span>options</span>
            {:else}
                <span>sidebar</span>
            {/if}
        </p>

        <div class="flex gap-2">
            <button
                class="btn variant-filled"
                onclick={async () => {
                    loading = true;

                    await api.startInstance();

                    load();
                }}
            >
            <SvgIcon className="h-4 w-4" name="bolt" />
                Start
            </button>
            {#if !isHomePage}
                <a
                    href="/z/pages/startpage/home"
                    class="btn variant-soft-secondary">Show Options</a
                >
            {/if}
        </div>
    {/if}
</div>
