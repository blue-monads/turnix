<script lang="ts">
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { AutoForm } from "$lib/compo";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);
</script>

<AutoForm
    data={{}}
    schema={{
        name: "Add Report",
        fields: [
            {
                name: "Name",
                ftype: "TEXT",
                key_name: "name",                
            },
            {
                name: "Type",
                ftype: "SELECT",
                key_name: "report_type",
                options: ["html_report", "sql_report"],
            },
        ],
        required_fields: ["name"],
    }}
    onSave={async (data) => {

        const res = await api.addReportTemplate(pid, data);
        if (res.status !== 200) {
            return
        }

        goto(`/z/pages/portal/projects/books/${pid}/report`);

    }}
/>
