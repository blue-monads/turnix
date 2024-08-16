<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";

    import CodeMirror from "svelte-codemirror-editor";
    import { sql, SQLite } from "@codemirror/lang-sql";
    import { html } from "@codemirror/lang-html";

    import { page } from "$app/stores";

    import * as sampleCode from "./sampleCode";
    import { samplePreview } from "./samplePreview";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    let iframe: HTMLIFrameElement;

    let tabSet: number = 0;
    let previewTabSet: number = 0;
    let sqlCode = sampleCode.sqlCode;
    let htmlCode = samplePreview;
    let savedHtmlCode = "";

    const pid = $page.params["pid"];

    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);

    interface Message {
        type: "sql_query" | "api_call" | "ping";
        name?: string;
        data: any;
        msgId: number;
    }

    let port: MessagePort;

    const onFrameMessage = async (ev: MessageEvent) => {
        const data = ev.data as Message;
        console.log("onFrameMessage", data);

        if (data.type === "sql_query") {
            console.log("sql_query", data);

            try {
                const resp = await rootApi.runProjectSQL(pid, {
                    input: sqlCode,
                    name: data.name as string,
                    data: data.data,
                });

                if (resp.status !== 200) {
                    port?.postMessage({
                        msgId: data.msgId,
                        data: { msg: "error", data: resp.data },
                    });
                    return;
                }

                port?.postMessage({
                    msgId: data.msgId,
                    data: { msg: "success", data: resp.data },
                });
            } catch (error) {}
        } else if (data.type === "api_call") {
            console.log("api_call", data);
        } else if (data.type === "ping") {
            port?.postMessage({
                msgId: data.msgId,
                data: { msg: "pong" },
            });
        }
    };
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Reports</li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Editor</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <button
            class="btn variant-filled btn-sm"
            on:click={() => {
                savedHtmlCode = htmlCode;
            }}
        >
            <SvgIcon className="w-4 h-4" name="play" />
            Run
        </button>

        <button
            class="btn variant-filled-secondary btn-sm"
            on:click={() => {
                iframe.contentWindow?.print();
            }}
        >
            <SvgIcon className="w-4 h-4" name="printer" />
            Print
        </button>
    </svelte:fragment>
</AppBar>

<div class=" flex flex-col md:flex-row w-full h-[94vh]">
    <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full">
        <TabGroup>
            <Tab bind:group={tabSet} name="SQL Query" value={0}>
                <span>SQL Query</span>
            </Tab>
            <Tab bind:group={tabSet} name="UI" value={1}>
                <span>UI</span>
            </Tab>

            <svelte:fragment slot="panel">
                <div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"> 
                    {#if tabSet === 0}
                        <CodeMirror
                            bind:value={sqlCode}
                            lang={sql({
                                dialect: SQLite,
                            })}
                        />
                    {:else if tabSet === 1}
                        <CodeMirror bind:value={htmlCode} lang={html()} />
                    {/if}
                </div>
            </svelte:fragment>
        </TabGroup>
    </div>

    <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2">
        <div class="card p-2 h-full w-full">
            <iframe
                on:load={(ev) => {
                    try {
                        let chan = new MessageChannel();
                        chan.port1.onmessage = onFrameMessage;
                        port = chan.port1;

                        console.log(
                            "chan.port2 type:",
                            chan.port2 instanceof MessagePort,
                        );

                        iframe?.contentWindow?.postMessage(
                            "transfer_port",
                            "*",
                            [chan.port2],
                        );
                    } catch (error) {
                        console.error("Error in postMessage:", error);
                    }
                }}
                bind:this={iframe}
                title="preview"
                srcdoc={savedHtmlCode}
                width="100%"
                height="100%"
                class="border-green-200 w-full h-full transition-all"
                allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"
                sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"
            >
            </iframe>
        </div>
    </div>
</div>
