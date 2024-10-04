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



    let run_as_type: "specific_user" | "context_user" | "no_user" =
        $state(runas_user_id === -1
            ? "no_user"
            : runas_user_id === 0
              ? "context_user"
              : "specific_user");

    interface Props {
        id?: number;
        hook_code?: any;
        runas_user_id?: any;
        hook_type?: string;
        target?: string;
        envs?: string;
        extrameta?: string;
        event?: string;
        name?: string;
        onSave?: any;
    }

    let {
        id = 0,
        hook_code = $bindable(code),
        runas_user_id = $bindable(-1),
        hook_type = $bindable("webhook"),
        target = $bindable(""),
        envs = $bindable("{}"),
        extrameta = $bindable("{}"),
        event = $bindable(""),
        name = $bindable(""),
        onSave = async (data: Record<string, any>) => {}
    }: Props = $props();

    const ptype = $params["ptype"];

    const api = getContext("__api__") as RootAPI;

    let data: ProjectDef = $state();
    let loading = $state(true);

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
                        {#each data.event_types || [] as evt}
                            <option value={evt}>{evt}</option>
                        {/each}
                    </select>
                </label>

                <label class="flex items-center space-x-2">
                    <p>Run as</p>
                    <select bind:value={run_as_type} class="select w-44">
                        <option value="specific_user">Specific User</option>
                        <option value="context_user">Context User</option>
                        <option value="no_user">No User</option>
                    </select>
                </label>

                {#if run_as_type === "specific_user"}
                    <div class="flex gap-2 items-center">
                        <input
                            class="input p-1 w-44"
                            value={usersIndex[runas_user_id] || ""}
                            type="text"
                            disabled
                            placeholder="John"
                        />

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
                    <!-- svelte-ignore a11y_label_has_associated_control -->
                    <label class="label">
                        <span>Script Code</span>
                        <div class="p-1 rounded border bg-white max-h-screen">
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

                <!-- svelte-ignore a11y_label_has_associated_control -->
                <div class="label">
                    <span>Extrameta</span>
                    <KvEditor
                        data={extrameta ? JSON.parse(extrameta) : {}}
                        onChange={handleExtrametaChange}
                    />
                </div>

                <!-- svelte-ignore a11y_label_has_associated_control -->
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
                    onclick={async () => {
                        loading = true;
                        if (run_as_type === "context_user") {
                            runas_user_id = 0;
                        } else if (run_as_type === "no_user") {
                            runas_user_id = -1
                        }

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
