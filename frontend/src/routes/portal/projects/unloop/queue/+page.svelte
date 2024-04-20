<script lang="ts">
  import { AutoForm, PageLayout, AutoTable } from "$lib/compo";
  import { getContext, onMount } from "svelte";
  import * as nav from "$lib/modules/unloop/route";
  import type { RootAPI } from "$lib/api";
  import { params } from "$lib/params";
  import { NewUnloopAPI } from "$lib/modules/unloop";

  let pid = $params["pid"];

  const api = NewUnloopAPI(getContext("__api__") as RootAPI);

  let datas: any[] = [];

  const load = async () => {
    const resp = await api.listQueueMessages(pid);
    if (resp.status !== 200) {
      return;
    }
    datas = resp.data;
  };

  load();
</script>

<PageLayout
  title="Onloop Queue Messages"
  actions={[
    {
      name: "add",
      icon: "plus",
      actionFn: () => {
        nav.gotoProjectOnloopQueueAdd(pid, $params["tid"]);
      },
    },
  ]}
>
  <AutoTable
    action_key={"id"}
    key_names={[
      ["id", "ID"],

      ["submitter", "Submitter"],
      ["status", "Status"],
      ["createdAt", "Created At"],
    ]}
    {datas}
    color={["ttype"]}
    actions={[
      {
        Name: "preview",
        Class: "bg-green-400",
        Action: async (id, data) => {
          nav.gotoProjectOnloopQueuePreview(pid, $params["tid"], id);
        },
      },
      {
        Name: "delete",
        Class: "bg-red-400",
        Action: async (id) => {
          await api.removeUpdateQueueMessage(pid, id);

          load();
        },
      },
    ]}
  />
</PageLayout>
