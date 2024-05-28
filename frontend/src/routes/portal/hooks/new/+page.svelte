<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import { params } from "$lib/params";
    import { AppBar } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";
    import HookForm from "../hookForm.svelte";
    import { gotoPorjects } from "$lib/nav";

    const pid = $params["pid"];

    const api = getContext("__api__") as RootAPI;



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
            <li>New</li>
        </ol>
    </svelte:fragment>
</AppBar>

<HookForm onSave={ async (data) => {

    const resp = await api.addProjectHook(pid, data)
    if (resp.status !== 200) {
        return
    }

    gotoPorjects()

}}  />