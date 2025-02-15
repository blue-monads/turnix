<script lang="ts">
    
    import { page } from "$app/stores";
    import { NewZblogAPI } from "../../../zblogApi";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib";
    import { goto } from "$app/navigation";
    import { AppBar } from "@skeletonlabs/skeleton";
    import AutoForm from "$lib/compo/autoform/auto_form.svelte";
    import { base } from "$app/paths";

    const pid = $page.params["pid"];

    const api = NewZblogAPI(getContext("__api__") as RootAPI);


</script>


<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/project/list/zblog">ZBlogs</a>
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Sites</li>
        </ol>
    </svelte:fragment>
</AppBar>




<AutoForm
    data={{
        provider: "github_pages",
    }}
    schema={{
        name: "Add Site",
        fields: [
            {
                name: "Name",
                ftype: "TEXT",
                key_name: "name",
            },
            {
                name: "API Key",
                ftype: "TEXT",
                key_name: "api_key",
            },
            {
                name: "Provider",
                ftype: "TEXT",
                key_name: "provider",
                options: ["github_pages", "custom"],
            },
            {
                name: "Domain",
                ftype: "TEXT",
                key_name: "domain",
            },
            {
                name: "Base Path",
                ftype: "TEXT",
                key_name: "base_path",
            },
            {
                name: "Deploy Repo",
                ftype: "TEXT",
                key_name: "deploy_repo",
            }
        ],
        required_fields: ["name", "provider"],
    }}
    onSave={async (data) => {

        if ( !data.provider) {
            data.provider = "github_pages";
        }

        const res = await api.addSite(pid, data);
        if (res.status !== 200) {
            return
        }

        goto(`/z/pages/portal/projects/zblog/${pid}/sites`);

    }}
/>