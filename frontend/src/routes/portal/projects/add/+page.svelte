<script lang="ts">
    import AutoForm from "../../../../compo/autoform/auto_form.svelte";
</script>

<AutoForm
    data={{}}
    schema={{
        fields: [
            { ftype: "TEXT", key_name: "name", name: "Name" },
            {
                ftype: "SELECT",
                key_name: "ptype",
                name: "type",
                options: ["laction"],
            },
        ],
        name: "Add Project",
        required_fields: ["name"],
    }}
    onSave={async (newdata) => {}}
/>
