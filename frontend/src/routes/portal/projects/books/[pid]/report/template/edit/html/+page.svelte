<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    import CodeMirror from "svelte-codemirror-editor";
    import { html } from "@codemirror/lang-html";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import RenderBox from "../../../../../../../playground/RenderBox/RenderBox.svelte";
    import type { RenderBoxHandle } from "../../../../../../../playground/RenderBox/renderbox";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";

    let savedHtmlCode = $state("");
    let runningSavedHtmlCode = $state("");

    const pid = $page.params["pid"];
    const template_id = $params["template_id"];

    const rootApi = getContext("__api__") as RootAPI;
    const api = NewBookAPI(rootApi);
    let loading = $state(true);

    const load = async () => {
        loading = true;

        const resp = await api.getReportTemplate(pid, template_id);
        if (resp.status !== 200) {
            return;
        }

        savedHtmlCode = resp.data.template;
        loading = false;
    };

    load();

    let handle: RenderBoxHandle | undefined;
</script>

{#if loading}
    <Loader />
{:else}
    <AppBar>
        <svelte:fragment slot="lead">
            <ol class="breadcrumb">
                <li class="crumb">
                    <a class="anchor" href="/z/pages/portal/projects/books"
                        >Books</a
                    >
                </li>
                <li class="crumb-separator" aria-hidden="true">&rsaquo;</li>
                <li>
                    <a
                        class="anchor"
                        href={`/z/pages/portal/projects/books/${pid}/report`}
                        >Reports</a
                    >
                </li>

                <li class="crumb-separator" aria-hidden="true">&rsaquo;</li>
                <li>Editor</li>
            </ol>
        </svelte:fragment>

        <svelte:fragment slot="trail">
            <button
                class="btn variant-filled btn-sm"
                onclick={() => {
                    runningSavedHtmlCode = savedHtmlCode;
                    handle?.reload();
                }}
            >
                <SvgIcon className="w-4 h-4" name="play" />
                Run
            </button>

            <button
                class="btn variant-filled-secondary btn-sm"
                onclick={() => {
                    handle?.print();
                }}
            >
                <SvgIcon className="w-4 h-4" name="printer" />
                Print
            </button>

            <button
                class="btn variant-filled-tertiary btn-sm"
                onclick={async () => {
                    loading = true;
                    await api.updateReportTemplate(pid, template_id, {
                        template: savedHtmlCode,
                    });

                    loading = false;
                }}
            >
                <SvgIcon className="w-4 h-4" name="arrow-down-on-square" />
                Save
            </button>
        </svelte:fragment>
    </AppBar>

    <div class=" flex flex-col md:flex-row w-full h-[94vh]">
        <div
            class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full overflow-auto"
        >
            <CodeMirror bind:value={savedHtmlCode} lang={html()} />
        </div>

        <div
            class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"
        >
            <RenderBox
                {pid}
                htmlSource={runningSavedHtmlCode}
                {rootApi}
                bind:handle
            />
        </div>
    </div>
{/if}
