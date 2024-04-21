<script lang="ts">
  import { page } from "$app/stores";

  const pid = $page.params["pid"];
  console.log("@pid", pid);
</script>

<div>Books {pid}</div>
