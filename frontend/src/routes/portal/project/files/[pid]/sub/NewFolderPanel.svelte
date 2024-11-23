<script lang="ts">
    import { getModalStore } from "@skeletonlabs/skeleton";

    const store = getModalStore();


    let value = $state("");

    const onNewName = $store[0].meta["onNewName"];


</script>

<div class="w-full w-modal max-w-md bg-white rounded">
    <div class="text-center">
        <h2 class="text-xl font-bold text-gray-900">New Folder!</h2>
    </div>
    <div class="p-4 flex flex-col gap-2">
        <label class="text-sm font-bold text-gray-500 tracking-wide"
            >Name
            <input
                bind:value
                class="text-base p-2 border border-gray-300 rounded-lg focus:outline-none focus:border-indigo-500 w-full"
                type="text"
                placeholder="new_folder"
            />
        </label>

        <div class="inline-flex justify-end">
            <button
                disabled={value === ""}
                onclick={() => {
                    onNewName(value)
                    store.close()
                }}
                class="btn variant-filled-primary">Create</button
            >
        </div>
    </div>
</div>
