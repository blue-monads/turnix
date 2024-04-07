<script lang="ts">
    import { gotoPorjects } from "$lib/nav";
    import { AutoForm, PageLayout, AutoTable } from "$lib/compo";

    import type { API } from "$lib/api";
    import { getContext } from "svelte";

    const api = getContext("__api__") as API;

    let message = "";
</script>

<div class="p-2">
    <div class="card p-2">
        <header class="flex justify-center">
            <div class="flex flex-col gap-2">
                <h3 class="h3">New Project</h3>
            </div>
        </header>

        <section class="flex flex-col gap-4">
            <label class="label">
                <span>Name</span>
                <input
                    class="input p-1"
                    type="text"
                    placeholder="TestProject1"
                />
            </label>

            <label class="label">
                <span>Project Type</span>
                <select class="select" disabled>
                    <option value="5">onloop</option>
                </select>
            </label>

            <label class="label">
                <span>Info</span>
                <textarea
                    class="textarea"
                    rows="4"
                    placeholder="This is a project description"
                />
            </label>
        </section>

        <section class="mt-2">
            <AutoForm
                data={{}}
                {message}
                schema={{
                    fields: [
                        { ftype: "TEXT", key_name: "name", name: "Name" }
                    ],
                    name: "",
                    required_fields: ["name"],

                }}               
            />
        </section>

        <footer class="flex justify-end p-2 gap-2">
            <button class="btn variant-filled" disabled> Update </button>
        </footer>
    </div>
</div>
