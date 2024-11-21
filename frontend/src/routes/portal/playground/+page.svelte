<script lang="ts">
        import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import VerySimpleModal from "../../playground/VerySimpleModal/VerySimpleModal.svelte";
    import FilePicker from "../filestore/FilePicker/FilePicker.svelte";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib";
    import RenderBox from "./RenderBox/RenderBox.svelte";

    const store = getModalStore();
    const api = getContext("__api__") as RootAPI;


    let showModal = $state(false);

</script>

<button 
    class="btn btn-sm variant-filled" 
    onclick={() => {
        showModal = !showModal;

    }}
    >
    toggle
</button>



<RenderBox 
    pid={"1"} 
    rootApi={api} 
    htmlSource={"<h1>RenderBox</h1>"}  
    title={"RenderBox TEST"}

    />


<VerySimpleModal bind:showModal={showModal}>
    <div class="w-full md:min-w-[724px] h-full">
        <FilePicker pid={1} tabMode="project" />
    </div>

</VerySimpleModal>