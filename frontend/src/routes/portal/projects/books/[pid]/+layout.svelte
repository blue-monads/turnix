<script>
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import WithSimpleSidebar from "./WithSimpleSidebar.svelte";

    let { children } = $props();
    const pid = $page.params["pid"];


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

<WithSimpleSidebar 
    sidebarList={sideBarList} 
    activeKey={sidebarKey}
    name="Books"
    >
    {@render children?.()}
</WithSimpleSidebar>