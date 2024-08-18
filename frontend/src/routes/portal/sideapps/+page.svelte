<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { AppBar } from "@skeletonlabs/skeleton";
    import { page } from "$app/stores";
    import { params } from "$lib/params";

    const pid = $params["pid"];
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Side Apps</a>
                
            </li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <a
            href={`/z/pages/portal/sideapps/editor?pid=${pid}`}
            class="btn variant-filled-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="pencil-square" />


            Editor
        </a>
    </svelte:fragment>
</AppBar>
