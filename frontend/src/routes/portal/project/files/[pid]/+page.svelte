<script lang="ts">
    import { RootAPI } from "$lib/api";
    import { getContext } from "svelte";
    import {
        AppBar,
        FileDropzone,
        getModalStore,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { page } from "$app/stores";
    import { Loader } from "$lib/compo";

    const pid = $page.params["pid"];


    const api = getContext("__api__") as RootAPI;

    let projects = [];

    const load = async () => {
        const resp = await api.listProjectFiles(pid);
        if (resp.status !== 200) {
            return;
        }

        projects = resp.data;
    };
    const store = getModalStore();

    load();

    const onDrop = async (e: any) => {
        console.log("onDrop", e);

      await  api.addProjectFile(pid, "testxyz", e.target.files[0]);
    };


</script>

<svelte:head>
  <title>Project Files</title>
</svelte:head>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4">Project Files</h4>
  </svelte:fragment>
</AppBar>

<FileDropzone 
on:change={onDrop}
name="files"


>
	<svelte:fragment slot="lead">(icon)</svelte:fragment>
	<svelte:fragment slot="message">Upload files to your project</svelte:fragment>
</FileDropzone>



