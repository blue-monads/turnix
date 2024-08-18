<script lang="ts">
    import { createEventDispatcher, getContext } from "svelte";
    import type { RootAPI, File } from "$lib/api";
    import { FolderIcon, Loader } from "$lib/compo";
    import FileIcon from "$lib/compo/FileIcons/FileIcon.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { page } from "$app/stores";
    import { sampleFiles } from "./sample";
    import { params } from "$lib/params";
    import { goto } from "$app/navigation";

    export let files: File[] = []; //sampleFiles

    export let selected;

    const dispatcher = createEventDispatcher();
    const api = getContext("__api__") as RootAPI;
    const pid = $page.params["pid"];

    $: _path = $params["folder"] || "";

    let loading = false;

    const load = async (lpath?: string) => {
        loading = true;
        const resp = await api.listProjectFiles(pid, lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
        loading = false;
    };

    $: load(_path);

    let size = "32";
</script>

<div class="table-container">
    <table class="table table-hover">
        <thead>
            <tr>
                <td />
                <th>Name</th>
                <th>Created At</th>
                <th>Size</th>
                <th>Action</th>
            </tr>
        </thead>
        <tbody>
            {#if loading}
                <Loader />
            {:else}
                {#each files as row, i}
                    <tr
                        on:dblclick={() => dispatcher("open_item", row)}
                        on:click={() => dispatcher("select_item", row)}
                        class="cursor-pointer hover:bg-gray-700 {selected ===
                        row.name
                            ? 'variant-outline-primary'
                            : ''}"
                    >
                        <td class="w-1">
                            {#if selected === row.name}
                                <input type="checkbox" checked={true} />
                            {/if}
                        </td>
                        <td>
                            <button
                                class="mr-1 text-indigo-500"
                                type="button"
                                on:click={() => {
                                    goto(
                                        `/z/pages/portal/project/files/${pid}?folder=${_path ? _path + "/" + row.name : row.name}`,
                                    );
                                }}
                            >
                                {#if row.is_folder}
                                    <FolderIcon {size} />
                                {:else}
                                    <FileIcon {size} name={row.name} />
                                {/if}

                                {row.name}
                            </button>
                        </td>
                        <td>{row.created_at || ""}</td>
                        <td>{row.size || ""}</td>
                        <td>
                            <button
                                type="button"
                                class="btn btn-sm variant-filled"
                            >
                                <SvgIcon name="eye" className="w-4 h-4" />
                            </button>
                        </td>
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>
