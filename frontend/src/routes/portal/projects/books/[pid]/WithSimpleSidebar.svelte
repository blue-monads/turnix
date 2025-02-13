<script lang="ts">
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    interface Props {
        activeKey?: string;
        sidebarList: { name: string; link: string; icon: string; key: string }[];
        children?: () => any;
        name: string;
    }

    let { sidebarList, children, activeKey: sidebarKey, name }: Props = $props();



    let active = $state(true);


</script>

<div
    class="flex relative h-screen flex-col md:flex-row md:w-full md:h-32 border-r border-gray-800"
>
    <div
        id="sidebar"
        class="p-4 flex flex-col gap-2 transition-all duration-300 ease-in-out md:flex-row flex-wrap
        {active ? 'w-48' : 'hidden'}"
    >
        <div
            class="mb-10 text-center variant-outline font-medium font-token hidden md:block p-1 w-full"
        >
            <h4 class="h4">{name}</h4>
        </div>

        {#each sidebarList as item}
            <a
                href={item.link}
                class="flex items-center px-2 py-2 w-full text-gray-600 transition-colors duration-300 transform rounded-lg hover:bg-gray-100 hover:text-gray-700 {sidebarKey ===
                item.key
                    ? 'bg-gray-100 text-gray-700'
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
            onclick={() => {
                active = !active;
            }}
        >

        <span class="hidden md:block">
            <SvgIcon className="h-4 w-4" name={ active ? 'chevron-double-left' : 'chevron-double-right' } />
        </span>

        <span class="block md:hidden"> 
            <SvgIcon className="h-4 w-4" name={ active ? 'chevron-double-up' : 'chevron-double-down' } />
        </span>


        </button>
    </div>

    <main class="grow w-full h-screen overflow-auto">
        {@render children?.()}
    </main>
</div>
