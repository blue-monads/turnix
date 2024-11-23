<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { getContext } from "svelte";
    import FilePicker from "../../../../../filestore/FilePicker/FilePicker.svelte";
    import { page } from "$app/stores";

    const pid = $page.params["pid"];

    const modal = getContext("__vs_modal__") as any;
    const rootApi = getContext("__api__") as RootAPI;

    interface Props {
        data?: Record<string, any>;
        onSave: (data: Record<string, any>) => Promise<void>;
    }

    let { data, onSave }: Props = $props();

    let title = $state(data?.title || "");
    let notes = $state(data?.notes || "");
    let attachments: string = $state(data?.attachments || "");

    let files = $derived(attachments.split(",").filter(Boolean));
</script>

<div class="card">
    <header class="card-header">
        <h4 class="h4">{data ? "Edit Note" : "Add Note"}</h4>
    </header>
    <section class="p-4 flex flex-col gap-4">
        <label class="label"
            ><span>Title</span>
            <input
                class="input p-1"
                type="text"
                placeholder="Input"
                bind:value={title}
            />
        </label>

        <label class="label"
            ><span>Notes</span>
            <textarea
                bind:value={notes}
                class="textarea p-1"
                rows="4"
                placeholder="Notes description"
            ></textarea></label
        >

        <div class="flex flex-col gap-2">
            <span>Files</span>


                <div class="flex justify-items-start flex-wrap gap-2 p-2">
                    {#each files as file}
                        <div
                            class="p-1 rounded bg-white relative min-w-20 min-h-20 self-start"
                        >
                            <button
                                class="btn btn-sm absolute top-0 right-0"
                                onclick={async () => {
                                    console.log("@delete_file", file);

                                    try {
                                        await rootApi.deleteShareFile(file);                                        
                                    } catch (error) {                                        
                                        console.error(error);
                                    }
                                    
                                    attachments = attachments.replace(file, "");
                                }}
                            >
                                <SvgIcon name="x-mark" className="w-4 h-4" />
                            </button>

                            <img
                                src={rootApi.getSharedFileURL(file)}
                                alt=""
                                class="w-full max-w-32 h-auto"
                            />
                        </div>
                    {/each}
                </div>

            <button
                class="btn btn-sm variant-filled-secondary w-10"
                onclick={() => {
                    modal.show(FilePicker, {
                        pid: pid,
                        onPick: async (file: any) => {
                            modal.close();

                            const resp = await rootApi.sharedFile(file.id);
                            if (resp.status !== 200) {
                                return;
                            }

                            attachments =
                                attachments === ""
                                    ? resp.data
                                    : attachments + "," + resp.data;
                        },
                    });
                }}
            >
                <SvgIcon name="plus" className="w-6 h-6" />
            </button>
        </div>
    </section>

    <footer class="card-footer flex justify-end">
        <button
            type="button"
            class="btn variant-filled"
            onclick={() => {
                onSave({
                    title,
                    notes,
                    attachments,
                });
            }}
            disabled={!title || !notes}
            >save
        </button>
    </footer>
</div>
