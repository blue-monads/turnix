


<script lang="ts">
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { AutoTable } from "$lib/compo";
    import Loader from "$lib/compo/loader/loader.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
    let loading = $state(true);
    let datas: any[] = $state([]);

    const load = async () => {
        loading = true;
        const resp = await api.listReportTemplate(pid);
        if (resp.status !== 200) {
            return;
        }

        datas = resp.data;
        loading = false;
    };

    load();
</script>



{#if loading}
    <Loader />
{:else}
    <AutoTable
        action_key={"id"}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["report_type", "Report Type"],
            ["created_at", "Created At"],
        ]}
        hashSeed={35}
        datas={datas}
        color={["report_type"]}
        actions={[
            {
                Name: "explore",
                Class: "variant-filled-primary",
                Action: async (id, data) => {

                },
            },
            {
                Name: "edit",
                Class: "variant-filled-primary",
                Action: async (id, data) => {
                    if (data.report_type === "html_report") {
                        goto(
                            `/z/pages/portal/projects/books/${pid}/report/template/edit/html?template_id=${id}`,
                        );
                    } else {
                        goto(
                            `/z/pages/portal/projects/books/${pid}/report/template/edit/sql?template_id=${id}`,
                        );
                    }
                },
            },
            {
                Name: "delete",
                Class: "variant-filled-error",
                Action: async (id, data) => {
                    const res = await api.deleteReportTemplate(pid, id);
                    if (res.status !== 200) {
                        return;
                    }

                    load();
                },
            },
        ]}
    />
{/if}
