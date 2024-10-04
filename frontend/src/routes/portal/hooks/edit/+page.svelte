<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import { params } from "$lib/params";
    import { AppBar } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";
    import HookForm from "../hookForm.svelte";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $params["pid"];
    const hid = $params["hid"];

    const api = getContext("__api__") as RootAPI;

    let data = $state({});
    let loading = $state(true);

    const load = async () => {
        loading = true;

        const resp = await api.getProjectHook(pid, hid);
        if (resp.status !== 200) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    load();
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books"
                    >Projects</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Hooks</li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Edit</li>
        </ol>
    </svelte:fragment>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <HookForm
        onSave={async (data) => {
            const resp = await api.updateProjectHook(pid, hid, data);
            if (resp.status !== 200) {
                return;
            }

            location.pathname.replace("/edit", "");
        }}
        {...data}
    />
{/if}
