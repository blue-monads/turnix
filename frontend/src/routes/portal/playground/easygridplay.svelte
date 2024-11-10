<script lang="ts">
    import AutoForm from "$lib/compo/autoform/auto_form.svelte";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import EasyGrid from "./EasyGrid/EasyGrid.svelte";
    import type { GridHandle } from "./EasyGrid/easyTypes";
    import sampleData from "./EasyGrid/sample";
  
    const store = getModalStore();
  
    const items = [
      { name: "rename", icon: "pencil-square" },
      { name: "download", icon: "arrow-down-on-square" },
      { name: "delete", icon: "trash" },
    ];
    let handle: GridHandle | undefined = $state();
  </script>
  
  <div class="flex flex-col p-4">
    <EasyGrid
      key="id"
      enableSort={true}
      enableSidebar={true}
      bind:handle
      columns={[
        { title: "Name", key: "name" },
        { title: "Info", key: "info" },
        { title: "Type", key: "acc_type", rendererOptions: { autoColor: true } },
        { title: "Created At", key: "created_at" },
      ]}
      onLoad={async (params) => {
        console.log("@load", params);
  
        return new Promise((resolve) => {
          setTimeout(() => {
            resolve(sampleData);
          }, 2000);
        });
  
      }}
    />
  </div>
  
  
  <button class="btn btn-sm variant-filled" onclick={handle?.reload}>Reload</button>