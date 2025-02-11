<script lang="ts">
  import type { RootAPI } from "$lib/api";
  import Loader from "$lib/compo/loader/loader.svelte";

  import { getModalStore } from "@skeletonlabs/skeleton";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

  const modalStore = getModalStore();

  let datas = [];
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

  const data = $modalStore[0].meta["data"];
  const pid = data["id"];
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
        onclick={() => {
          modalStore.close();
        }}
        href="/z/pages/portal/projects/{['ptype']}?pid={data['id']}"
      >
        <SvgIcon name="document-text" className="w-6 h-6" />
        <span>Project Home</span>
      </a>
      <a
        class="logo-item"
        onclick={() => {
          modalStore.close();
        }}
        href={`/z/pages/portal/project/files/${pid}`}
      >
        <SvgIcon name="folder" className="w-6 h-6" />
        <span>Project Files</span>
      </a>

      <a
        class="logo-item"
        onclick={() => {
          modalStore.close();
        }}
        href={`/z/pages/portal/project/plugins?pid=${pid}`}
      >
        <SvgIcon name="squares-2x2" className="w-6 h-6" />
        <span>Plugins</span>
      </a>

      <a
      class="logo-item"
      onclick={() => {
        modalStore.close();
      }}
      href={`/z/pages/portal/project/tools/livequery?pid=${pid}`}
    >
      <SvgIcon name="wrench-screwdriver" className="w-6 h-6" />
      <span>Tools</span>
    </a>

      <a
        class="logo-item"
        onclick={() => {
          modalStore.close();
        }}
        href={`/z/pages/portal/project/users?pid=${pid}`}
      >
        <SvgIcon name="users" className="w-6 h-6" />
        <span>Users and Permissions</span>
      </a>



      <a
        class="logo-item"
        onclick={() => {
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
