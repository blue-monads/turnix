<script lang="ts">
  import AutoForm from "$lib/compo/autoform/auto_form.svelte";
</script>

<AutoForm
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
    required_fields: [],
  }}
/>
