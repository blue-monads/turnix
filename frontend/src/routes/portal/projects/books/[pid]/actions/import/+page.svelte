<script lang="ts">
    import { goto } from "$app/navigation";
    import { AutoTable, FloatyButton, Loader } from "$lib/compo";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";

    const pid = $page.params["pid"];

    const store = getModalStore();
    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);

    let loading = $state(false);

    let file: FileList | null = $state(null);
    let importAsNewTxn = $state(false);
</script>

{#if loading}
    <Loader />
{:else}
    <div class="p-2">
        <div class="card p-2">
            <header class="flex justify-center">
                <div class="flex flex-col gap-2">
                    <h3 class="h3">Import Project</h3>
                </div>
            </header>

            <section class="flex flex-col gap-4">
                <label class="label">
                    <span>File</span>
                    <input type="file" bind:files={file} required />
                </label>

                <label class="label">
                    <span>As New Transaction</span>
                    <input type="checkbox" bind:checked={importAsNewTxn} />
                </label>
            </section>

            <header class="flex justify-end p-2">
                <button
                    disabled={!file}
                    class="btn btn-sm variant-filled"
                    onclick={async () => {
                        if (!file) {
                            return;
                        }

                        loading = true;

                        const resp = await api.importData(pid, {
                            data: JSON.parse(await file[0].text()),
                            as_new_txn: importAsNewTxn,
                        });

                        if (resp.status !== 200) {
                            return;
                        }

                        loading = false;
                        goto(`/z/pages/portal/projects/books/${pid}`);
                    }}>Submit</button
                >
            </header>
        </div>
    </div>
{/if}
