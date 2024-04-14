<script lang="ts">
    import Icon from "@krowten/svelte-heroicons/Icon.svelte";
    import { FileDropzone } from "@skeletonlabs/skeleton";

    let files: FileList;
</script>

<div class="p-2">
    <div class="card">
        <header class="card-header">
            <h4 class="h4">Add Transaction</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <label class="label">
                <span>Title</span>
                <input class="input p-1" type="text" placeholder="title" />
            </label>

            <div class="label border p-2 flex flex-col">
                <span>Lines</span>
                <div class="flex flex-row justify-start gap-4 items-center">
                    <label class="label">
                        <span>Debit Account</span>
                        <input
                            class="input p-1"
                            type="text"
                            disabled
                            placeholder="Petty cash"
                        />
                    </label>

                    <button class="pt-4 cursor-pointer">
                        <Icon name="arrow-up-right" class="w-5 h-5" />
                    </button>

                    <label class="label">
                        <span>Amount</span>
                        <input class="input p-1" type="number" />
                    </label>
                </div>

                <div class="flex flex-row justify-start gap-4">
                    <label class="label">
                        <span>Credit Account</span>
                        <input
                            class="input p-1"
                            type="text"
                            disabled
                            placeholder="Bank XYZ"
                        />
                    </label>

                    <button class="pt-4 cursor-pointer">
                        <Icon name="arrow-up-right" class="w-5 h-5" />
                    </button>

                    <label class="label">
                        <span>Amount</span>
                        <input class="input p-1" type="number" />
                    </label>
                </div>
            </div>

            <label class="label">
                <span>Notes</span>
                <textarea
                    class="textarea p-1"
                    rows="4"
                    placeholder={"Extra information about txn"}
                />
            </label>

            <!-- svelte-ignore a11y-label-has-associated-control -->
            <label class="label">
                <span>Attachments</span>
                <FileDropzone bind:files name="files">
                    <div slot="lead" class="flex justify-center">
                        <Icon name="arrow-up-tray" class="w-6 h-6" />
                    </div>

                    <svelte:fragment slot="message">
                        <strong>Upload a file</strong> or drag and drop
                    </svelte:fragment>
                    <svelte:fragment slot="meta"
                        >PNG, JPG or PDF</svelte:fragment
                    >
                </FileDropzone>
            </label>

            <label class="label">
                <span>Reference ID</span>
                <input class="input p-1" type="text" placeholder="REF-123" />
            </label>
        </section>
        <footer class="card-footer flex justify-end">
            <button class="btn variant-filled"> save </button>
        </footer>
    </div>
</div>
