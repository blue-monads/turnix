<script lang="ts">
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { AutoTable } from "$lib/compo";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
    let loading = $state(true);
    let datas: any[] = $state([]);

    const load = async () => {
        loading = true;
        const resp = await api.listNotepad(pid);
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

            <li class="crumb">Notepad</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <a
            href={`/z/pages/portal/projects/books/${pid}/notepad/new`}
            class="btn variant-ghost-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Note
        </a>
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

        ]}
        datas={datas}
        color={["ctype"]}
        actions={[
            {
                Name: "edit",
                Class: "variant-filled-primary",
                Action: async (id, data) => {

                },
            },
            {
                Name: "delete",
                Class: "variant-filled-error",
                Action: async (id, data) => {
                    const res = await api.deleteNotepad(pid, id);
                    if (res.status !== 200) {
                        return;
                    }

                    load();
                },
            },
        ]}
    />
{/if}
