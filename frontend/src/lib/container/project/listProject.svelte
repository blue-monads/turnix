<script lang="ts">
  import { RootAPI, type ProjectHook } from "$lib/api";
  import { getContext } from "svelte";
  import {
    AppBar,
    getModalStore,
    type ModalSettings,
  } from "@skeletonlabs/skeleton";
  import { FloatyButton } from "$lib/compo";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const api = getContext("__api__") as RootAPI;
  interface Props {
    ptype?: string | undefined;
  }

  let { ptype = undefined }: Props = $props();

  let projects: any[] = $state([]);

  const load = async () => {
    const resp = await api.listProjects(ptype);
    if (resp.status !== 200) {
      return;
    }

    projects = resp.data;
  };
  const store = getModalStore();

  load();
</script>



<div class="p-4 flex flex-wrap gap-2">
  {#each projects as proj}
    <div class="card p-1 w-96 hover:bg-secondary-backdrop-token">
      <header class="card-header">
        <img src={`https://picsum.photos/seed/${proj.id}/512/300`} alt="" />
      </header>
      <section class="p-4 flex flex-col">
        <div class="flex justify-end">
          <span class="chip variant-filled-tertiary bg-tertiary-hover-token"
            >{proj.ptype}</span
          >
        </div>

        <h3 class="h3">{proj.name}</h3>

        <p class="p">{proj.info}</p>
      </section>

      <footer class="card-footer flex justify-end gap-2">
        <a
          href="/z/pages/portal/projects/{proj['ptype']}/{proj['id']}"
          class="btn variant-filled"
        >
          explore
        </a>

        <button
          onclick={() => {
            store.trigger({
              type: "component",
              component: "project_options",
              meta: { api, data: proj },
            });
          }}
          class="btn variant-filled variant-filled-warning"
        >
          <SvgIcon name="adjustments-horizontal" className="w-6 h-6" />
        </button>
      </footer>
    </div>
  {/each}
</div>
