<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import CodeMirror from "svelte-codemirror-editor";
    import { html } from "@codemirror/lang-html";
    import {
        javascript,
        localCompletionSource,
        scopeCompletionSource,
    } from "@codemirror/lang-javascript";
    import { autocompletion } from "@codemirror/autocomplete";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";
    import EasyGrid from "../../playground/EasyGrid/EasyGrid.svelte";

    let tabSet: number = $state(0);
    let savedClientCode = $state("");
    let serverCode = $state("");

    let { runAction = $bindable() } = $props();

    runAction = () => {
        savedClientCode = clientCode;
    };

    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);
</script>

<div class="flex flex-col w-full h-[94vh] p-2">
    <div class="w-full border overflow-auto resize-y p-2 h-24 bg-white rounded">
        <CodeMirror
            bind:value={serverCode}
            extensions={[]}
            lang={javascript()}
        />
    </div>

    <div class="w-full border  p-2">
        <div class="flex-1 w-full">
            <TabGroup>
                <Tab bind:group={tabSet} name="SQL Query" value={0}>
                    <span>Grid</span>
                </Tab>
                <Tab bind:group={tabSet} name="UI" value={1}>
                    <span>Mappings</span>
                </Tab>

                <div class="flex-1" />

                <div class="flex justify-end p-2">
                    <button
                        class="inline-flex rounded p.05 variant-filled self-center"
                        onclick={() => {
                            runAction();
                        }}
                    >
                        <SvgIcon className="w-4 h-4 mt-1" name="play" />
                        Run
                    </button>
                </div>

                <svelte:fragment slot="panel">
                    <div class="max-h-[40vh] md:max-h-[90vh] overflow-auto">
                        {#if tabSet === 0}
                            <EasyGrid
                                columns={[]}
                                enablePagination={false}
                                enableSidebar={false}
                                enableSort={false}
                                onLoad={(params) => {
                                    return [];
                                }}
                            />
                        {:else if tabSet === 1}
                            <span>Mappings HERE</span>
                        {/if}
                        
                    </div>
                </svelte:fragment>
            </TabGroup>

            <!-- <div>AAA</div> -->
        </div>
    </div>
</div>
