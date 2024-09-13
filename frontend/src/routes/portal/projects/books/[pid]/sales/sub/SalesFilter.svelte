<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
</script>

<div class="p-2">
    <div class="card">
        <header class="card-header">
            <h4 class="h4">Sales Filter</h4>
        </header>

        <section
            class="pl-4 py-2 flex justify-between items-center flex-wrap gap-2"
        >
            <div class="flex gap-2 items-center flex-wrap">
                

                <div class="flex flex-col gap-1">
                    <span class="font-medium">Date range</span>

                    <label class="label flex">
                        <span class="w-24">From</span>
                        <input class="input" type="datetime-local" />
                    </label>

                    <label class="label flex">
                        <span class="w-24">To</span>
                        <input class="input" type="datetime-local" />
                    </label>

                    <label class="label">
                        <select class="select">
                            <option value="">1 Hour</option>
                            <option value="">1 Day</option>
                            <option value="">1 Week</option>
                            <option value="">1 Month</option>
                            <option value="">3 Month</option>
                        </select>
                    </label>

                </div>

                

                <div class="flex flex-col">
                    <label class="label">
                        <span>Client</span>

                        <div class="flex gap-2">
                            <input
                                type="text"
                                class="input p-1"
                                placeholder="Client"
                            />

                            <button class="btn btn-sm variant-filled">
                                <SvgIcon className="h-4 w-4" name="plus" />
                            </button>
                        </div>
                    </label>
                </div>

                <label class="label inline-flex">
                    <span>Show Deleted</span>
                    <input type="checkbox" class="checkbox ml-2" />
                </label>
            </div>


        </section>

        <footer class="card-footer flex justify-end">
            <button
                class="btn btn-sm variant-filled disabled:variant-ghost disabled:text-slate-800"
            >
                <SvgIcon className="h-4 w-4" name="bolt" />

                Run
            </button>
        </footer>

    </div>
</div>
