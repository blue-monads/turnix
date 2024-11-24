<script lang="ts">
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { goto } from "$app/navigation";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";

    const rootApi = getContext("__api__") as RootAPI;

    let name = $state("");
    let bio = $state("");
    let loading = $state(true);

    const uid = $params["uid"];

    const load = async () => {
        loading = true;
        const resp = await rootApi.getSelfUser(uid);
        if (resp.status !== 200) {
            return;
        }

        name = resp.data.name;
        bio = resp.data.bio;
        loading = false;
    };

    load();

    const update = async () => {
        const resp = await rootApi.updateSelfUser(uid, {
            name,
            bio,
        });
        if (resp.status !== 200) {
            return;
        }

        goto(`/z/pages/portal/self/users`);
    };
</script>

<div class="p-2">
    {#if loading}
        <Loader />
    {:else}
        <div class="card p-4">
            <header class="flex justify-center">
                <div class="flex flex-col gap-2">
                    <h3 class="h3">Update User</h3>
                </div>
            </header>

            <section class="">
                <form class="flex flex-col gap-4">
                    <label class="label"
                        ><span>Name</span>
                        <input
                            class="input p-1"
                            type="text"
                            required
                            bind:value={name}
                        /></label
                    >

                    <label class="label">
                        <span>Bio</span>
                        <textarea
                            class="textarea p-1"
                            bind:value={bio}
                            rows={4}
                        >
                        </textarea>
                    </label>

                    <div class="flex justify-end pt-8 gap-2">
                        <button
                            type="button"
                            class="btn btn-sm variant-filled"
                            onclick={update}
                        >
                            Update
                        </button>
                    </div>
                </form>
            </section>
        </div>
    {/if}
</div>
