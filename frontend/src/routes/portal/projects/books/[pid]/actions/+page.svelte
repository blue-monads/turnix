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
    let data: any = undefined;

    const load = async () => {
        const resp = await rootApi.getProject(pid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        console.log(data);
    };

    load();
</script>

{#if loading}
    <div class="p-2">Loading..</div>
{:else}
    <div class="p-2">
        <div class="card p-2">
            <header class="flex justify-center">
                <div class="flex flex-col gap-2">
                    <h3 class="h3">Project Actions</h3>
                </div>
            </header>
            <section class="flex flex-col gap-4">
                <label class="label"
                    ><span>Name</span>
                    <input
                        class="input p-1"
                        type="text"
                        disabled={true}
                        value={data?.name || ""}
                    /></label
                >
                <label class="label"
                    ><span>Info</span>
                    <input
                        class="input p-1"
                        disabled={true}
                        value={data?.info || ""}
                    /></label
                >
            </section>

            <footer class="flex justify-start p-2 gap-2">
                <a
                    href="/z/pages/portal/projects/books/{pid}/actions/import"
                    class="btn variant-filled-secondary btn-sm"
                >
                    <SvgIcon className="h-4 w-4" name="arrow-up-tray" />
                    Import
                </a>

                <a
                    href="/z/pages/portal/projects/books/{pid}/actions/export"
                    class="btn variant-filled-primary btn-sm"
                >
                    <SvgIcon className="h-4 w-4" name="arrow-down-tray" />
                    Export
                </a>


            </footer>
        </div>
    </div>
{/if}
