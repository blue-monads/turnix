<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
</script>

<div class="p-2">
    <div class="card p-2">
        <header class="card-header">
            <h4 class="h4">Sales Filter</h4>
        </header>

        <section
            class="pl-4 py-2 flex justify-between items-center flex-wrap gap-2"
        >
            <div class="flex gap-2 items-center flex-wrap">
                <label class="label">
                    <span>Payment Status</span>

                    <select class="select">
                        <option value="paid">paid</option>
                        <option value="paid">not paid</option>
                    </select>
                </label>

                <label class="label">
                    <span>From Date</span>
                    <input class="input" type="datetime-local" />
                </label>

                <label class="label">
                    <span>To Date</span>
                    <input class="input" type="datetime-local" />
                </label>

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

                <label class="label">
                    <span>Include Deleted</span>
                    <input type="checkbox" class="checkbox" />
                </label>
            </div>

            <div>
                <button
                    class="btn btn-sm variant-filled disabled:variant-ghost disabled:text-slate-800"
                >
                    <SvgIcon className="h-4 w-4" name="bolt" />

                    Run
                </button>
            </div>
        </section>
    </div>
</div>
