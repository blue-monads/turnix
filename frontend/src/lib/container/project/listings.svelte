<script lang="ts">
  import { AutoTable, PageLayout } from "$lib/compo";

  import { getContext, onMount } from "svelte";
  import * as nav from "$lib/nav";
  import type { RootAPI } from "$lib/api";
  interface Props {
    ptype?: string;
    title?: string;
  }

  let { ptype = "", title = "projects" }: Props = $props();

  const api = getContext("__api__") as RootAPI;

  let datas: any[] = $state([]);

  const load = async () => {
    const resp = await api.listProjects(ptype);
    if (resp.status !== 200) {
      return;
    }
    datas = resp.data;
  };

  load();
</script>

<PageLayout {title} actions={[{ name: "add", actionFn: nav.gotoAddProject }]}>
  <AutoTable
    action_key={"id"}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["ptype", "Project type"],
      ["owned_by", "Owner"],
    ]}
    {datas}
    color={["ptype"]}
    actions={[
      {
        Name: "explore",
        Class: "bg-green-400",
        Action: async (id, data) => {},
      },

      {
        Name: "edit",

        Action: async (id) => {
          nav.gotoEditProject(id);
        },
      },

      {
        Name: "hooks",
        Class: "bg-yellow-400",
        Action: async (id, data) => {
          nav.gotoProjectHooks( data["ptype"],   id, );
        },
      },
      {
        Name: "delete",
        Class: "bg-red-400",
        Action: async (id) => {
          await api.removeProject(id);

          load();
        },
      },
    ]}
  />
</PageLayout>
