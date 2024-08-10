<script lang="ts">
    import { goto } from "$app/navigation";
    import { AutoTable, FloatyButton } from "$lib/compo";
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

    let loading = false;

    const saveAsJsonFile = (data: any, filename: string) => {
        const blob = new Blob([JSON.stringify(data, null, 2)], {
            type: "application/json",
        });
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.href = url;
        link.download = filename;
        link.click();
        URL.revokeObjectURL(url);
    };
</script>

<div class="p-2">
    <div class="card p-2 w-80 mx-auto my-10">
        <header class="flex justify-center">
            <div class="flex flex-col gap-2">
                <h3 class="h3">Export Project</h3>

                {#if loading}
                    <div class="p-2">Processing..</div>
                {/if}
            </div>
        </header>

        <footer class="flex justify-end p-2">
            <button
                disabled={loading}
                class="btn btn-sm variant-filled"
                on:click={async () => {
                    loading = true;

                    const resp = await api.exportData(pid);
                    if (resp.status !== 200) {
                        return;
                    }

                    console.log(resp.data);

                    saveAsJsonFile(
                        resp.data,
                        `books_export_${new Date().toISOString()}.json`,
                    );

                    loading = false;
                }}>Start</button
            >
        </footer>
    </div>
</div>
