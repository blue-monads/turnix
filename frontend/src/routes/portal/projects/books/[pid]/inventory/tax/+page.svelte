<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoTable } from "$lib/compo";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let taxes: any[] = $state([]);
    let loading = true;

    const load = async () => {
        loading = true;

        const resp = await api.listTax(pid);
        if (resp.status !== 200) {
            return;
        }

        taxes = resp.data;
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

            <li class="crumb">Taxes</li>
        </ol>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <a
            href={`/z/pages/portal/projects/books/${pid}/inventory/tax/new`}
            class="btn variant-ghost-primary btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />
            Tax
        </a>
    </svelte:fragment>
</AppBar>

<AutoTable
    action_key={"id"}
    key_names={[
        ["id", "ID"],
        ["name", "Name"],
        ["ttype", "Type"],
        ["info", "Info"],
        ["rate", "Rate"],
    ]}
    datas={taxes}
    color={["ttype"]}
    actions={[
        {
            Name: "edit",
            Class: "variant-filled-secondary",
            Action: async (id) => {
                goto(
                    `/z/pages/portal/projects/books/${pid}/inventory/tax/edit?pid=${pid}&tid=${id}`,
                );
            },
        },
        {
            Name: "delete",
            Class: "variant-filled-error",
            Action: async (id) => {
                await api.deleteTax(pid, id);

                load();
            },
        },
    ]}
/>
