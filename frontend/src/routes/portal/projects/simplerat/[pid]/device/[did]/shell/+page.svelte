<script lang="ts">
    import { AutoTable } from "$lib/compo";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { NewSimpleRATApi } from "../../../../lib/SimpleRATApi";
    import { Terminal } from "@xterm/xterm";

    const pid = $page.params["pid"];
    const did = $page.params["did"];

    const api = NewSimpleRATApi(getContext("__api__") as RootAPI);
    const terminal = new Terminal();
    let termRef: HTMLElement | null = null;

    const load = async () => {
        let url = api.getNewRoomUrl(pid, did);

        if (url.startsWith("https")) {
            url = url.replace("https", "wss");
        }

        if (url.startsWith("http")) {
            url = url.replace("http", "ws");
        }

        const parsedURL = new URL(url);
        parsedURL.searchParams.append("mtype", "service.shell")



        const ws = new WebSocket(parsedURL.toString());

        ws.onopen = () => {
            console.log("Connected");
            terminal.open(termRef as HTMLElement);
            terminal.onData((data) => {
                console.log("Data", data);
                ws.send(data);
            });
        };


        ws.onmessage = (e) => {
            console.log("Message", e.data);
            terminal.write(e.data);
        };

        ws.onclose = () => {
            console.log("Closed");
        };

        ws.onerror = (e) => {
            console.log("Error", e);
        };
    };

    load();
</script>

<svelte:head>
    <title>Terminal Shell</title>
    <link rel="stylesheet" href="https://xtermjs.org/css/xterm.css" />
</svelte:head>

<AppBar>
    <svelte:fragment slot="lead">
        <h3 class="shell">Terminal Shell</h3>
    </svelte:fragment>
</AppBar>

<div class="shell p-4 rounded">

    <div bind:this={termRef} id="terminal"></div>
</div>