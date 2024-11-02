<script lang="ts">
    interface Props {
        title?: string;
        notes?: string;
        attachments?: string;
    }

    let { title = "", notes = "", attachments = "" }: Props = $props();
</script>

<div class="card">
    <header class="card-header"><h4 class="h4">Add Note</h4></header>
    <section class="p-4 flex flex-col gap-4">
        <label class="label"
            ><span>Title</span>
        <input class="input p-1" type="text" placeholder="Input" bind:value={title} />
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

        <label>
            <span>Files</span>

            <div
                class="mx-auto cursor-pointer flex w-full max-w-lg flex-col items-center justify-center rounded-xl border-2 border-dashed p-2 text-center"
            >
                <svg
                    class="h-10 w-10 text-secondary-500"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                >
                    <path
                        d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
                    />
                </svg>

                <h2
                    class="mt-4 text-xl font-medium text-gray-700 tracking-wide"
                >
                    Category image
                </h2>

                <p class="mt-2 text-gray-500 tracking-wide">
                    Upload or drag & drop your file PNG, PDF, JPG or GIF.
                </p>

                <input
                    type="file"
                    class="hidden"
                    name="category_image"
                    accept="image/png, image/jpeg, image/webp"
                />
            </div>
        </label>
    </section>

    <footer class="card-footer flex justify-end">
        <button type="submit" class="btn variant-filled">save</button>
    </footer>
</div>
