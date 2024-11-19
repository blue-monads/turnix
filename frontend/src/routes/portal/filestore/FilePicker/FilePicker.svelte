<script lang="ts">
    import { RadioGroup, RadioItem } from "@skeletonlabs/skeleton";
    import FileListings from "../../self/files/FileListings.svelte";
    import { getContext } from "svelte";
    import type { File, RootAPI } from "$lib";
    import FilePreview from "../../self/files/preview/FilePreview.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    interface Props {
        api?: RootAPI;
        pid?: number;
    }

    let { api = getContext("__api__") as RootAPI, pid }: Props = $props();

    let files = $state([]);
    let loading = $state(false);
    let personalPath = $state("");
    let projectPath = $state("");

    let previewFile: File | undefined = $state();

    let mode: "listing" | "preview" = $state("listing");
    let value: "personal" | "project" = $state("personal");

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

    $effect(() => {
        console.log("@effect");

        mode = "listing";
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

        if (value === "personal") {
            isBackDisabled = personalPath === "";
        } else {
            isBackDisabled = projectPath === "";
        }
    });

    let _paths = $derived(
        (value === "personal" ? personalPath : projectPath).split("/"),
    );
</script>

<div class="w-full bg-white">
    <div class="flex justify-between bg-slate-200 items-center px-2 rounded">
        <div class="flex items-center">
            <h4 class="h4 py-2">File Picker</h4>
        </div>

        <div class="flex items-center">
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
                <li class="crumb">
                    {path}
                </li>
                <li class="crumb-separator" aria-hidden="true">/</li>
            {/each}
        </ol>
    </div>

    <div class="py-2 text-xs">
        {#if mode === "listing"}
            <FileListings
                {files}
                loading={false}
                selected=""
                hidePreview={true}
                onExplore={(row) => {
                    if (row.is_folder) {
                        if (value === "personal") {
                            personalPath = `${row.path}/${row.name}`;
                        } else {
                            projectPath = `${row.path}/${row.name}`;
                        }
                    } else {
                        previewFile = row;
                        mode = "preview";
                    }

                    console.log("@@explore", row.path);
                }}
            />
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

    <div class="flex justify-start">
        <button
            class="btn btn-sm bg-gray-100"
            disabled={isBackDisabled}
            onclick={() => {
                personalPath = "";
                projectPath = "";
                mode = "listing";
            }}
        >
            <SvgIcon className="h-4 w-4" name="chevron-left" />
        </button>
    </div>
</div>
