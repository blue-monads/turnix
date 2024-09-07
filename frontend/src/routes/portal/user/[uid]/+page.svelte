<script lang="ts">
    import { RootAPI, type User } from "$lib/api";
    import { getContext } from "svelte";
    import { page } from "$app/stores";
    import Loader from "$lib/compo/loader/loader.svelte";

    const api = getContext("__api__") as RootAPI;
    const uid = $page.params["uid"];

    let user: Partial<User>;

    const load = async () => {
        const resp = await api.userInfo(uid);
        if (resp.status !== 200) {
            return;
        }

        user = resp.data;
    };

    load();

    let message = "";

    const senUserMessage = async () => {
        const uid = $page.params["uid"];
        const resp = await api.messageUser(uid, {
            title: "Hello",
            body: message,
        });

    }


</script>

<svelte:head>
    <title>User Profile</title>
</svelte:head>

<div class="p-2">
    {#if user}
        <div class="card p-2">
            <header class="flex flex-col items-center justify-center">
                <div class="flex flex-col gap-2 p-2">
                    <div class="ml-4 w-16 h-16 rounded-full bg-gray-500"></div>
                    <h3 class="h3">{user?.name}</h3>
                    <span class="text-sm opacity-50">{user?.email || ""}</span>
                    <span class="text-sm opacity-50">{user?.phone || ""}</span>
                    <span class="chip variant-filled">{user?.utype}</span>
                </div>

                <p>
                    {user?.bio || ""}

                </p>
            </header>
            <section class="flex flex-col gap-4">
                <label class="label"
                    ><span class="font-semibold">Send message</span>
                    <textarea
                        bind:value={message}
                        class="textarea"
                        rows="4"
                        placeholder="I am a nobody"
                    ></textarea></label
                >

                <p class="text-xs italic">
                    User might not be able to see this message, based on his/her
                    notification settings
                </p>

                <div class="flex justify-end pt-8 gap-2">
                    <button
                        disabled={message.length === 0}
                        on:click={senUserMessage}
                        class="btn btn-sm variant-filled"
                    >
                        Send
                    </button>
                </div>
            </section>
        </div>
    {:else}
        <Loader />
    {/if}
</div>
