<script lang="ts">
    import Autotable from "$lib/compo/autotable/autotable.svelte";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    import { AppBar } from "@skeletonlabs/skeleton";
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4>Peers</h4>
    </svelte:fragment>

    <div class="flex flex-wrap justify-end gap-2">
        <a
            href={`/z/pages/startpage/peers/account/new`}
            class="btn variant-filled btn-sm"
        >
            <SvgIcon className="h-4 w-4" name="plus" />

            peer
        </a>
    </div>
</AppBar>

<Autotable
    actions={[
        {
            Action: async (id) => {},
            Name: "explore",
            Class: "variant-filled-primary",
        },

        {
            Action: async (id) => {},
            Name: "shell",
            Class: "variant-filled-secondary",
        },

        {
            Action: async (id) => {},
            Name: "remove",
            Class: "variant-filled-error",
        },
    ]}
    color={["status"]}
    key_names={[
        ["id", "ID"],
        ["lastSeen", "Last seen"],
        ["status", "Status"],
    ]}

    action_key={"id"}
    
    datas={[
        {
            id: "xx-yy-zz",
            lastSeen: "2022-12-31",
            status: "online",
        },
    ]}
/>
