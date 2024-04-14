<script lang="ts">
</script>

<div class="p-2">
    <div class="card">
        <header class="card-header">
            <h4 class="h4">Add Account</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <label class="label">
                <span>Name</span>
                <input class="input p-1" type="text" placeholder="Input" />
            </label>

            <label class="label">
                <span>Type</span>
                <select class="select">
                    <option value="expenses">Expenses</option>
                    <option value="revenue">Revenue</option>
                    <option value="assets">Assets</option>
                    <option value="liabilities">Liabilities</option>
                    <option value="equity">Equity</option>
                </select>
            </label>

            <label class="label">
                <span>Info</span>
                <textarea
                    class="textarea p-1"
                    rows="4"
                    placeholder={"information about account"}
                />
            </label>
        </section>
        <footer class="card-footer flex justify-end">
            <button class="btn variant-filled"> save </button>
        </footer>
    </div>
</div>
