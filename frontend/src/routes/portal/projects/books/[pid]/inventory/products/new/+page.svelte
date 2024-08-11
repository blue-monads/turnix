<script lang="ts">
    import { NewBookAPI, type Catagory } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoForm } from "$lib/compo";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let message = "";

    let name = "";
    let info = "";
    let catid = 0;
    let price = 0;

    let allCatagories: Catagory[] = [];

    const load = async () => {
        const resp = await api.listCatagories(pid);
        if (resp.status !== 200) {
            return;
        }
        allCatagories = resp.data;
    };

    load();

    const submit = async () => {
        const resp = await api.addProduct(pid, {
            name,
            info,
            catagory_id: Number(catid),
            price,
        });
        if (resp.status !== 200) {
            message = resp.data.message;
            return;
        }

        goto(`/z/pages/portal/projects/books/${pid}/inventory/products`);
    };
</script>

<form class="p-2" on:submit|preventDefault={submit}>
    <div class="card">
        <header class="card-header">
            <h4 class="h4">Add Product</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <label class="label">
                <span>Name</span>
                <input
                    bind:value={name}
                    class="input p-1"
                    type="text"
                    placeholder="Input"
                />
            </label>

            <label class="label">
                <span>Catagory</span>
                <select bind:value={catid} class="select" required>
                    {#each allCatagories as catagory}
                        <option value={catagory.id}>{catagory.name}</option>
                    {/each}
                </select>
            </label>


            <!-- price -->
            <label class="label">
                <span>Price</span>
                <input
                    bind:value={price}
                    class="input p-1"
                    type="number"
                    placeholder="Input"
                />

            <label class="label">
                <span>Info</span>
                <textarea
                    bind:value={info}
                    class="textarea p-1"
                    rows="4"
                    placeholder={"information about account"}
                />
            </label>


        </section>
        <footer class="card-footer flex justify-end">
            <button type="submit" class="btn variant-filled"> save </button>
        </footer>
    </div>
</form>
