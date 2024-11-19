<script lang="ts">
        import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import VerySimpleModal from "../../playground/VerySimpleModal.svelte";
    import FilePicker from "../filestore/FilePicker/FilePicker.svelte";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib";

    const store = getModalStore();
    const api = getContext("__api__") as RootAPI;


    let showModal = $state(false);

</script>

<div class="flex flex-col p-4">
    <h1>Hello</h1>

    <!-- FilePicker -->
  
</div>


<button 
    class="btn btn-sm variant-filled" 
    onclick={() => {
        showModal = !showModal;

    }}
    >
    toggle
</button>


<VerySimpleModal bind:showModal={showModal}>
    <div class="w-full min-w-[724px] h-full">
        <FilePicker pid={1} />
    </div>

</VerySimpleModal>