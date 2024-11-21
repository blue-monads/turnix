<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import RenderBox from "../../../../../../../playground/RenderBox/RenderBox.svelte";
    import type { RenderBoxHandle } from "../../../../../../../playground/RenderBox/renderbox";
    import { params } from "$lib/params";
    import Loader from "$lib/compo/loader/loader.svelte";

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

        runningSavedHtmlCode = resp.data.template;
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
                <li>Reports</li>
                <li class="crumb-separator" aria-hidden="true">&rsaquo;</li>
                <li>Editor</li>
            </ol>
        </svelte:fragment>

        <svelte:fragment slot="trail">
            <button
                class="btn variant-filled btn-sm"
                onclick={() => {
                    handle?.reload();
                }}
            >
                <SvgIcon className="w-4 h-4" name="play" />
                Reload
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

        </svelte:fragment>
    </AppBar>

    <div
        class="border border-slate-50 h-full w-full p-2 max-h-[90vh]"
    >
        <RenderBox
            {pid}
            htmlSource={runningSavedHtmlCode}
            {rootApi}
            bind:handle
            title="Report Preview"
        />
    </div>
{/if}
