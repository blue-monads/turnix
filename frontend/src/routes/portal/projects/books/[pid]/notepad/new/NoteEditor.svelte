<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import FilePicker from "../../../../../filestore/FilePicker/FilePicker.svelte";

    const modal = getContext("__vs_modal__") as any;
    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);

    
    let title = "";
    let notes = "";
    let attachments = $state("");

    let files = $derived(
        attachments.split(",").map((f) => rootApi.getSharedFileURL(f)),
    );
</script>

<div class="card">
    <header class="card-header"><h4 class="h4">Add Note</h4></header>
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

        <!-- svelte-ignore a11y_label_has_associated_control -->
        <label>
            <span>Files</span>

            <div
                class="mx-auto cursor-pointer flex w-full max-w-lg flex-col items-center justify-center rounded-xl border-2 border-dashed p-2 text-center"
            >
                <div class="flex items-center gap-2">
                    {#each files as file}
                        <div class="flex-shrink-0 w-32 h-32">
                            <img src={file} alt="" class="w-full h-auto" />
                        </div>
                    {/each}
                </div>
            </div>

            <button class="btn btn-sm variant-filled-secondary" onclick={() => {
                modal.show(FilePicker, {
                    onPick: async (file: any) => {
                        modal.close();

                        const resp = await rootApi.sharedFile(file.id);
                        if (resp.status !== 200) {
                            return;
                        }

                        attachments = attachments === "" ? resp.data : attachments + "," + resp.data;
                   }
                });
                    
            }}>
                <SvgIcon name="plus" className="w-6 h-6" />
            </button>


        </label>
    </section>

    <footer class="card-footer flex justify-end">
        <button
            type="button"
            class="btn variant-filled"
            disabled={!title || !notes}
            >save
        </button>
    </footer>
</div>
