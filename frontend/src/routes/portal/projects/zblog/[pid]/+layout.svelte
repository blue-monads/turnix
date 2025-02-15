<script>
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import WithSimpleSidebar from "../../books/[pid]/WithSimpleSidebar.svelte";

    let { children } = $props();
    const pid = $page.params["pid"];


    const sideBarList = [   
        {
            name: "Posts",
            link: `/z/pages/portal/projects/zblog/${pid}`,
            icon: "users",
            key: "",
        },
        {
            name: "Sites",
            link: `/z/pages/portal/projects/zblog/${pid}/sites`,
            icon: "globe-alt",
            key: "sites",            
        },
        {
            name: "Tags",
            link: `/z/pages/portal/projects/zblog/${pid}/tags`,
            icon: "tag",
            key: "tags",
        },
        {
            name: "Themes",
            link: `/z/pages/portal/projects/zblog/${pid}/themes`,
            icon: "photo",
            key: "themes",
        }

    ];

    let sidebarKey = $derived($page.url.pathname.split("/").at(7));

</script>

<WithSimpleSidebar
    sidebarList={sideBarList} 
    activeKey={sidebarKey}
    name="ZBlog"
    >
    {@render children?.()}
</WithSimpleSidebar>