<script lang="ts">
    import { getContext } from "svelte";
    import type { ProjectPlugin, RootAPI } from "$lib/api";
    import { params } from "$lib/params";
    import { AutoTable, Loader } from "$lib/compo";
    import { goto } from "$app/navigation";
    import {
        AppBar,
        getModalStore,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";

    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    const api = getContext("__api__") as RootAPI;
    let pid = $params["pid"];

    let plugins: ProjectPlugin[] = $state([]);
    let loading = $state(true);

    const load = async () => {
        loading = true;
        const resp = await api.listProjectPlugins(pid);
        if (resp.status !== 200) {
            return;
        }
        plugins = resp.data;
        loading = false;
    };

    load();

    const store = getModalStore();

    const remove = async (id: string) => {
        const resp = await api.removeProjectPlugin(pid, id);
        if (resp.status !== 200) {
            return;
        }

        load();
    };
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4">Projects Plugins</h4>
    </svelte:fragment>

    <svelte:fragment slot="trail">
        <button
            class="btn btn-sm variant-ghost-secondary"
            onclick={() => {

                store.trigger({
                    type: "component",
                    component: "project_new_plugin_picker",
                    meta: {
                        pid: pid,
                    },
                });
            }}
        >
            <SvgIcon name="plus" className="w-6 h-6" />

            Plugin
        </button>
    </svelte:fragment>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <AutoTable
        datas={plugins}
        action_key="id"
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["ptype", "type"],
            ["created_at", "Created At"],
            ["updated_at", "Updated At"],
        ]}
        color={["ptype"]}
        actions={[
            {
                Name: "Run",
                Class: "variant-filled-success",
                Action: async (id, data) => {
                    console.log("@run", id, data);
                    goto(
                        `/z/pages/portal/project/plugins/run?pid=${pid}&id=${id}`,
                    );
                },
            },
            {
                Name: "Live Editor",
                Class: "variant-filled-primary",
                Action: async (id, data) => {
                    console.log("@view", id, data);

                    goto(
                        `/z/pages/portal/project/plugins/view?pid=${pid}&id=${id}`,
                    );
                },
            },

            {
                Name: "Raw Edit",
                Class: "variant-filled-secondary",

                Action: async (id, data) => {
                    console.log("@edit", id, data);

                    goto(
                        `/z/pages/portal/project/plugins/edit?pid=${pid}&id=${id}`,
                    );
                },
            },
            {
                Name: "remove",
                Class: "variant-filled-error",
                Action: remove,
            },
        ]}
    />
{/if}
