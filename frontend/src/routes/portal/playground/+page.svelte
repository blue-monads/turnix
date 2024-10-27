<script lang="ts">
  import AutoForm from "$lib/compo/autoform/auto_form.svelte";
  import { getModalStore } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import EasyGrid from "./EasyGrid/EasyGrid.svelte";
  import sampleData from "./EasyGrid/sample";

  const store = getModalStore();

  const items = [
    { name: "rename", icon: "pencil-square" },
    { name: "download", icon: "arrow-down-on-square" },
    { name: "delete", icon: "trash" },
  ];
</script>

<div class="flex flex-col p-4">
  <EasyGrid
    enableSort={true}
    enableSidebar={true}
    columns={[
      { title: "Name", key: "name" },
      { title: "Info", key: "info" },
      { title: "Type", key: "acc_type", rendererOptions: { autoColor: true } },
      { title: "Created At", key: "created_at" },
    ]}
    onLoad={(params) => {
      console.log("@load", params);

      return sampleData;
    }}
  />
</div>
