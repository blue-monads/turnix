<script lang="ts">
    let name = "";
    let file: FileList | null = null;
</script>

<div class="p-2">
    <div class="card p-2">
        <header class="flex justify-center">
            <div class="flex flex-col gap-2">
                <h3 class="h3">Import Project</h3>
            </div>
        </header>

        <section class="flex flex-col gap-4">
            <label class="label">
                <span>Name</span>
                <input
                    bind:value={name}
                    class="input p-1"
                    type="text"
                    placeholder="TestProject1"
                />
            </label>

            <label class="label">
                <span>File</span>
                <input type="file" bind:files={file} required />
            </label>
        </section>

        <header class="flex justify-end p-2">
            <button disabled={!name || !file} class="btn btn-sm variant-filled" on:click={() => {

            }}
                >Submit</button
            >
        </header>
    </div>

    
</div>
