<script lang="ts">
    import { NewBookAPI, type Catagory } from "$lib/projects/books";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import type { NewSaleLine } from "./sub/sales";
    import {
        getModalStore,
        RadioGroup,
        RadioItem,
    } from "@skeletonlabs/skeleton";
    import { Loader } from "$lib/compo";
    import SaleMaker from "./SaleMaker.svelte";

    const pid = $page.params["pid"];
    const api = NewBookAPI(getContext("__api__") as RootAPI);

    let message = "";

    const store = getModalStore();

    let title = "";
    let info = "";
    let totalPrice = 0;
    let client_id = 0;
    let sales_date = new Date().toISOString().slice(0, 16);

    let showTitle = !!title;

    let lines: NewSaleLine[] = [];
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
        bind:client_id={client_id}
        
       
        
        />
{/if}
