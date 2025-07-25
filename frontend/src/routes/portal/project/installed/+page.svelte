<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import {
        AppBar,
        getModalStore,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";

    const api = getContext("__api__") as RootAPI;


    const store = getModalStore();


</script>

<svelte:head>
  <title>Installed Project types</title>
</svelte:head>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4">Installed Project types</h4>
  </svelte:fragment>
</AppBar>

