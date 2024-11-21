<script lang="ts">
    import type { RootAPI } from "$lib";
    import type { IFrameIPCMessage, RenderBoxHandle } from "./renderbox";

    interface Props {
        pid: string;
        rootApi: RootAPI;
        htmlSource: string;
        title?: string;
        handle?: RenderBoxHandle;
        ipcHandler?:  (ev: MessageEvent) => Promise<void>;
    };


    let {
        pid,
        title,
        rootApi,
        htmlSource,
        handle = $bindable(),
        ipcHandler,
    }: Props = $props();

    let epoch = $state(0);
    let port: MessagePort | undefined;
    let iframe: HTMLIFrameElement | undefined;

    handle = {
        reload: () => {
            console.log("reload");
            epoch = epoch + 1;
            iframe = undefined;
        },
    };

    const onFrameMessage = async (ev: MessageEvent) => {
        const data = ev.data as IFrameIPCMessage;
        console.log("onFrameMessage", data);

        if (ipcHandler) {
            await ipcHandler(ev);
            return;
        }


        if (data.type === "sql_query") {
            console.log("sql_query", data);

            const args = data.data["args"];
            const sqlCode = data.data["sql"];

            try {
                const resp = await rootApi.runProjectSQL2(pid, {
                    qstr: sqlCode,                   
                    data: args,
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

<div class="card p-2 h-full w-full">
    {#key epoch}
        <iframe
            onload={(ev) => {
                try {
                    let chan = new MessageChannel();
                    chan.port1.onmessage = onFrameMessage;
                    port = chan.port1;

                    console.log(
                        "chan.port2 type:",
                        chan.port2 instanceof MessagePort,
                    );

                    iframe?.contentWindow?.postMessage("transfer_port", "*", [
                        chan.port2,
                    ]);
                } catch (error) {
                    console.error("Error in postMessage:", error);
                }
            }}
            bind:this={iframe}
            title={title || ""}
            srcdoc={htmlSource}
            width="100%"
            height="100%"
            class="border-green-200 w-full h-full transition-all"
            allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write;
        document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations;
        legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment;
        picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock;
        xr-spatial-tracking"
            sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox
        allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"
        >
        </iframe>
    {/key}
</div>
