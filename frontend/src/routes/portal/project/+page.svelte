<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import {
        AppBar,
        getModalStore,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { FloatyButton } from "$lib/compo";
    import ListProject from "$lib/container/project/listProject.svelte";

    const api = getContext("__api__") as RootAPI;


    const store = getModalStore();


</script>

<svelte:head>
  <title>Projects</title>
</svelte:head>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4">Projects</h4>
  </svelte:fragment>
</AppBar>


<ListProject />

<FloatyButton
    handler={() => {
        store.trigger({
            type: "component",
            component: "project_picker",
            meta: { api },
        });
    }}
/>
