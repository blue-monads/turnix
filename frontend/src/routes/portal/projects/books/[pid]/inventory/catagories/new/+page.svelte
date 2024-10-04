<script lang="ts">
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { AutoForm } from "$lib/compo";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let message = $state("");

    const submit = async (data: Record<string, any>) => {
        const resp = await api.addCatagory(pid, {
            name: data.name,
            info: data.info,
        });
        if (resp.status !== 200) {
            message = resp.data.message;
            return;
        }

        goto(`/z/pages/portal/projects/books/${pid}/inventory/catagories`);
    };
</script>

<AutoForm
    {message}
    schema={{
        fields: [
            {
                name: "Name",
                ftype: "TEXT",
                key_name: "name",
            },
            {
                name: "Info",
                ftype: "LONG_TEXT",
                key_name: "info",

            },
        ],
        name: "Add Catagory",
        required_fields: ["name"],
    }}
    onSave={submit}
/>
