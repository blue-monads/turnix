<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { page } from "$app/stores";
    import { NewSimpleRATApi } from "../../lib/SimpleRATApi";

    let name = $state("");
    let registerURL = $state("");

    interface Props {
        pid: string;
        onAddDevice: () => void;
    }
    

    let { pid, onAddDevice }: Props = $props();

    let mode: "form" | "result" | "loading" = $state("form");

    const api = NewSimpleRATApi(getContext("__api__") as RootAPI);

    const onChange = async (data: Record<string, any>) => {
        mode = "loading";

        const resp = await api.addDevice(pid, data);
        if (resp.status !== 200) {
            console.error("Failed to add device", resp);
            return;
        }

        // "http://localhost:7703/z/project/simplerat/1/device/finish-setup?token=xyz"

        const token = (resp.data)

        registerURL = `${window.location.origin}/z/project/simplerat/${pid}/device/finish-setup?token=${token}`;

        mode = "result";
    };
</script>

<div class="p-2">
    <div class="card">
        {#if mode === "form"}
            <header class="card-header">
                <h4 class="h4">Add Device</h4>
            </header>

            <section class="p-4 flex flex-col gap-4">
                <label class="label">
                    <span>Name</span>
                    <input
                        bind:value={name}
                        class="input p-1"
                        type="text"
                        placeholder="Name"
                    />
                </label>
            </section>
            <footer class="card-footer flex justify-end">
                <button
                    onclick={() => {
                        onChange({ name });
                    }}
                    class="btn variant-filled"
                >
                    create
                </button>
            </footer>

        {:else if mode === "loading"}
            <div class="p-4 flex justify-center">
                <div class="spinner">
                    loading...
                </div>
            </div>
        {:else if mode === "result"}

            <header class="card-header">
                <h4 class="h4">Device Register token</h4>
            </header>

            <section class="p-4 flex flex-col gap-4">
                <label class="label">
                    <span>register URL</span>
                    <textarea
                        bind:value={registerURL}
                        class="input p-1 textarea"
                        cols="30"
                        disabled
                    >
                    </textarea>
                    <button
                        class="rounded p-0.5 bg-gray-200"
                        onclick={() => {
                            navigator.clipboard.writeText(registerURL);
                        }}
                    >
                        Copy
                    </button>
                </label>
            </section>
            <footer class="card-footer flex justify-end">
                <button
                    onclick={() => {
                        onAddDevice();
                    }}
                    class="btn variant-filled"
                >
                    close
                </button>
            </footer>
        {/if}
    </div>
</div>
