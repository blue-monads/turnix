<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import { page } from "$app/stores";
    import { NewSimpleRATApi } from "../../lib/SimpleRATApi";


    let name = "";

    interface Props {
        pid: string;
        onAddDevice: () => void;
    }

    let { pid, onAddDevice }: Props = $props();


    const api = NewSimpleRATApi(getContext("__api__") as RootAPI);

    let loading = $state(false);
    const onChange = async (data: Record<string, any>) => {
        loading = true;

        await api.addDevice(pid, data);

        onAddDevice();


        loading = false;
    };
</script>

<div class="p-2">
    <div class="card">
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
                class="btn variant-filled"> create 
        </button>
        </footer>
    </div>
</div>
