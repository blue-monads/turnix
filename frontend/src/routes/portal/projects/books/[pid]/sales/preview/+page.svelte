<script lang="ts">
    import { NewBookAPI, type Sale, type SalesData } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { Loader } from "$lib/compo";
    import { params } from "$lib/params";
    import SalesPreview from "./SalesPreview.svelte";

    const pid = $page.params["pid"];
    const sid = $params["sid"];

    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let data: SalesData | undefined = $state();
    let contacts: object[] = $state([]);
    let loadingContacts = $state(true);
    let loadingData = $state(true);

    const loadContact = async () => {
        loadingContacts = true;
        const resp = await api.listContacts(pid);
        if (resp.status !== 200) {
            return;
        }

        contacts = resp.data;
        loadingContacts = false;
    };

    const loadData = async () => {
        loadingData = true;
        const resp = await api.getSale(pid, sid);
        if (resp.status !== 200) {
            return;
        }
        data = resp.data;
        loadingData = false;

        console.log("@sale_data", resp.data);
    };

    loadData();
    loadContact();

    let __clientIndex = $derived(
        contacts.reduce((acc: any, item: any) => {
            const item_id = item["id"];
            const name = item["name"];
            acc[item_id] = name;
            return acc;
        }, {}),
    );
</script>

{#if loadingContacts || loadingData}
    <Loader />
{:else}
    <SalesPreview
        lines={data?.lines || []}
        data={data?.sale as any}
        pid={Number(pid)}
        contactsNameIndex={__clientIndex}
    />
{/if}
