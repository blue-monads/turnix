<script lang="ts">
    import { Tab, TabGroup } from "@skeletonlabs/skeleton";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { params } from "$lib/params";
    import { } from "svelte/store";
    import { goto, onNavigate } from "$app/navigation";

    let pid = $params["pid"];

    interface Props {
        children?: import("svelte").Snippet;
    }

    let { children }: Props = $props();

    onNavigate((nav) => {
        console.log("Navigating", nav);
    });

    const tools = [
        {
            name: "Query",
            href: "/z/pages/portal/project/tools/livequery",
        },
        {
            name: "Explore",
            href: "/z/pages/portal/project/tools/explore",
        },
        {
            name: "Hooks",
            href: "/z/pages/portal/project/tools/hooks",
        },
        {
            name: "Config",

            href: "#",
        },
        {
            name: "Reload",
            href: "#",
        },
        {
            name: "Pair",
            href: "#",
        },
    ];

    let tabSet: string = $state(tools[0].name);
</script>

<div class="flex-1 w-full">
    <TabGroup>
        {#each tools as item}
            <Tab
                bind:group={tabSet}
                name={item.name}
                value={item.name}
                on:change={() => {
                    // window.location.href = `${item.href}?pid=${pid}`;
                    goto(`${item.href}?pid=${pid}`);
                }}
            >
                <span>{item.name}</span>
            </Tab>
        {/each}
    </TabGroup>
</div>

{@render children?.()}
