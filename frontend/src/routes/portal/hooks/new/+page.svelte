<script lang="ts">
    import type { ProjectDef, RootAPI } from "$lib/api";
    import KvEditor from "$lib/compo/autoform/_kv_editor.svelte";
    import Loader from "$lib/compo/loader/loader.svelte";
    import { params } from "$lib/params";
    import { AppBar } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";

    const ptype = $params["ptype"];

    export let codeValue = `const handle = (ctx) => {

    }`;

    let hook_type = "webhook";

    const api = getContext("__api__") as RootAPI;
    let runas_user_id = 0;
    let runas_specific_user = false;

    let data: ProjectDef;
    let loading = true;

    let usersIndex: Record<number, string> = {};

    const load = async () => {
        loading = true;
        const resp = await api.getProjectType(ptype);
        if (resp.status !== 200) {
            return;
        }
        data = resp.data;
        loading = false;
    };

    load();
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books"
                    >Projects</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Hooks</li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>New</li>
        </ol>
    </svelte:fragment>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <div class="flex p-2 flex-col gap-2">
        <div class="card p-4 flex flex-col gap-2">
            <header class="card-header">
                <h4 class="h4">New Hook</h4>
            </header>

            <section class="p-4 flex flex-col gap-4">
                <label class="label">
                    <span>Name</span>
                    <input class="input" type="text" placeholder="Input" />
                </label>

                <label class="label">
                    <span>Event Type</span>
                    <select class="select">
                        {#each data.event_types as evt}
                            <option value={evt}>{evt}</option>
                        {/each}
                    </select>
                </label>

                <label class="flex items-center space-x-2">
                    <input
                        class="checkbox"
                        type="checkbox"
                        checked={runas_specific_user}
                        on:change={(ev) => {
                            if (runas_specific_user) {
                                runas_user_id = 0;
                                runas_specific_user = false;
                            } else {
                                runas_specific_user = true;
                            }
                        }}
                    />
                    <p>Run as Specific User</p>
                </label>

                {#if runas_specific_user || runas_user_id}
                    <div class="flex gap-2 items-center">
                        <label class="label">
                            <span>Run as User</span>
                            <input
                                class="input p-1"
                                value={usersIndex[runas_user_id] || ""}
                                type="text"
                                disabled
                                placeholder="John"
                            />
                        </label>

                        <div class="pt-4 inline-flex justify-self-center">
                            <button class="variant-filled btn btn-sm"
                                >pick</button
                            >
                        </div>
                    </div>
                {/if}

                <label class="label">
                    <span>Hook Type</span>
                    <select bind:value={hook_type} class="select">
                        <option value="script">Script</option>
                        <option value="native_function">Native Function</option>
                        <option value="webhook">Webhook</option>
                        <option value="webhook">Webhook</option>
                        <option value="webhook">Webhook</option>
                    </select>
                </label>

                {#if hook_type === "script"}
                    <label class="label">
                        <span>Script Code</span>
                        <textarea
                            bind:value={codeValue}
                            class="textarea"
                            rows="10"
                        />
                    </label>
                {/if}

                {#if hook_type === "webhook"}
                    <label class="label">
                        <span>Endpoint</span>
                        <input
                            class="input"
                            type="url"
                            placeholder="http://example.com/receive"
                        />
                    </label>
                {/if}

                <!-- svelte-ignore a11y-label-has-associated-control -->
                <label class="label">
                    <span>Extrameta</span>
                    <KvEditor />
                </label>
            </section>

            <footer class="card-footer flex justify-end">
                <button class="btn btn-sm variant-filled"> save </button>
            </footer>
        </div>
    </div>
{/if}
