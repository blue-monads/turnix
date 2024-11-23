<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { AutoTable } from "$lib/compo";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
    let loading = $state(true);
    let datas: any[] = $state([]);

    const load = async () => {
        loading = true;
        const resp = await api.listReportSaved(pid);
        if (resp.status !== 200) {
            return;
        }

        datas = resp.data;
        loading = false;
    };


    load();


</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>    
                <a class="anchor" href={`/z/pages/portal/projects/books/${pid}/report`}>Reports</a>    
              </li>

            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Saved</li>
        </ol>
    </svelte:fragment>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <AutoTable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["created_at", "Created At"],
        ]}
        hashSeed={35}
        datas={datas}
        color={["report_type"]}
        actions={[
            {
                Name: "view",
                Class: "variant-filled-primary",
                Action: async (id, data) => {

                },
            },
            {
                Name: "delete",
                Class: "variant-filled-error",
                Action: async (id) => {
                    const res = await api.deleteReportSaved(pid, id);
                    if (res.status !== 200) {
                        return;
                    }

                    load();
                },
            },
        ]}
    />
{/if}
