<script lang="ts">
  import { AutoForm, PageLayout, AutoTable, Loader } from "$lib/compo";

  import { gotoProjectOnloopTemplates } from "$lib/projects/unloop/route";
  import { params } from "$lib/params";
  import type { RootAPI } from "$lib/api";
  import { getContext } from "svelte";
  import { NewUnloopAPI } from "$lib/projects/unloop";
  import { page } from "$app/stores";

  const api = NewUnloopAPI(getContext("__api__") as RootAPI);

  let pid = $page.params["pid"];
  let tid = $params["tid"];


  let message = $state("");
  let loading = $state(true);
  let data = $state({});

  const load = async () => {
    loading = true;
    const resp = await api.getTemplate(pid, tid);
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
        { ftype: "TEXT", key_name: "name", name: "Name" },
        {
          ftype: "KEY_VALUE_TEXT",
          key_name: "contents",
          name: "Contents",
        },
        {
          ftype: "KEY_VALUE_TEXT",
          key_name: "result",
          name: "Result",
        },
        {
          ftype: "LONG_TEXT",
          key_name: "contentScript",
          name: "Content script",
        },
      ],
      name: "Edit Template",
      required_fields: ["name"],
    }}
    onSave={async (newdata) => {
      const resp = await api.updateTemplate(pid, tid, newdata);
      if (resp.status !== 200) {
        message = resp.data;
        return;
      }

      gotoProjectOnloopTemplates(pid);
    }}
  />
{/if}
