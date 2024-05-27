<script lang="ts">
  import type { RootAPI } from "$lib/api";
  import Loader from "$lib/compo/loader/loader.svelte";

  import { getModalStore } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const modalStore = getModalStore();

  let datas = [];
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
          >Options</span
        >
      </h4>
    </header>

    <div class="logo-cloud flex flex-col gap-1 border p-2">
      <a
        class="logo-item"
        on:click={() => {
          modalStore.close();
        }}
        href="/z/pages/portal/projects/{$modalStore[0].meta['data'][
          'ptype'
        ]}?pid={$modalStore[0].meta['data']['id']}"
      >
        <SvgIcon name="document-text" className="w-6 h-6" />
        <span>Project Home</span>
      </a>

      <a
        class="logo-item"
        on:click={() => {
          modalStore.close();
        }}
        href="/z/pages/portal/hooks"
      >
        <SvgIcon name="folder" className="w-6 h-6" />
        <span>Project Files</span>
      </a>

      <a
        class="logo-item"
        on:click={() => {
          modalStore.close();
        }}
        href={`/z/pages/portal/hooks?ptype=${$modalStore[0].meta['data']['ptype']}&pid=${$modalStore[0].meta['data']['id']}`}
      >
        <SvgIcon name="code-bracket-square" className="w-6 h-6" />
        <span>Project Hooks</span>
      </a>

      <a
        class="logo-item"
        on:click={() => {
          modalStore.close();
        }}
        href="/z/pages/portal/projects"
      >
        <SvgIcon name="pencil-square" className="w-6 h-6" />
        <span>Edit Project</span>
      </a>
    </div>
  </div>
{/if}
