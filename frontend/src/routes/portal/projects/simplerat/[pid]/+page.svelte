<script lang="ts">
    import { AutoTable } from "$lib/compo";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { NewSimpleRATApi } from "../lib/SimpleRATApi";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import AddDeviceDialog from "./AddDeviceDialog.svelte";
    import { goto } from "$app/navigation";
    import { webspcketTest } from "./wstest";

    const pid = $page.params["pid"];

    const store = getModalStore();
    const api = NewSimpleRATApi(getContext("__api__") as RootAPI);
    const modal = getContext("__vs_modal__") as any;

    let devices: any[] = $state([]);

    const load = async () => {
        const resp = await api.listDevices(pid);
        if (resp.status !== 200) {
            return;
        }

        devices = resp.data;
    };

    load();
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/simplerat"
                    >SimpleRATS</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Devices</li>
        </ol>
    </svelte:fragment>

    <div class="flex flex-wrap justify-end gap-2">
        <button
            class="btn variant-filled btn-sm"
            onclick={() => {
                modal.show(AddDeviceDialog, {
                    pid: pid,
                    onAddDevice: async () => {
                        modal.close();
                        load();
                    },
                });
            }}
        >
            <SvgIcon className="h-4 w-4" name="plus" />

            Device
        </button>
    </div>
</AppBar>

<AutoTable
    action_key={"id"}
    key_names={[
        ["id", "ID"],
        ["name", "Name"],
        ["created_at", "Created At"],
        ["status", "Status"],
    ]}
    datas={devices}
    hashSeed={99}
    color={["status"]}
    show_drop={true}
    actions={[
        {
          Name: "Ping",
          Class: "variant-filled-primary",
          icon: "plus",
          Action: async (id) => {

            const resp = await api.performDeviceAction(pid, id, "ping", {});
            if (resp.status !== 200) {
                console.error("Failed to ping device", resp);
                return;
            }

            console.log("Pinged device", resp.data);
    
          },
        },

        {
            Name: "Service",
            Class: "",
            Action: async (id) => {


               await webspcketTest(api, pid, id);
        
            },
            
      
        },




        {
          Name: "Shell",
          Class: "variant-filled-secondary",
          icon: "plus",
          Action: async (id) => {
            goto(`/z/pages/portal/projects/simplerat/${pid}/device/${id}/shell`);
          },
        },

        {
          Name: "FS",
          Class: "variant-filled-tertiary",
          icon: "plus",        
          Action: async (id) => {
            goto(`/z/pages/portal/projects/simplerat/${pid}/device/${id}/fs`);
            
          },

        },



    ]}
/>
