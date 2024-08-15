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

    import * as sampleCode from "./sampleCode";
    import { samplePreview } from "./samplePreview";

    let iframe: HTMLIFrameElement;

    let tabSet: number = 0;
    let sqlCode = sampleCode.sqlCode;
    let htmlCode = samplePreview;
    let savedHtmlCode = "";

    interface Message {
        type: "sql_query" | "api_call" | "ping";
        name?: string;
        data: any;
        msgId: number
    }

    let port: MessagePort;

   const onFrameMessage = (ev: MessageEvent) => {
        const data = ev.data as Message;
        console.log("onFrameMessage", data);

        if (data.type === "sql_query") {
            console.log("sql_query", data);
        }  else if (data.type === "api_call") {
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
            Run
        </button>
    </svelte:fragment>
</AppBar>

<div class=" flex flex-col md:flex-row w-full h-full max-h-[90vh]">
    <div
        class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full overflow-auto"
    >
        <TabGroup>
            <Tab bind:group={tabSet} name="tab1" value={0}>
                <span>SQL Query</span>
            </Tab>
            <Tab bind:group={tabSet} name="tab2" value={1}>
                <span>UI</span>
            </Tab>

            <svelte:fragment slot="panel">
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
            </svelte:fragment>
        </TabGroup>
    </div>

    <div
        class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"
    >
        <h4 class="h4 p-1 uppercase">Preview</h4>

        <div class="card p-2 h-full w-full">
            <iframe
                on:load={(ev) => {
                    try {

                        let chan = new MessageChannel();
                        chan.port1.onmessage = onFrameMessage;
                        port = chan.port1;

                        console.log("chan.port2 type:", chan.port2 instanceof MessagePort);

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
