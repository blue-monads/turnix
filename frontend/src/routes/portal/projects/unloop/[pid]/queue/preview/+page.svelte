<script lang="ts">
  import { Loader, AutoForm } from "$lib/compo";

  import { params } from "$lib/params";
  import type { RootAPI } from "$lib/api";
  import { getContext } from "svelte";
  import { NewUnloopAPI } from "$lib/projects/unloop";
  import { page } from "$app/stores";


  let pid = $page.params["pid"];
  let mid = $params["mid"];
  let message = "";
  let data = {};
  let loading = true;

  const api = NewUnloopAPI(getContext("__api__") as RootAPI);

  const load = async () => {
    const resp = await api.getQueueMessage(pid, mid);
    if (resp.status !== 200) {
      return;
    }
    data = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <Loader />
{:else}
  <AutoForm
    {data}
    {message}
    schema={{
      fields: [
        {
          ftype: "KEY_VALUE_TEXT",
          key_name: "contents",
          name: "Contents",
          disabled: true,
        },
        {
          ftype: "SELECT",
          key_name: "status",
          name: "Status",
          options: ["submited", "processed", "denied"],
          disabled: true,
        },

        {
          ftype: "KEY_VALUE_TEXT",
          key_name: "extrameta",
          name: "Extra Meta",
          disabled: true,
        },
      ],
      name: "Preview Queue message",
      required_fields: [],
    }}
    onSave={async (newdata) => {}}
  />
{/if}
