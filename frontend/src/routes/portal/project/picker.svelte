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
              href={data.is_external
                ? `/z/x/${data["ptype"]}`
                : `/z/pages/portal/projects/${data["ptype"]}`}
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

            <button class="hover:bg-gray-100 p-1"
            onclick={async () => {
              let api: RootAPI = $modalStore[0].meta["api"];
              await api.getProjectTypeReload(data["ptype"]);
              load();
            }}
            
            >

                <SvgIcon className="w-4 h-4" name="arrow-path" />
            </button>


          </div>
        </div>
      {/each}
    </div>

    <div class="flex justify-end px-10 py-2">
      <a
        href="/z/pages/portal/store"
        class="btn btn-sm variant-filled"
        onclick={() => {
          modalStore.close();
        }}
      >
        <SvgIcon name={"shopping-bag"} className="w-5 h-5" />
        Store
      </a>
    </div>
  </div>
{/if}
