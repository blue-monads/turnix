<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import type { File } from "$lib/api";
    import { FolderIcon, Loader } from "$lib/compo";
    import FileIcon from "$lib/compo/FileIcons/FileIcon.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { goto } from "$app/navigation";

    export let files: File[] = [];
    export let loading = false;
    export let baseUrl = "/z/pages/portal/self/files";
    export let path = "";
    export let selected: string;

    const dispatcher = createEventDispatcher();

    let size = "32";

    const expore = (row: File) => () => {
        if (row.is_folder) {
            goto(
                `${baseUrl}?folder=${path ? path + "/" + row.name : row.name}`,
            );
        } else {
            goto(`${baseUrl}/preview?folder=${path}`);
        }
    };

    const actions = [
        { name: "rename", icon: "pencil-square" },
        { name: "download", icon: "arrow-down-on-square" },
        { name: "delete", icon: "trash" },
    ];
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
                                on:click|preventDefault={expore(row)}
                            >
                                {#if row.is_folder}
                                    <FolderIcon {size} />
                                {:else}
                                    <FileIcon {size} name={row.name} />
                                {/if}

                                {row.name}
                            </button>
                        </td>
                        <td>
                            {#if row.created_at}
                                {new Date(row.created_at).toLocaleString()}
                            {/if}
                        </td>
                        <td>{row.size || ""}</td>
                        <td>
                            <button
                                on:click={expore(row)}
                                class="btn btn-sm variant-filled"
                            >
                                <SvgIcon name="eye" className="w-4 h-4" />
                            </button>

                            <button
                                class="btn btn-sm variant-filled-secondary relative group transition-all duration-200 focus:overflow-visible overflow-hidden"
                            >
                                <SvgIcon name="bars-3" className="w-4 h-4" />

                                <div
                                    class="absolute shadow-lg top-8 -left-4 w-32 h-max p-1 bg-white border border-zinc-200 rounded-lg flex flex-col gap-2 z-100"
                                >
                                    {#each actions as action}
                                        <span
                                            class="flex gap-1 justify-start items-center hover:bg-zinc-100 p-1 rounded-lg"
                                        >
                                            <SvgIcon
                                                name={action.icon}
                                                className="w-4 h-4"
                                            />

                                            <p>{action.name}</p>
                                        </span>
                                    {/each}
                                </div>
                            </button>
                        </td>
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>
