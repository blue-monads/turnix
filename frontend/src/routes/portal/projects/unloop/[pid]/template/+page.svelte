<script lang="ts">
  import { PageLayout, AutoTable } from "$lib/compo";
  import { getContext, onMount } from "svelte";
  import * as nav from "$lib/projects/unloop/route";

  import { params } from "$lib/params";
  import { NewUnloopAPI } from "$lib/projects/unloop";
  import type { RootAPI } from "$lib/api";

  let pid = $params["pid"];

  const api = NewUnloopAPI(getContext("__api__") as RootAPI);

  let datas: any[] = [];

  const load = async () => {
    const resp = await api.listTemplates(pid);
    if (resp.status !== 200) {
      return;
    }
    datas = resp.data;
  };

  load();
</script>

<PageLayout
  title="OnLoop Templates"
  actions={[
    { name: "add", actionFn: () => nav.gotoProjectOnloopTemplateAdd(pid) },
  ]}
>
  <AutoTable
    action_key={"id"}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["ttype", "Template type"],
    ]}
    {datas}
    color={["ttype"]}
    actions={[
      {
        Name: "watch",
        Class: "bg-yellow-400",
        Action: async (id) => {
          nav.gotoProjectOnloopTemplateWatch(pid, id);
        },
      },

      {
        Name: "explore",
        Class: "bg-green-400",
        Action: async (id) => {
          nav.gotoProjectOnloopQueues(pid, id);
        },
      },

      {
        Name: "push",
        Class: "bg-orange-400",
        Action: async (id) => {
          nav.gotoProjectOnloopQueueAdd(pid, id);
        },
      },

      {
        Name: "edit",
        Action: async (id) => {
          nav.gotoProjectOnloopTemplateEdit(pid, id);
        },
      },
      {
        Name: "delete",
        Class: "bg-red-400",
        Action: async (id) => {
          await api.removeTemplate(pid, id);

          load();
        },
      },
    ]}
  />
</PageLayout>
