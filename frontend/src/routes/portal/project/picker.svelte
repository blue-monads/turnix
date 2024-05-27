<script lang="ts">

  import type { RootAPI } from "$lib/api";
  import Loader from "$lib/compo/loader/loader.svelte";

  import { getModalStore } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const modalStore = getModalStore();

  let datas: Record<string, any>[] = [];
  let loading = true;

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
  <div class="card p-2 w-modal">
    <header class="header flex justify-center">
      <h4 class="h4">
        <span
          class="bg-gradient-to-br from-pink-500 to-violet-500 bg-clip-text text-transparent box-decoration-clone uppercase"
          >Select New Project type</span
        >
      </h4>
    </header>

    <div class="logo-cloud grid-cols-1 lg:!grid-cols-3 gap-1 border p-2">
      {#each datas as data}
        {@const icon = data["icon"]}
        <a
          class="logo-item"
          on:click={() => {
            modalStore.close();
          }}
          href="/z/pages/portal/project/new?ptype={data['ptype']}"
        >
          {#if icon}
            <SvgIcon name={icon} className="w-6 h-6" />
          {/if}
          <span>{data["name"]}</span>
        </a>
      {/each}
    </div>
  </div>
{/if}
