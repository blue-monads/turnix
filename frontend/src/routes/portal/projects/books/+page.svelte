<script lang="ts">
  import ListProject from "../../listProject.svelte";
  import { params } from "$lib/params";
  import ListAccount from "./account/+page.svelte";

  let pid = $params["pid"];


</script>

{#if pid}
  <ListAccount />
{:else}
  <ListProject ptype="books" />
{/if}
