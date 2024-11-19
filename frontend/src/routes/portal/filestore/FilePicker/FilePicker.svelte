<script lang="ts">
    import { RadioGroup, RadioItem } from "@skeletonlabs/skeleton";
    import FileListings from "../../self/files/FileListings.svelte";
    import { getContext } from "svelte";
    import type { File, RootAPI } from "$lib";
    import FilePreview from "../../self/files/preview/FilePreview.svelte";

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
        mode = "listing"
        if (value === "personal") {
            files = [];
            loadPersonal(personalPath);            
        } else {
            files = [];
            loadProject(String(pid), projectPath);
        }
    });
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
                            personalPath = row.path;
                        } else {
                            projectPath = row.path;
                        }
                    } else {
                        previewFile = row;
                        mode = "preview";
                    }

                    console.log("@@explore", row);
                }}
            />
        {:else if previewFile}
            <div class="max-h-96 max-w-screen-md w-auto h-full ">
                <FilePreview
                    {api}
                    fileId={String(previewFile?.id)}
                    filename={previewFile?.name as string}
                />
            </div>
        {/if}
    </div>
</div>
