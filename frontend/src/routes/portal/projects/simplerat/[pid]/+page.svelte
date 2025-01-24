<script lang="ts">
    import { AutoTable } from "$lib/compo";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { NewSimpleRATApi } from "../lib/SimpleRATApi";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import AddDeviceDialog from "./new/AddDeviceDialog.svelte";

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
    ]}
    datas={devices}
    color={[]}
    actions={[]}
/>
