<script lang="ts">
    import type { ProjectDef, RootAPI } from "$lib/api";
    import KvEditor from "$lib/compo/autoform/_kv_editor.svelte";
    import Loader from "$lib/compo/loader/loader.svelte";
    import { params } from "$lib/params";
    import { AppBar } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";
    import HookForm from "../hookForm.svelte";

    const ptype = $params["ptype"];

    export let codeValue = `const handle = (ctx) => {

    }`;

    let hook_type = "webhook";

    const api = getContext("__api__") as RootAPI;
    let runas_user_id = 0;
    let runas_specific_user = false;

    let data: ProjectDef;
    let loading = true;

    let usersIndex: Record<number, string> = {};

    const load = async () => {
        loading = true;
        const resp = await api.getProjectType(ptype);
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
            <li>New</li>
        </ol>
    </svelte:fragment>
</AppBar>

<HookForm  />