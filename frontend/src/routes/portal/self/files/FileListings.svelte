<script lang="ts">
    import { createEventDispatcher, getContext, onMount } from "svelte";
    import type { File, RootAPI } from "$lib/api";
    import { FolderIcon, Loader } from "$lib/compo";
    import FileIcon from "$lib/compo/FileIcons/FileIcon.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { goto } from "$app/navigation";
    import { getModalStore } from "@skeletonlabs/skeleton";

    interface Props {
        files?: File[];
        loading?: boolean;
        baseUrl?: string;
        path?: string;
        selected: string;
        hidePreview?: boolean;
        onExplore?: (row: File) => void;
    }

    const explore = (row: File) => {
        console.log("@exp");
        if (row.is_folder) {
            goto(
                `${baseUrl}?folder=${path ? path + "/" + row.name : row.name}`,
            );
        } else {
            goto(
                `${baseUrl}/preview?folder=${path}&fid=${row.id}&filename=${row.name}`,
            );
        }
    };

    let {
        files = [],
        loading = false,
        baseUrl = "/z/pages/portal/self/files",
        path = "",
        selected = $bindable(),
        hidePreview = false,
        onExplore = explore
    }: Props = $props();

    let api = getContext("__api__") as RootAPI;

    const dispatcher = createEventDispatcher();
    const store = getModalStore();

    let size = "32";
    const fileActions = [
        { name: "rename", icon: "pencil-square" },
        { name: "download", icon: "arrow-down-on-square" },
        { name: "delete", icon: "trash" },
    ];

    const folderActions = [
        { name: "rename", icon: "pencil-square" },
        { name: "delete", icon: "trash" },
    ];

    let isActive = $state(false);
    let activeFileId = $state(0);
    let dropdownPosition = $state({ top: 0, left: 0 });

    const handler = (e: any) => {
        isActive = false;
        activeFileId = 0;
    };

    $effect(() => {
        document.addEventListener("click", handler);

        return () => {
            document.removeEventListener("click", handler);
        };
    });

    const bytesToHumanReadable = (bytes: number) => {
        const sizes = ["Bytes", "KB", "MB", "GB", "TB"];
        if (bytes === 0) return "0 Byte";
        const i = parseInt(
            String(Math.floor(Math.log(bytes) / Math.log(1024))),
        );
        if (i === 0) return `${bytes} ${sizes[i]}`;
        return `${(bytes / 1024 ** i).toFixed(1)} ${sizes[i]}`;
    };

    const previewFile = (row: File) => {
        store.trigger({
            type: "component",
            component: "file_preview_dialog",
            meta: {
                api,
                fileId: row.id,
                filename: row.name,
            },
        });
    };

    const toggleDropdown = (event: MouseEvent, row: File) => {
        event.stopPropagation();
        const button = event.currentTarget as HTMLButtonElement;
        const rect = button.getBoundingClientRect();
        dropdownPosition = {
            top: rect.bottom + window.scrollY,
            left: rect.left + window.scrollX - 5
        };
        isActive = !isActive || activeFileId !== row.id;
        activeFileId = isActive ? row.id : 0;
    };
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
                        ondblclick={() => dispatcher("open_item", row)}
                        onclick={() => dispatcher("select_item", row)}
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
                                onclick={(ev) => {
                                    console.log("@clicked1", onExplore);
                                    onExplore(row)
                                    console.log("@clicked2", onExplore);

                                    ev.preventDefault();
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
                        <td>
                            {#if row.created_at}
                                {new Date(row.created_at).toLocaleString()}
                            {/if}
                        </td>
                        <td>
                            {#if row.size}
                                {bytesToHumanReadable(row.size)}
                            {/if}
                        </td>
                        <td>
                            {#if !row.is_folder && !hidePreview}
                                <button
                                    onclick={() => previewFile(row)}
                                    class="btn btn-sm variant-filled"
                                >
                                    <SvgIcon
                                        name="eye"
                                        className="w-4 h-4"
                                    />
                                </button>
                            {/if}

                            <button
                                onclick={(e) => toggleDropdown(e, row)}
                                class="btn btn-sm variant-filled-secondary relative group transition-all duration-200 focus:overflow-visible overflow-hidden"
                            >
                                <SvgIcon
                                    name="bars-3"
                                    className="w-4 h-4"
                                />
                            </button>

                            {#if isActive && activeFileId === row.id}
                                <div
                                    class="fixed z-10 shadow-lg w-28 h-max p-1 border border-zinc-200 rounded-lg flex flex-col gap-2 variant-filled"
                                    style="top: {dropdownPosition.top}px; left: {dropdownPosition.left}px;"
                                >
                                    {#each row.is_folder ? folderActions : fileActions as action}
                                        <button
                                            onclick={((ev) => {
                                                ev.stopPropagation();

                                                dispatcher("action", {
                                                    action,
                                                    row,
                                                });
                                            })}
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
                            {/if}
                        </td>
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>