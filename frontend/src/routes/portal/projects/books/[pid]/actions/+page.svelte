<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { Loader } from "$lib/compo";

    const pid = $page.params["pid"];

    const rootApi = getContext("__api__") as RootAPI;

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
    <Loader />
{:else}
    <div class="p-2">
        <div class="card p-4">
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

                <label class="label">
                    <span>Default unpaid sales account</span>

                    <div class="flex gap-1">
                        <input class="input p-1" disabled={true} />
                        <button class="btn variant-filled-secondary btn-sm">
                            <SvgIcon
                                className="h-4 w-4"
                                name="chevron-up"
                            />
                        </button>
                    </div>
                </label>

                <label class="label">
                    <span>Default  paid sales account</span>

                    <div class="flex gap-1">
                        <input class="input p-1" disabled={true} />
                        <button class="btn variant-filled-secondary btn-sm">
                            <SvgIcon
                                className="h-4 w-4"
                                name="chevron-up"
                            />
                        </button>
                    </div>
                </label>

                <label class="label">
                    <span>Default bills account</span>

                    <div class="flex gap-1">
                        <input class="input p-1" disabled={true} />
                        <button class="btn variant-filled-secondary btn-sm">
                            <SvgIcon
                                className="h-4 w-4"
                                name="chevron-up"
                            />
                        </button>
                    </div>
                </label>


            </section>

            <footer class="flex justify-start pt-8 gap-2">
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
