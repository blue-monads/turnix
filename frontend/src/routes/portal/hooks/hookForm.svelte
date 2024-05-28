<script lang="ts">
    import type { ProjectDef, RootAPI } from "$lib/api";
    import KvEditor from "$lib/compo/autoform/_kv_editor.svelte";
    import Loader from "$lib/compo/loader/loader.svelte";
    import { params } from "$lib/params";

    import { getContext } from "svelte";

    import CodeMirror from "svelte-codemirror-editor";
    import {
        javascript,
        localCompletionSource,
        scopeCompletionSource,
    } from "@codemirror/lang-javascript";
    import { autocompletion } from "@codemirror/autocomplete";
    import { code, completionObject } from "./code";

    export let id = 0;
    export let hook_code = code;
    export let runas_user_id = 0;
    export let hook_type = "webhook";
    export let target = "";

    export let envs = "{}";
    export let extrameta = "{}";
    export let event = "";
    export let name = "";

    export let onSave = async (data: Record<string, any>) => {};

    const ptype = $params["ptype"];

    const api = getContext("__api__") as RootAPI;

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

    const handleExtrametaChange = (data: object) => {
        console.log("handleExtrametaChange", data);
        extrameta = JSON.stringify(data);
    };

    const handleEnvVarChange = (data: object) => {
        console.log("handleEnvVarChange", data);
        envs = JSON.stringify(data);
    };
</script>

{#if loading}
    <Loader />
{:else}
    <div class="flex p-2 flex-col gap-2">
        <div class="card p-4 flex flex-col gap-2">
            <header class="card-header">
                <h4 class="h4">
                    {#if id === 0}
                        New Hook
                    {:else}
                        Edit Hook
                    {/if}
                </h4>
            </header>

            <section class="p-4 flex flex-col gap-4">
                <label class="label">
                    <span>Name</span>
                    <input
                        bind:value={name}
                        class="input p-0.5"
                        type="text"
                        placeholder="name"
                    />
                </label>

                <label class="label">
                    <span>Event Type</span>
                    <select bind:value={event} class="select">
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
                    </select>
                </label>

                {#if hook_type === "script"}
                    <!-- svelte-ignore a11y-label-has-associated-control -->
                    <label class="label">
                        <span>Script Code</span>
                        <div class="p-1 rounded border bg-white">
                            <CodeMirror
                                extensions={[
                                    autocompletion({
                                        activateOnTyping: true,
                                        override: [
                                            localCompletionSource,
                                            scopeCompletionSource(
                                                completionObject,
                                            ),
                                        ],
                                    }),
                                ]}
                                bind:value={hook_code}
                                lang={javascript()}
                            />
                        </div>
                    </label>
                {/if}

                {#if hook_type === "webhook"}
                    <label class="label">
                        <span>Endpoint</span>
                        <input
                            class="input p-0.5"
                            type="url"
                            bind:value={target}
                            placeholder="http://example.com/receive"
                        />
                    </label>
                {/if}

                {#if hook_type === "native_function"}
                    <label class="label">
                        <span>Handler</span>
                        <input
                            class="input p-0.5"
                            type="url"
                            bind:value={target}
                            placeholder="handleXyz"
                        />
                    </label>
                {/if}

                <!-- svelte-ignore a11y-label-has-associated-control -->
                <div class="label">
                    <span>Extrameta</span>
                    <KvEditor
                        data={extrameta ? JSON.parse(extrameta) : {}}
                        onChange={handleExtrametaChange}
                    />
                </div>

                <!-- svelte-ignore a11y-label-has-associated-control -->
                <div class="label">
                    <span>Env Vars</span>
                    <KvEditor
                        sensitive={true}
                        data={envs ? JSON.parse(envs) : {}}
                        onChange={handleEnvVarChange}
                    />
                </div>
            </section>

            <footer class="card-footer flex justify-end">
                <button
                    class="btn btn-sm variant-filled"
                    on:click={async () => {
                        loading = true;

                        await onSave({
                            id,
                            name,
                            hook_code,
                            runas_user_id,
                            hook_type,
                            envs,
                            extrameta,
                            target,
                            event,
                        });

                        loading = false;
                    }}
                >
                    save
                </button>
            </footer>
        </div>
    </div>
{/if}
