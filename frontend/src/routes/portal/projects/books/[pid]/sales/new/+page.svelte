<script lang="ts">
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { Loader } from "$lib/compo";
    import SaleMaker from "./SaleMaker.svelte";
    import { goto } from "$app/navigation";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let contacts: object[] = $state([]);
    let loadingContacts = $state(true);

    const loadData = async () => {
        loadingContacts = true;
        const resp = await api.listContacts(pid);
        if (resp.status !== 200) {
            return;
        }

        contacts = resp.data;
        loadingContacts = false;
    };

    loadData();

    let __clientIndex = $derived(contacts.reduce((acc: any, item: any) => {
        const item_id = item["id"];
        const name = item["name"];
        acc[item_id] = name;
        return acc;
    }, {}));

    const submit = async (data: Record<string, any>) => {
        console.log("@submit/data", data);

        data.sale.sales_date = new Date(data.sale.sales_date).toISOString();

        const resp = await api.addSale(pid, data as any);
        if (resp.status !== 200) {
            return resp.data.message;
        }

        goto(`/z/pages/portal/projects/books/${pid}/sales`);
        
        return "";

    };

    let mode = "invoice";
</script>

{#if loadingContacts}
    <Loader />
{:else}
    <SaleMaker 
        name="New Sale"
        pid={Number(pid)} 
        {api} 
        contactsNameIndex={__clientIndex}
        submit={submit}
        
       
        
        />
{/if}
