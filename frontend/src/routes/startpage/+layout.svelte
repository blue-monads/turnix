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

    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    const drawer = getDrawerStore();

    let sibarItems: Record<string, any>[] = [
        {
            name: "Home",
            link: "/z/pages/startpage",
            icon: "home",
        },
        {
            name: "Peers",
            link: "/z/pages/startpage/peers",
            icon: "rss",
        },
        {
            name: "Setting",
            link: "/z/pages/startpage/setting",
            icon: "cog",
        },
    ];

    storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });
</script>

<AppShell>
    <div
        slot="sidebarLeft"
        class="h-screen bg-gradient-to-t variant-gradient-success-warning grad"
    >
        <AppRail width="w-14" background="bg-gradient-to-b from-purple-100 to-indigo-50">
            <div class="flex flex-col">
                {#each sibarItems as item, index}
                    <a
                        class="p-4 hover:bg-secondary-50 [&>*]:pointer-events-none"
                        href={item["link"]}
                    >
                        <SvgIcon
                            name={item["icon"] || "link"}
                            className="w-6 h-6"
                        />
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
                <span></span>
            </div>
        </AppRail>
    </div>

    <div class="h-[calc(100vh-4rem)] md:h-screen bg-gradient-to-b from-indigo-50 to-purple-50">
        <slot />
    </div>
</AppShell>
