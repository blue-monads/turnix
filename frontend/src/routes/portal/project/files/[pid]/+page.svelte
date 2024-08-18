<script>
    import { FolderIcon } from "$lib/compo";
    import FileIcon from "$lib/compo/FileIcons/FileIcon.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { createEventDispatcher } from "svelte";

    export let files = [
        {
            name: "public",
            last_modified: "",
            size: "",
            is_dir: true,
        },
        {
            name: "test.mp4",
            last_modified: "2022/03/45",
            size: "23M",
        },
        {
            name: "test.zip",
            last_modified: "2022/03/45",
            size: "1K",
        },
        {
            name: "test.png",
            last_modified: "2022/03/45",
            size: "1K",
        },
        {
            name: "xyz.mp3",
            last_modified: "2022/03/45",
            size: "1K",
        },
        {
            name: "test.txt",
            last_modified: "2022/03/45",
            size: "1K",
        },
    ];

    export let selected;

    const dispatcher = createEventDispatcher();

    let size = "32";
</script>

<div class="table-container">
    <table class="table table-hover">
        <thead>
            <tr>
                <td />
                <th>Name</th>
                <th>Last Modified</th>
                <th>Size</th>
                <th>Action</th>
            </tr>
        </thead>
        <tbody>
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
                        <span class="mr-1 text-indigo-500">
                            {#if row.is_dir}
                                <FolderIcon {size} />
                            {:else}
                                <FileIcon {size} name={row.name} />
                            {/if}
                        </span>

                        {row.name}
                    </td>
                    <td>{row.last_modified || ""}</td>
                    <td>{row.size || ""}</td>
                    <td>
                        <button
                            type="button"
                            class="btn btn-sm variant-filled"
                        >
                            <SvgIcon name="eye" className="w-4 h-4"  />
                        </button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
