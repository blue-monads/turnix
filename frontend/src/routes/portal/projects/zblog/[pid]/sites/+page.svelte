<script lang="ts">
    import { goto } from "$app/navigation";
    import { AutoTable, FloatyButton } from "$lib/compo";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { NewZblogAPI } from "../../zblogApi";
  
    const pid = $page.params["pid"];
  
    const store = getModalStore();
    const api = NewZblogAPI(getContext("__api__") as RootAPI);
  
    let sites: any[] = $state([]);
  
    const load = async () => {
      const resp = await api.listSite(pid);
      if (resp.status !== 200) {
        return;
      }
  
      sites = resp.data;
    };
  
    load();
  </script>
  
  <AppBar>
    <svelte:fragment slot="lead">
      <ol class="breadcrumb">
        <li class="crumb">
          <a class="anchor" href="/z/pages/portal/project/list/zblog">ZBlogs</a>
        </li>
        <li class="crumb-separator" aria-hidden>&rsaquo;</li>
        <li>Sites</li>
      </ol>
    </svelte:fragment>
  
    <div class="flex flex-wrap justify-end gap-2">
      <a
        href={`/z/pages/portal/projects/zblog/${pid}/sites/new`}
        class="btn variant-filled btn-sm"
      >
        <SvgIcon className="h-4 w-4" name="plus" />
  
        Site
      </a>
  
  
    </div>
  </AppBar>
  
  <AutoTable
    action_key={"id"}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["provider", "Provider"],
      ["domain", "Domain"],
      ["base_path", "Base Path"],
      ["created_at", "Created At"],

    ]}
    datas={sites}

    actions={[  
      {
        Name: "edit",
        Class: "variant-filled-secondary",
        Action: async (id) => {
          goto(`/z/pages/portal/projects/zblog/${pid}/sites/${id}/edit`);
        },
      },

      {
        Name: "delete",
        Class: "bg-red-500",
        Action: async (id) => {
          const resp = await api.deleteSite(pid, id);
          if (resp.status !== 200) {
            return;
          }
          load();
        },
      },


    ]}
  />
  