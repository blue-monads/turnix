<script lang="ts">
    import { page } from "$app/stores";
    import ListProject from "$lib/container/project/listProject.svelte";
    const ptype = $page.params["ptype"];
</script>

<ListProject ptype={ptype} />
  