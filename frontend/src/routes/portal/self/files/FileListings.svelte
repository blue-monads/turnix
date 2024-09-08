<script lang="ts">
    import { createEventDispatcher, onMount } from "svelte";
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

    const fileActions = [
        { name: "rename", icon: "pencil-square" },
        { name: "download", icon: "arrow-down-on-square" },
        { name: "delete", icon: "trash" },
    ];

    const folderActions = [
        { name: "rename", icon: "pencil-square" },
        { name: "delete", icon: "trash" },
    ];

    let isActive = false;
    let activeFileId = 0;

    const handler = (e: any) => {
        isActive = false;
        activeFileId = 0;
    };

    onMount(() => {
        document.addEventListener("click", handler);

        return () => {
            document.removeEventListener("click", handler);
        };
    });
</script>

<div class="table-container">
    <table class="table table-hover overflow-auto">
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
                            {#if !isActive || activeFileId === row.id}
                                <button
                                    on:click={expore(row)}
                                    class="btn btn-sm variant-filled"
                                >
                                    <SvgIcon name="eye" className="w-4 h-4" />
                                </button>

                                <button
                                    on:click|stopPropagation={() => {
                                        isActive = true;
                                        activeFileId = row.id;
                                        console.log(
                                            "activeFileId",
                                            activeFileId,
                                        );
                                    }}
                                    class="btn btn-sm variant-filled-secondary relative group transition-all duration-200 focus:overflow-visible overflow-hidden"
                                >
                                    <SvgIcon
                                        name="bars-3"
                                        className="w-4 h-4"
                                    />

                                    <div
                                        class="absolute shadow-lg top-8 -left-16 w-28 h-max p-1 border border-zinc-200 rounded-lg flex flex-col gap-2 variant-filled"
                                    >
                                        {#each row.is_folder ? folderActions : fileActions as action}
                                            <button
                                                on:click|stopPropagation={() => {
                                                    dispatcher("action", {
                                                        action,
                                                        row,
                                                    });
                                                }}
                                                class="flex gap-1 justify-start items-center p-1 rounded-lg hover:bg-white hover:text-secondary-600"
                                            >
                                                <SvgIcon
                                                    name={action.icon}
                                                    className="w-4 h-4"
                                                />

                                                <p>{action.name}</p>
                                            </button>
                                        {/each}
                                    </div>
                                </button>
                            {/if}
                        </td>
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>
