<script lang="ts">
    import { RootAPI, type User } from "$lib/api";
    import { getContext } from "svelte";
    import { FloatyButton, Loader } from "$lib/compo";
    import { AppBar } from "@skeletonlabs/skeleton";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    const api = getContext("__api__") as RootAPI;

    let users: User[] = [];
    let loading = true;

    const load = async () => {
        loading = true;
        const resp = await api.listSelfUsers();
        if (resp.status !== 200) {
            return;
        }

        users = resp.data;
        loading = false;
    };

    const deleteUser = async (user: User) => {
        await api.deleteSelfUser(String(user.id));

        load();
    };

    load();
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4">Users</h4>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <a
            class="btn btn-sm variant-soft-secondary"
            href="/z/pages/portal/self/users/new"
        >
            <SvgIcon className="h-4 w-4" name="plus" />

            Add User
        </a>
    </svelte:fragment>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <div class="p-2">
        {#if users.length === 0}
            <div class="alert alert-info">
                <div class="flex-auto">
                    <span class="ml-2">No users</span>
                </div>
            </div>
        {:else}
            <div class="card p-2">
                <dl class="list-dl">
                    {#each users as user}
                        <div>
                            <span class="badge-icon p-4 variant-soft-secondary">
                                <i class="fa-solid fa-book"></i>
                            </span>

                            <span class="flex-auto"
                                ><dt class="font-bold">{user.name}</dt>
                                <dd class="text-sm opacity-50">
                                    {user.bio}
                                </dd></span
                            >
                            <div class="flex gap-0 sm:gap-1 flex-wrap space-y-2">



                                <a
                                
                                href="/z/pages/portal/user/{user.id}"
                                class="btn btn-sm variant-filled-tertiary"
                                    >Profile
                            </a>

                                <a
                                
                                href="/z/pages/portal/self/users/reset-password?uid={user.id}"
                                class="btn btn-sm variant-filled-primary"
                                    >Reset Password
                            </a>
                                <a
                                    href="/z/pages/portal/self/users/edit?uid={user.id}"
                                    class="btn btn-sm variant-filled-warning"
                                    >Edit
                                </a>
                                <button 
                                    class="btn btn-sm variant-filled-error"
                                    on:click={() => deleteUser(user)}
                                    >Delete</button
                                >
                            </div>
                        </div>
                    {/each}
                </dl>
            </div>
        {/if}
    </div>
{/if}
