<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";

    let pid = $params["pid"];

    const tools = [
        {
            name: "Query",
            icon: "squares-2x2",
            href: "/z/pages/portal/project/livequery",
        },
        {
            name: "Explore",
            icon: "squares-2x2",
            href: "/z/pages/portal/project/explore",
        },
        {
            name: "Hooks",
            icon: "code-bracket-square",
            href: "/z/pages/portal/hooks",
        },
        {
            name: "Config",
            icon: "cog",
            href: "#",
        },
        {
            name: "Reload",
            icon: "arrow-path",
            href: "#",
        }, {
            name: "Pair",
            icon: "signal",
            href: "#",
        }

    ];
</script>

<section class="p-6 space-y-6 rounded-xl">
    <div class="flex flex-col gap-6">
        {#each tools as item}
            <a
                href={`${item.href}?pid=${pid}`}
                class="px-4 py-2 text-lg text-gray-100 transition variant-filled rounded-md h-14 w-44 btn"
            >
                <SvgIcon name={item.icon} className="w-6 h-6 mr-2" />

                {item.name}
            </a>
        {/each}
    </div>
</section>

