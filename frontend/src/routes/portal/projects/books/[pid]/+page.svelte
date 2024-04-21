<script lang="ts">
  import { page } from "$app/stores";
  import AccountList from "./account/+page.svelte";

  const pid = $page.params["pid"];
  console.log("@pid", pid);
</script>

<AccountList />
