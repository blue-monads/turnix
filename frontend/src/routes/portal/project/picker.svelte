<script lang="ts">
  import type { RootAPI } from "$lib/api";
  import Loader from "$lib/compo/loader/loader.svelte";

  import { getModalStore } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const modalStore = getModalStore();

  let datas: Record<string, any>[] = $state([]);
  let loading = $state(true);

  const load = async () => {
    let api: RootAPI = $modalStore[0].meta["api"];

    const resp = await api.listProjectTypes();
    if (resp.status !== 200) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  let message = "";
</script>

{#if loading}
  <Loader />
{:else}
  <div class="card px-2 w-modal-slim py-4">
    <header class="header flex justify-center">
      <h4 class="h4 mb-2">Select a project type</h4>
    </header>

    <div class="flex flex-col gap-2 overflow-auto max-h-96">
      {#each datas as data}
        {@const icon = data["icon"]}

        <div
          class="flex flex-row justify-between items-center hover:bg-gray-100 rounded py-2 px-10"
        >
          <div class="flex flex-col">
            <h5 class="h-5 font-semibold flex gap-2">
              {#if icon}
                <SvgIcon name={icon} className="w-5 h-5 mt-1" />
              {/if}
              <span>{data["name"]}</span>
            </h5>
            <p class="p">{data["description"]}</p>
          </div>
          <div class="flex flex-row gap-4">
            <a
              href={ data.is_external ? `/z/x/${data['ptype']}` : `/z/pages/portal/projects/${data['ptype']}`  }
              class="underline text-blue-500 text-sm self-center"
              onclick={() => {
                modalStore.close();
              }}
            >
              Home
            </a>

            <a
              href="/z/pages/portal/project/new?ptype={data['ptype']}"
              class="btn btn-sm variant-filled"
              onclick={() => {
                modalStore.close();
              }}
            >
              new
            </a>
          </div>
        </div>
      {/each}
    </div>
  </div>
{/if}

