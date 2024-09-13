<script lang="ts">
    import { NewBookAPI } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { Loader } from "$lib/compo";
    import SaleMaker from "./SaleMaker.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let contacts: object[] = [];
    let loadingContacts = true;

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

    $: __clientIndex = contacts.reduce((acc: any, item: any) => {
        const item_id = item["id"];
        const name = item["name"];
        acc[item_id] = name;
        return acc;
    }, {});

    const submit = async () => {};

    let mode = "invoice";
</script>

{#if loadingContacts}
    <Loader />
{:else}
    <SaleMaker 
        pid={Number(pid)} 
        {api} 
        contactsNameIndex={__clientIndex}
        submit={submit}
        
       
        
        />
{/if}
