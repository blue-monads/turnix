<script>
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    /** @type {{children?: import('svelte').Snippet}} */
    let { children } = $props();
    const pid = $page.params["pid"];

    let active = $state(true);

    const sideBarList = [
        {
            name: "Accounts",
            link: `/z/pages/portal/projects/books/${pid}/account`,
            icon: "users",
            key: "account",
        },
        {
            name: "Transactions",
            link: `/z/pages/portal/projects/books/${pid}/txn`,
            icon: "clipboard-document",
            key: "txn",
        },
        {
            name: "Inventory",
            link: `/z/pages/portal/projects/books/${pid}/inventory`,
            icon: "archive-box",
            key: "inventory",
        },
        {
            name: "Estimates",
            link: `/z/pages/portal/projects/books/${pid}/estimates`,
            icon: "document",
            key: "estimates",

        },
        {
            name: "Sales",
            link: `/z/pages/portal/projects/books/${pid}/sales`,
            icon: "banknotes",
            key: "sales",
        },
        {
            name: "Reports",
            link: `/z/pages/portal/projects/books/${pid}/report`,
            icon: "document-chart-bar",
            key: "report",
        },

        {
            name: "Contacts",
            link: `/z/pages/portal/projects/books/${pid}/contacts`,
            icon: "user-group",
            key: "contacts",
        },

        {
            name: "Notepad",
            link: `/z/pages/portal/projects/books/${pid}/notepad`,
            icon: "book-open",
            key: "notepad",
        },
        {
            name: "Settings",
            link: `/z/pages/portal/projects/books/${pid}/actions`,
            icon: "cog",
            key: "actions",
        },
    ];

    let sidebarKey = $derived($page.url.pathname.split("/").at(7));

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
            class="mb-10 text-center variant-outline uppercase font-medium font-token hidden md:block p-1 w-full"
        >
            <h4 class="h4">Books</h4>
        </div>

        {#each sideBarList as item}
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
            <SvgIcon className="h-4 w-4" name="{ active ? 'chevron-double-left' : 'chevron-double-right' }" />
        </span>

        <span class="block md:hidden"> 
            <SvgIcon className="h-4 w-4" name="{ active ? 'chevron-double-up' : 'chevron-double-down' }" />
        </span>


        </button>
    </div>

    <main class="grow w-full h-screen overflow-auto">
        {@render children?.()}
    </main>
</div>
