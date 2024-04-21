<script lang="ts">

  import AccountList from "./account/+page.svelte";
  import { page } from "$app/stores";

  const pid = $page.params["pid"];
  console.log("@pid", pid);
</script>

<AccountList />
