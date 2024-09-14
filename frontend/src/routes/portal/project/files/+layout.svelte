<script lang="ts">
    import { getContext } from "svelte";
    import { params } from "$lib/params";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { AppBar } from "@skeletonlabs/skeleton";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import { page } from "$app/stores";
    import type { RootAPI } from "$lib/api";

    const pid = $page.params["pid"];

    const store = getModalStore();

    const api = getContext("__api__") as RootAPI;

    $: _paths = ($params["folder"] || "").split("/");
    $: __filename = $params["filename"];

    $: __epoch = 1;

    const complete_new_folder = async (name: string) => {
        await api.addProjectFolder(pid, name, $params["folder"]);
        __epoch = __epoch + 1;
    };
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4">Project Files</h4>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        {#if !$params["file"]}
            <div class="flex flex-row gap-1">
                <a
                    href="/z/pages/portal/project/files/{pid}/upload?folder={$params[
                        'folder'
                    ] || ''}"
                    class="btn btn-sm variant-filled"
                >
                    <SvgIcon name="cloud-arrow-up" className="h-4 w-4" />
                    <span class="hidden md:inline">Upload</span>
                </a>

                <button
                    class="btn btn-sm variant-filled-primary"
                    on:click={() => {
                        store.trigger({
                            type: "component",
                            component: "new_folder_panel",
                            meta: {
                                onNewName: complete_new_folder,
                            },
                        });
                    }}
                >
                    <SvgIcon name="folder-plus" className="h-4 w4" />
                    <span class="hidden md:inline">Folder</span>
                </button>
            </div>
        {/if}
    </svelte:fragment>
</AppBar>

<div class="flex justify-between p-2 pl-4">
    <ol class="breadcrumb">
        <li class="crumb">
            <a class="anchor" href="/z/pages/portal/project/files/{pid}">Home</a
            >
        </li>

        {#each _paths as path, i}
            <li class="crumb-separator" aria-hidden>/</li>

            <li class="crumb">
                <a
                    class="anchor"
                    href="/z/pages/portal/project/files/{pid}?folder={_paths
                        .slice(0, i + 1)
                        .join('/')}">{path || ""}</a
                >
            </li>
        {/each}

        {#key __filename}
            {#if __filename}
                <li class="crumb-separator" aria-hidden>/</li>
                <li class="crumb">
                    {__filename}
                </li>
            {/if}
        {/key}
    </ol>
</div>

<div class="px-2 h-full">
    <div class="card p-2 h-full">
        {#key __epoch}
            <slot />
        {/key}
    </div>
</div>
