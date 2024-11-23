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
    import LiveEditor from "./LiveEditor.svelte";

    let tabSet: number = $state(0);
    let serverCode = $state(sampleJs);
    let clientCode = $state(sampleHtml);
    let savedClientCode = $state("");

    const pid = $params["pid"];    

    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);

    let runAction = $state();

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

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/project">Projects</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>
                <a class="anchor" href={`/z/pages/portal/project/plugins?pid=${pid}`}>
                    Plugins
                </a>
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Editor</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <button
            class="btn variant-filled btn-sm"
            onclick={runAction}
        >
            <SvgIcon className="w-4 h-4" name="play" />
            Run
        </button>

    </svelte:fragment>
</AppBar>

<LiveEditor bind:runAction={runAction} />

