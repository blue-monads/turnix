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
    import { sampleHtml, sampleJs } from "./sample";
    import { params } from "$lib/params";
    import Runner from "../../run/Runner.svelte";


    let tabSet: number = $state(0);
    let serverCode = $state(sampleJs);
    let clientCode = $state(sampleHtml);
    let savedClientCode = $state("");
    
    let {
        runAction = $bindable(),
    } = $props();


    runAction = () => {
        savedClientCode = clientCode;
    }


    const pid = $params["pid"];

    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);

    interface Message {
        type: "api_call" | "ping";
        name?: string;
        data: any;
        msgId: number;
    }


    const  completionObject = {}


    const onFrameMessage = async (ev: MessageEvent) => {
       
    };
</script>


<div class=" flex flex-col md:flex-row w-full h-[94vh]">
    <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full">
        <TabGroup>
            
            <Tab padding="ml-6 p-2" bind:group={tabSet} name="SQL Query" value={0}>
                <span>Server Code</span>
            </Tab>
            <Tab bind:group={tabSet} name="UI" value={1}>
                <span>Client Code</span>
            </Tab>

            <svelte:fragment slot="panel">
                <div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"> 
                    {#if tabSet === 0}
                        <CodeMirror
                            bind:value={serverCode}

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

                            lang={javascript()}
                            
                        />
                    {:else if tabSet === 1}
                        <CodeMirror bind:value={clientCode} lang={html()} />
                    {/if}
                </div>
            </svelte:fragment>
        </TabGroup>
    </div>

    <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2">
        <Runner client_code={savedClientCode} />        
    </div>
</div>
