<script lang="ts">
    import { createEventDispatcher } from "svelte";

    interface Props {
        client_code?: any;
    }

    let { client_code = `` }: Props = $props();

    
    let iframe: HTMLIFrameElement = $state();

    interface Message {
        type: "plugin_ipc" | "ping";
        name?: string;
        data: any;
        msgId: number;
    }

    let port: MessagePort = $state();

    const onFrameMessage = async (ev: MessageEvent) => {
        const data = ev.data as Message;
        console.log("onFrameMessage", data);

        if (data.type === "plugin_ipc") {
            console.log("sql_query", data);
            try {
                
                
                port?.postMessage({
                    msgId: data.msgId,
                    data: { msg: "success", data: {} },
                });


            } catch (error) {}
        } else if (data.type === "ping") {
            port?.postMessage({
                msgId: data.msgId,
                data: { msg: "pong" },
            });
        }
    };
</script>

<div class="card p-2 h-full w-full">
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
        title="preview"
        srcdoc={client_code}
        width="100%"
        height="100%"
        class="border-green-200 w-full h-full transition-all"
        allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"
        sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"
    >
    </iframe>
</div>
