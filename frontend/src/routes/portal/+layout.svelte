<script lang="ts">
  import {
    AppBar,
    AppRail,
    AppRailAnchor,
    AppRailTile,
    AppShell,
    type DrawerSettings,
    type PopupSettings,
  } from "@skeletonlabs/skeleton";
  import { popup } from "@skeletonlabs/skeleton";
  import {
    computePosition,
    autoUpdate,
    offset,
    shift,
    flip,
    arrow,
  } from "@floating-ui/dom";
  import { getDrawerStore } from "@skeletonlabs/skeleton";

  import { storePopup } from "@skeletonlabs/skeleton";
  import { RootAPI } from "$lib/api/api";
  import ContextThis from "./contextThis.svelte";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  interface Props {
    children?: import('svelte').Snippet;
  }

  let { children }: Props = $props();

  let currentTile = "";
  let showMobileModal = $state(false);
  const drawer = getDrawerStore();

  let sibarItems: Record<string, any>[] = [
    {
      name: "Home",
      link: "/z/pages/portal",
      icon: "home",
    },
    {
      name: "Projects",
      link: "/z/pages/portal/project",
      icon: "queue-list",
    },
    {
      name: "Profile",
      link: "/z/pages/portal/self",
      icon: "user-circle",
    },

    {
      name: "Users",
      link: "/z/pages/portal/self/users",
      icon: "users",
    },
    {
      name: "Files",
      link: "/z/pages/portal/self/files",
      icon: "folder",
    },
  ];

  sibarItems.push({
    name: "Playground",
    link: "/z/pages/portal/playground",
    icon: "code-bracket-square",
  });

  let api: RootAPI;

  const load = () => {
    api = new RootAPI();
  };

  const toggle = () => {
    showMobileModal = !showMobileModal;
  };

  const popupHover: PopupSettings = {
    event: "hover",
    target: "popupHover",
    placement: "right",
  };

  const drawerSettings: DrawerSettings = {
	id: 'notification',
  width: "w-modal-slim",
	meta: { 
    getRootAPI: () => api 
  }
};

  storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

  load();
</script>

{#if showMobileModal}
  <div class="h-screen w-screen bg-gray-600 bg-opacity-90 fixed z-50">
    <div
      class="h-full w-full absolute transform translate-x-0 transition ease-in-out duration-700 p-5"
    >
      <div class="h-full bg-white rounded p-5">
        <div class="absolute right-4">
          <button
            onclick={toggle}
            aria-label="close menu modal"
            class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 cursor-pointer rounded border mr-2"
          >
            <SvgIcon name="x-mark" className="h-4 w-4" />
          </button>
        </div>

        <ul class="space-y-4 pt-10">
          {#each sibarItems as item, index}
            <li class="border rounded">
              <a
                onclick={toggle}
                class="h-12 pl-24  flex justify-start items-center w-full text-gray-700 focus:text-primary-600 hover:text-primary-800 uppercase gap-2"
                href={item["link"]}
              >
                <SvgIcon name={item["icon"] || "link"} className="w-6 h-6" />
                {item["name"] || ""}
              </a>
            </li>
          {/each}
        </ul>
      </div>
    </div>
  </div>
{/if}

<AppShell>
  <div
    slot="header"
    class="flex flex-row md:hidden justify-between bg-surface-100-800-token"
  >
    <div class="flex p-1">
      <button onclick={toggle}>
        <SvgIcon name="bars-3" className="w-6 h-6" />
      </button>
    </div>

    <div class="flex gap-4 p-1">
      <button
        onclick={() => {
          drawer.open(drawerSettings);
        }}
        class="flex flex-col justify-center items-center rounded-full bg-secondary-100 hover:bg-secondary-200 p-2"
      >
        <SvgIcon name="bell" className="w-6 h-6" />
      </button>
    </div>
  </div>

  <div
    slot="sidebarLeft"
    class="h-screen hidden md:block bg-gradient-to-br variant-gradient-primary-secondary"
  >
    <AppRail width="w-14">
      <div slot="lead" class="mb-4">
        <AppRailAnchor href="/z/pages/portal">
          <div class="flex justify-center items-center p-2">
            <img src="/z/lib/images/logo.png" alt="logo" />
          </div>
        </AppRailAnchor>
      </div>

      <div class="flex flex-col">
        {#each sibarItems as item, index}
          {@const popupHover = {
            event: "hover",
            target: `popupHover-${index}`,
            placement: "right",
          }}
          <a
            use:popup={popupHover as any}
            class="p-4 hover:bg-secondary-50 [&>*]:pointer-events-none [&]:z-20"
            href={item["link"]}
          >
            <SvgIcon name={item["icon"] || "link"} className="w-6 h-6" />
          </a>

          <div
            class="card p-2 variant-filled-secondary"
            data-popup={`popupHover-${index}`}
          >
            <p>{item["name"] || ""}</p>
          </div>
        {/each}
      </div>

      <div slot="trail" class="mb-4 p-2">
        <button
          onclick={() => {
            drawer.open(drawerSettings);
          }}
          class="flex flex-col justify-center items-center rounded-full bg-secondary-100 hover:bg-secondary-200 p-2"
        >
          <SvgIcon name="bell" className="w-6 h-6" />
        </button>
      </div>
    </AppRail>
  </div>

  <div class="h-[calc(100vh-4rem)] md:h-screen">
    <ContextThis {api}>
      <svelte:fragment>
        {@render children?.()}
      </svelte:fragment>
    </ContextThis>
  </div>
</AppShell>
