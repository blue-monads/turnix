<script lang="ts">
    import { getModalStore } from "@skeletonlabs/skeleton";

    const store = getModalStore();

    const data = $store[0].meta;
    const pid = data.pid;

    const pluginSources = [
        {
            name: "URL",
            description: "Import plugin from a URL",
            link: "/todo",
        },
        {
            name: "File",
            description: "Upload a plugin file",
            link: "/todo",
        },
        {
            name: "Live Editor",
            description: "Edit a plugin in the editor",
            link: "/todo",
        },
        {
            name: "Raw Create",
            description: "Create a plugin from scratch",
            link: `/z/pages/portal/project/plugins/new?pid=${pid}`,
        },
    ];
</script>

<div class="w-modal-slim bg-white rounded card p-4">
    <h3 class="h3">Plugins Sources</h3>

    <section class="flex flex-col gap-2">
        {#each pluginSources as pluginSource}
            <div
                class="flex flex-row justify-between items-center hover:bg-gray-100 rounded p-1"
            >
                <div class="flex flex-col">
                    <h5 class="h-5 font-semibold">{pluginSource.name}</h5>
                    <p class="p">{pluginSource.description}</p>
                </div>
                <div class="flex flex-row">
                    <a
                        href={pluginSource.link}
                        class="btn btn-sm variant-filled"
                        on:click={() => {
                            store.close();
                        }}
                    >
                        open
                    </a>
                </div>
            </div>
        {/each}
    </section>
</div>
