<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { AppBar } from "@skeletonlabs/skeleton";

    import { RootAPI, type User } from "$lib/api";
    import { getContext } from "svelte";
    import Loader from "$lib/compo/loader/loader.svelte";

    const api = getContext("__api__") as RootAPI;

    let user: User | undefined = $state();
    let loading = $state(true);

    let name = $state("");
    let bio = $state("");

    const load = async () => {
        loading = true;
        const resp = await api.getSelfSelf();
        if (resp.status !== 200) {
            return;
        }

        user = resp.data;
        name = user.name;
        bio = user.bio;

        loading = false;
    };

    load();
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4">Profile</h4>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <a
            href="/z/pages/portal/self/devices"
            class="btn variant-filled btn-sm mr-1"
        >
            <SvgIcon className="w-4 h-4" name="rectangle-group" />
            Devices
        </a>
        <a
            href="/z/pages/portal/self/inbox"
            class="btn variant-filled-secondary btn-sm mr-1"
        >
            <SvgIcon className="w-4 h-4" name="inbox" />
            Inbox
        </a>
    </svelte:fragment>
</AppBar>

<div class="p-2">
    {#if loading}
        <Loader />
    {:else}
        <div class="card p-2">
            <header class="flex justify-center">
                <div class="flex flex-col gap-2">
                    <div class="w-16 h-16 rounded-full bg-gray-500"></div>
                </div>
            </header>

            <section class="flex flex-col gap-4">
                <label class="label">
                    <span>Name</span>
                    <input
                        class="input p-1"
                        type="text"
                        placeholder="John"
                        bind:value={name}
                    />
                </label>

                <label class="label">
                    <span>Email</span>
                    <input
                        class="input p-1"
                        type="email"
                        value={user?.email}
                        disabled
                        placeholder="user@example.com"
                    />
                </label>

                <label class="label">
                    <span>User Type</span>
                    <input
                        class="input p-1"
                        type="text"
                        value={user?.utype}
                        disabled
                        placeholder="user"
                    />
                </label>

                <label class="label">
                    <span>Bio</span>
                    <textarea
                        class="textarea"
                        bind:value={bio}
                        rows="4"
                        placeholder="I am a nobody"
                    >
                    </textarea>
                </label>
            </section>

            <footer class="flex justify-end p-2 gap-2">
                <a
                    href="/z/pages/portal/self/change-password"
                    class="btn variant-ghost-secondary"
                >
                    Change Password
                </a>

                <a
                    href="/z/pages/portal/self/reset-email"
                    class="btn variant-ghost-tertiary"
                >
                    Reset email
                </a>

                <button class="btn variant-filled" disabled> Update </button>
            </footer>
        </div>
    {/if}
</div>
