<script lang="ts">
    import { NewBookAPI, type Catagory } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    const pid = $page.params["pid"];
    const id = $params["prid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let message = $state("");
    let data: any = $state();
    let loading = $state(true);

    let allCatagories: Catagory[] = $state([]);

    const load = async () => {
        loading = true;
        const resp = await api.getProduct(pid, id);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    const loadCats = async () => {
        const resp = await api.listCatagories(pid);
        if (resp.status !== 200) {
            return;
        }
        allCatagories = resp.data;
    };

    loadCats();
    load();

    const submit = async (data: Record<string, any>) => {
        delete data.id;

        const resp = await api.updateProduct(pid, id, data);
        if (resp.status !== 200) {
            message = resp.data.message;
            return;
        }

        goto(`/z/pages/portal/projects/books/${pid}/inventory/products`);
    };

    const setValueText = (key: string) => (ev: any) => {
        data[key] = ev.target["value"];
    }

    const setValueNumber = (key: string) => (ev: any) => {
        data[key] = Number(ev.target["value"]) ;
    }



</script>

{#if loading}
    <Loader />
{:else}
    <div class="p-2">
        <div class="card">
            <header class="card-header">
                <h4 class="h4">Edit Product</h4>

                <p class="text-red-500">{message}</p>
            </header>

            <section class="p-4 flex flex-col gap-4">
                <label class="label">
                    <span>Name</span>
                    <input
                        value={data.name}
                        class="input p-1"
                        type="text"
                        placeholder="Input"
                        onchange={setValueText("name")}
                    />
                </label>

                <label class="label">
                    <span>Catagory</span>
                    
                    <select 
                        value={data.catagory_id} 
                        class="select" 
                        onchange={setValueNumber("catagory_id")}
                        >
                        {#each allCatagories as catagory}
                            <option value={catagory.id}>{catagory.name}</option>
                        {/each}
                    </select>
                </label>

                <label class="label">
                    <span>Variant</span>
                    <input
                        type="text"
                        name="variant_id"
                        value={data.variant_id}
                        onchange={setValueNumber("variant_id")}
                        class="input p-1"
                        placeholder="LARGE XL"
                    />
                </label>

                <label class="label" for="price">
                    <span>Price</span>
                </label>

                <div
                    class="input-group input-group-divider grid-cols-[auto_1fr_auto]"
                >
                    <span class="w-8">
                        <SvgIcon
                            name="currency-dollar"
                            className="w-5 h-5 text-gray-700 inline-flex justify-self-center items-center mt-1 ml-1"
                        />
                    </span>

                    <input
                        id="price"
                        value={data.price}
                        onchange={setValueNumber("price")}
                        class="input p-1"
                        type="number"
                        placeholder="Input"
                    />
                </div>

                <label class="label">
                    <span>In Stock</span>
                    <input
                        id="stock_count"
                        value={data.stock_count}
                        onchange={setValueNumber("stock_count")}
                        class="input p-1"
                        type="number"
                        placeholder="Input"
                    />
                </label>

                <label class="label">
                    <span>Info</span>
                    <textarea
                        value={data.info}
                        class="textarea p-1"
                        onchange={setValueText("info")}
                        rows="4"
                        placeholder={"information about account"}
                    />
                </label>
            </section>
            <footer class="card-footer flex justify-end">
                <button
                    class="btn variant-filled"
                    onclick={() => {
                        submit({ ...data });
                    }}
                >
                    save
                </button>
            </footer>
        </div>
    </div>
{/if}
