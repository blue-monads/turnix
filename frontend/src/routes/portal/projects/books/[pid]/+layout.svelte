<script>
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    const pid = $page.params["pid"];

    let active = true;

    const sideBarList = [
        {
            name: "Accounts",
            link: `/z/pages/portal/projects/books/${pid}/account`,
            icon: "users",
            active: false,
        },
        {
            name: "Transactions",
            link: `/z/pages/portal/projects/books/${pid}/txn`,
            icon: "clipboard-document",
            active: false,
        },
        {
            name: "Reports",
            link: `/z/pages/portal/projects/books/${pid}/report`,
            icon: "document-chart-bar",
            active: false,
        },
        {
            name: "Inventory",
            link: `/z/pages/portal/projects/books/${pid}/inventory`,
            icon: "archive-box",
            active: false,
        },
        {
            name: "Sales",
            link: `/z/pages/portal/projects/books/${pid}/sales`,
            icon: "banknotes",
            active: false,
        },
        {
            name: "Invoice",
            link: `/z/pages/portal/projects/books/${pid}/invoice`,
            icon: "document-text",
            active: false,
        },
        {
            name: "Settings",
            link: `/z/pages/portal/projects/books/${pid}/actions`,
            icon: "cog",
            active: false,
        },
    ];
</script>

<div class="flex relative h-screen flex-col md:flex-row md:w-full md:h-32">
    <div
        id="sidebar"

        class="p-4  flex flex-col gap-2 transition-all duration-300 ease-in-out md:flex-row flex-wrap
        {active ? 'w-48': 'hidden'}"
    >
        <div class="mb-10 text-center variant-outline uppercase font-medium font-token hidden md:block p-1 w-full">
            <h4 class="h4">Books</h4> 
        </div>

        {#each sideBarList as item}
            <a
                href={item.link}
                class="flex items-center px-3 py-2 text-gray-600 transition-colors duration-300 transform rounded-lg hover:bg-gray-100 hover:text-gray-700"

            >
                <SvgIcon className="h-4 w-4 mr-1" name={item.icon} />
                {item.name}
            </a>
        {/each}


    </div>

    <div class="absolute md:bottom-8 ">
        <button
            class=" rounded border-2 bg-slate-100 p-1"
            on:click={() => {
                active = !active;
            }}
        >
            <SvgIcon className="h-4 w-4" name="bars-3-center-left" />
        </button>
    </div>

    <main class="grow w-full h-screen overflow-auto">
        <slot />
    </main>
</div>
