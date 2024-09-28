<script lang="ts">
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { onMount } from "svelte";
    import { buildApi, type StartAPI } from "./startAPI";
    import Api from "./api.svelte";
    const pid = $page.params["pid"];

    let active = true;

    let sibarItems: Record<string, any>[] = [
        {
            name: "Home",
            link: "/z/pages/startpage",
            icon: "home",
            key: "home",
        },
        {
            name: "Peers",
            link: "/z/pages/startpage/peers",
            icon: "rss",
            key: "peers",
        },
        {
            name: "Setting",
            link: "/z/pages/startpage/setting",
            icon: "cog",
            key: "setting",
        },
    ];

    $: sidebarKey = $page.url.pathname.split("/").at(4);
    $: console.log(sidebarKey);

    onMount(() => {
        const g = window as any;
        g["__handle_rpc__"] = function (name: string, data: any) {
            console.log("@bang/bang");
            console.log(name, data);
        };
    });

    let api: StartAPI;

    buildApi(location.origin).then((_api) => {
        api = _api;
    });
</script>

<div
    class="flex relative h-screen flex-col md:flex-row md:w-full border-r border-gray-800"
>
    {#if sidebarKey}
        <div
            id="sidebar"
            class="p-4 flex flex-col gap-2 transition-all duration-300 ease-in-out bg-blend-darken bg-gradient-to-b from-purple-50 shadow
{active ? 'w-48' : 'hidden'}"
        >
            <div
                class="p-2 border text-center variant-outline uppercase font-medium font-token"
            >
                <h4>Startpage</h4>
            </div>

            {#each sibarItems as item}
                <a
                    href={item.link}
                    class="flex items-center px-2 py-2 w-full text-gray-600 transition-colors duration-300 transform rounded-lg hover:bg-gray-300 hover:text-gray-700 {sidebarKey ===
                    item.key
                        ? 'bg-gray-400 text-gray-900'
                        : ''}"
                >
                    <SvgIcon className="h-4 w-4 mr-1" name={item.icon} />
                    {item.name}
                </a>
            {/each}
        </div>

        <div class="absolute md:bottom-8">
            <button
                class=" rounded border-2 bg-slate-100 p-1"
                on:click={() => {
                    active = !active;
                }}
            >
                <span class="hidden md:block">
                    <SvgIcon
                        className="h-4 w-4"
                        name={active
                            ? "chevron-double-left"
                            : "chevron-double-right"}
                    />
                </span>

                <span class="block md:hidden">
                    <SvgIcon
                        className="h-4 w-4"
                        name={active
                            ? "chevron-double-up"
                            : "chevron-double-down"}
                    />
                </span>
            </button>
        </div>
    {/if}

    <main class="grow w-full h-screen overflow-auto">
        {#if api}
            <Api {api}>
                <slot />
            </Api>
        {/if}
    </main>
</div>
