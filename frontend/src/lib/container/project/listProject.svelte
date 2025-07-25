<script lang="ts">
  import { RootAPI, type ProjectHook } from "$lib/api";
  import { getContext } from "svelte";
  import {
    AppBar,
    getModalStore,
    type ModalSettings,
  } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const api = getContext("__api__") as RootAPI;
  interface Props {
    ptype?: string | undefined;
  }

  let { ptype = undefined }: Props = $props();

  let projects: any[] = $state([]);
  let ptypesMap: Record<string, any> = $state({});

  const load = async () => {
    const resp1 = api.listProjects(ptype);
    const resp2 = await api.listProjectTypes()

    const resp = await resp1;

    if (resp.status !== 200) {
      return;
    }

    if (resp2.status !== 200) {
      return;
    }

    const nextPtypesMap: Record<string, any> = {};
    const ptypes = resp2.data;
    for (const ptype of ptypes) {
      nextPtypesMap[ptype.slug] = ptype;
    }

    ptypesMap = nextPtypesMap;
    projects = resp.data;
  };
  const store = getModalStore();

  load();
</script>

<div class="p-4 flex flex-wrap gap-2">
  {#each projects as proj}
    {@const ptypeData = ptypesMap[proj.ptype]}
    {@const pattern: string = ptypeData?.project_link_pattern}

    {#if ptypeData}
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
          {#if ptypeData.is_external}
            <a
              href={pattern
                ? pattern
                    .replace("{PID}", proj.id)
                    .replace("{PTYPE}", proj["ptype"])
                : `/z/x/${proj["ptype"]}`}
              class="btn variant-filled"
            >
              explore
            </a>
          {:else}
            <a
              href="/z/pages/portal/projects/{proj['ptype']}/{proj['id']}"
              class="btn variant-filled"
            >
              explore
            </a>
          {/if}

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
    {/if}
  {/each}
</div>
