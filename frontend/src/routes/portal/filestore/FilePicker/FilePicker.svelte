<script lang="ts">
    import {
        FileDropzone,
        RadioGroup,
        RadioItem,
    } from "@skeletonlabs/skeleton";
    import FileListings from "../../self/files/FileListings.svelte";
    import { getContext } from "svelte";
    import type { File, RootAPI } from "$lib";
    import FilePreview from "../../self/files/preview/FilePreview.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    interface Props {
        api?: RootAPI;
        pid?: number;
        tabMode?: "all" | "personal" | "project";
        pickMode?: "all" | "file" | "folder" | "images";
        onPick?: (file: File) => void;
    }

    let pp: Props = $props();

    let {
        api = getContext("__api__") as RootAPI,
        pid,
        tabMode = "all",
        onPick,
        pickMode = "all",
    } = pp;

    console.log("@props", $state.snapshot(pp));

    let files: File[] = $state([]);
    let loading = $state(false);
    let personalPath = $state("");
    let projectPath = $state("");
    let uploadFile: any | undefined = $state();
    let uploading = $state(false);

    let previewFile: File | undefined = $state();

    let mode: "listing" | "upload" | "preview" = $state("listing");
    let value: "personal" | "project" = $state(
        tabMode === "project" ? "project" : "personal",
    );

    const onDrop = async (e: any) => {
        console.log("onDrop", e);
        uploadFile = e.target.files[0];
    };

    const loadPersonal = async (lpath?: string) => {
        loading = true;
        const resp = await api.listSelfFiles(lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
        loading = false;
    };

    const loadProject = async (lpid: string, lpath?: string) => {
        loading = true;
        const resp = await api.listProjectFiles(lpid, lpath);
        if (resp.status !== 200) {
            return;
        }

        files = resp.data as any;
    };

    const submitUploadFile = async () => {
        uploading = true;

        if (value === "personal") {
            await api.addSelfFile(uploadFile.name, personalPath, uploadFile);
        } else {
            await api.addProjectFile(
                String(pid),
                uploadFile.name,
                personalPath,
                uploadFile,
            );
        }
        uploadFile = undefined;
        mode = "listing";
    };

    $effect(() => {
        console.log("@effect");

        mode = "listing";
        previewFile = undefined;
        if (value === "personal") {
            files = [];
            loadPersonal(personalPath);
        } else {
            files = [];
            loadProject(String(pid), projectPath);
        }
    });

    let isBackDisabled = $state(false);

    $effect(() => {
        if (mode === "preview") {
            isBackDisabled = false;
            return;
        }

        if (mode === "upload") {
            isBackDisabled = false;
            return;
        }

        if (value === "personal") {
            isBackDisabled = personalPath === "";
        } else {
            isBackDisabled = projectPath === "";
        }
    });

    let _paths = $derived(
        (value === "personal" ? personalPath : projectPath).split("/"),
    );

    let selected = $state("");
</script>

<div class="w-full bg-white md:min-w-[724px] h-full">
    <div class="flex justify-between bg-slate-200 items-center px-2 rounded">
        <div class="flex items-center">
            <h4 class="h4 py-2">File Picker</h4>
        </div>

        <div class="flex items-center">
            {#if tabMode === "all"}
                <RadioGroup>
                    <RadioItem
                        bind:group={value}
                        name="justify"
                        value={"personal"}
                        padding="p-0.5 px-2 text-sm"
                        >Personal
                    </RadioItem>
                    <RadioItem
                        bind:group={value}
                        name="justify"
                        value={"project"}
                        padding="p-0.5 px-2 text-sm"
                    >
                        Project
                    </RadioItem>
                </RadioGroup>
            {/if}
        </div>
    </div>

    <div class="flex justify-between p-1">
        <ol class="breadcrumb">
            <li class="crumb">
                <button
                    onclick={() => {
                        personalPath = "";
                        projectPath = "";
                        mode = "listing";
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="home" />
                </button>
            </li>

            {#each _paths as path, i}
                {#if path}
                    <li class="crumb">
                        {path}
                    </li>
                    <li class="crumb-separator" aria-hidden="true">/</li>
                {/if}
            {/each}

            <li class="crumb">
                {previewFile?.name}
            </li>
        </ol>
    </div>

    <div class="py-2 text-xs">
        {#if mode === "listing"}
            <FileListings
                {files}
                loading={false}
                pickMode={true}
                bind:selected
                hidePreview={true}
                onExplore={(row) => {
                    if (row.is_folder) {
                        const nextPath = row.path
                            ? `${row.path}/${row.name}`
                            : row.name;
                        if (value === "personal") {
                            personalPath = nextPath;
                        } else {
                            projectPath = nextPath;
                        }
                    } else {
                        previewFile = row;
                        mode = "preview";
                    }

                    selected = "";

                    console.log("@@explore", row.path);
                }}
            />
        {:else if mode === "upload"}
            <FileDropzone on:change={onDrop} name="files">
                <div slot="lead" class="flex justify-center">
                    <SvgIcon name="arrow-up-tray" className="w-6 h-6" />
                </div>

                <svelte:fragment slot="message">
                    {#if uploadFile}
                        <div class="flex-shrink-0 w-32 h-32">
                            <img
                                src={URL.createObjectURL(uploadFile)}
                                alt=""
                                class="w-full h-auto"
                            />
                        </div>
                        <span>
                            {uploadFile.name}
                        </span>
                    {:else}
                        <strong>Upload a file</strong> or drag and drop
                    {/if}
                </svelte:fragment>
            </FileDropzone>

            <div class="flex justify-end py-2">
                <button
                    class="btn btn-sm variant-filled"
                    disabled={!uploadFile || uploading}
                    onclick={submitUploadFile}
                >
                    {#if loading}
                        Uploading...
                    {:else}
                        Upload
                    {/if}
                </button>
            </div>
        {:else if previewFile}
            <div
                class="w-auto h-full max-h-96 max-w-xs md:max-w-screen-md overflow-auto"
            >
                <FilePreview
                    {api}
                    fileId={String(previewFile?.id)}
                    filename={previewFile?.name as string}
                />
            </div>
        {/if}
    </div>

    <div class="flex justify-between">
        <button
            class="btn btn-sm bg-gray-100"
            disabled={isBackDisabled}
            onclick={() => {
                if (mode === "preview" || mode === "upload") {
                    previewFile = undefined;
                    mode = "listing";
                } else {
                    if (value === "personal") {
                        personalPath = personalPath.substring(
                            0,
                            personalPath.lastIndexOf("/"),
                        );
                    } else {
                        projectPath = projectPath.substring(
                            0,
                            projectPath.lastIndexOf("/"),
                        );
                    }
                }
            }}
        >
            <SvgIcon className="h-4 w-4" name="chevron-left" />
        </button>

        <div class="flex gap-2">
            {#if mode !== "upload"}
                <button
                    class="btn btn-sm variant-filled-primary"
                    onclick={() => {
                        mode = "upload";
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="cloud-arrow-up" />
                    <span class="hidden md:inline">Upload</span>
                </button>

                <button
                    class="btn btn-sm variant-filled"
                    disabled={selected === ""}
                    onclick={() => {
                        const picked = files.find((f) => f.name === selected);
                        if (picked) {
                            onPick?.(picked);
                            return;
                        }
                        console.log("@clicked");
                    }}
                >
                    <SvgIcon className="h-4 w-4" name="check" />
                    <span class="hidden md:inline">Pick</span>
                </button>
            {/if}
        </div>
    </div>
</div>
