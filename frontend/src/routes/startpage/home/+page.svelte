<script lang="ts">
    import { type StartAPI } from "../startAPI";
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
    let remoteAddr = $state("");
    let meshAddr = $state("");

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
        meshAddr = resp.data["mesh_addr"];
    };

    load();

    const localNavigate = () => {
        // local-navigate
        const rpcFunc = (window as any)["__ebrowser_rpc__"];

        rpcFunc("local-navigate", {
            url: `http://localhost${port}/z/pages`,
        });
    };

    const remoteNavigate = () => {
        const rpcFunc = (window as any)["__ebrowser_rpc__"];

        rpcFunc("remote-navigate", {
            url: remoteAddr,
        });
    };
</script>

<div class="flex items-center justify-center min-h-screen flex-col gap-4">
    <div class="card flex flex-col p-4 max-w-md gap-4 justify-center">
        {#if loading}
            <div class="flex items-center justify-center h-96 flex-col gap-4">
                <div class="flex items-center justify-center flex-col gap-4">
                    <h2 class="h2">Loading...</h2>
                </div>
            </div>
        {:else if isInstanceRunning}
            <h2 class="h2">Turnix is running.</h2>
            <span class="text-indigo-500 chip">{workingDir} </span>
            <p class="p text-sm">
                click explore to start using it. Or click on
                {#if !isHomePage}
                    <span>options</span>
                {:else}
                    <span>sidebar</span>
                {/if}
            </p>

            <div class="flex gap-2">
                <button
                    class="btn variant-soft-primary"
                    onclick={localNavigate}
                >
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

    <div class="card p-4 max-w-md w-full">
        <label>
            <span>Remote Address</span>
            <div class="flex flex-col gap-2">
                <input
                    type="text"
                    class="input p-1"
                    placeholder="http://xxyyzz-7773.lpweb"
                    bind:value={remoteAddr}
                />
                <button
                    class="btn variant-soft-success"
                    onclick={remoteNavigate}
                >
                    Explore
                </button>
            </div>
        </label>
    </div>


    <div class="card p-4 max-w-md w-full">
        <label>
            <span>Current Node address</span>
            <div class="flex flex-col gap-2">
                <input
                    type="text"
                    class="input p-1"
                    value={meshAddr}
                />
            </div>
        </label>
    </div>


</div>
