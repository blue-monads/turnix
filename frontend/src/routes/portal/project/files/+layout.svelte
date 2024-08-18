<script lang="ts">
    import { getContext } from "svelte";
    import { params } from "$lib/params";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { AppBar } from "@skeletonlabs/skeleton";

    $: _paths = ($params["folder"] || "").split("/");

    $: __epoch = 1;

    const complete_new_folder = async (name: string) => {
        // await capi.newFolder(($params["folder"] || "") + "/" + name);
        // __epoch = __epoch + 1;
        // app.utils.small_modal_close();
    };
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4">Project Files</h4>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        {#if !$params["file"]}
            <div class="flex flex-row gap-1">
                {#if $params["folder"]}
                    <button
                        class="btn btn-sm variant-filled-primary"
                        on:click={() => {}}
                    >
                        <SvgIcon name="cloud-arrow-up" className="h-4 w4" />
                        <span class="hidden md:inline">Upload</span>
                    </button>
                {/if}

                <button
                    class="btn btn-sm variant-filled-primary"
                    on:click={() => {}}
                >
                    <SvgIcon name="folder-plus" className="h-4 w4" />
                    <span class="hidden md:inline">Folder</span>
                </button>
            </div>
        {/if}
    </svelte:fragment>
</AppBar>

<header class="card-header">
    <div class="flex justify-between">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/project/files">Home</a>
            </li>

            {#each _paths as path, i}
                <li class="crumb-separator" aria-hidden>/</li>

                <li class="crumb">
                    <a
                        class="anchor"
                        href="/z/pages/portal/cabinet/listings/folder?source={$params[
                            'source'
                        ]}&folder={_paths.slice(0, i + 1).join('/')}">{path}</a
                    >
                </li>
            {/each}
        </ol>
    </div>
</header>

<div class="h-full p-0 md:p-2">
    <div class="card p-2 h-full">
        <section class="p-2">
            {#key __epoch}
                <slot />
            {/key}
        </section>
    </div>
</div>
