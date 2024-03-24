<script lang="ts">
    import { getContext } from "svelte";
    import type { API } from "$lib/api";

    const api = getContext("__api__") as API;

    let message = "";
</script>

<div class="p-2">
    <div class="card">
        <header class="card-header">
            <h4 class="h4">Add Template</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <label class="label">
                <span>Name</span>
                <input class="input p-1" type="text" placeholder="Input" />
            </label>

            <label class="label">
                <span>Type</span>
                <select class="select">
                    <option value="bexec">Browser Execute</option>
                </select>
            </label>

            <label class="label">
                <span>Content script</span>
                <textarea
                    class="textarea p-1"
                    rows="4"
                    placeholder={"(ctx) => {}"}
                />
            </label>

            <label class="flex items-center space-x-2">
                <input class="checkbox" type="checkbox" />
                <p>Can publisher read result ?</p>
            </label>
            
        </section>
        <footer class="card-footer flex justify-end">
            <button class="btn variant-filled"> save </button>
        </footer>
    </div>
</div>
