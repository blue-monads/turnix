<script lang="ts">
    import { goto } from "$app/navigation";
    import type { RootAPI } from "$lib/api";
    import { Loader } from "$lib/compo";

    import {
        NewBookAPI,
        type TransactionWithLine,
        type TxnData,
        type TxnUpdateWithLineOptions,
    } from "$lib/projects/books";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import { getContext } from "svelte";
    import { page } from "$app/stores";
    import TxnForm from "../TxnForm.svelte";
    import { params } from "$lib/params";

    const pid = $page.params["pid"];
    const tid = $params["tid"];

    const api = NewBookAPI(getContext("__api__") as RootAPI);
    const store = getModalStore();

    let loading = true;

    let oldData: TxnData;
    let debitLineId = 0;
    let creditLineId = 0;
    let firstIsDebit = true;

    let title = "";
    let notes = "";
    let debit_account_id = 0;
    let credit_account_id = 0;
    let debit_amount = 0;
    let credit_amount = 0;
    let reference_id = "";

    const onSubmit = async (data: TxnData) => {
        let first_line_data: Record<string, any> = {};
        let second_line_data: Record<string, any> = {};

        const updateData: TxnUpdateWithLineOptions = {
            first_line_data,
            second_line_data,
            first_line_id: firstIsDebit ? debitLineId : creditLineId,
            second_line_id: firstIsDebit ? creditLineId : debitLineId,
            txn_data: {},
        };

        if (data.title !== oldData.title) {
           updateData.txn_data["title"] = data.title            
        }

        if (data.notes !== oldData.notes) {
            updateData.txn_data["notes"] = data.notes
        }

        // if (data.attachments !== oldData.attachments) {
        //     updateData.txn_data["attachments"] = data.attachments
        // }

        if (data.reference_id !== data.reference_id) {
            updateData.txn_data["reference_id"] = data.reference_id
        }

        if (!debitLineId) {
            second_line_data = updateData.first_line_data
            first_line_data = updateData.second_line_data
        }

        if (data.debit_account_id !== oldData.debit_account_id) {
            updateData.first_line_data["account_id"] = data.debit_account_id
        }

        if (data.credit_account_id !== oldData.credit_account_id) {
            updateData.second_line_data["account_id"] = data.credit_account_id
        }

        if (data.debit_amount !== oldData.debit_amount) {
            updateData.first_line_data["debit_amount"] = data.debit_amount
        }

        if (data.credit_amount !== oldData.credit_amount) {
            updateData.second_line_data["credit_amount"] = data.credit_amount
        }

        loading = true

        await api.updateTxnWithLine(pid, tid, updateData)

        await load()


    };

    const load = async () => {
        loading = true;

        const resp = await api.getTxnWithLines(pid, tid);
        if (resp.status !== 200) {
            console.log("@err", loading);
            return;
        }

        let data = resp.data;
        title = data.txn.title;
        notes = data.txn.notes;
        reference_id = data.txn.reference_id;

        let debitLine = data.first_line;
        let creditLine = data.second_line;
        if (debitLine?.debit_amount === 0) {
            firstIsDebit = false;
            creditLine = data.first_line;
            debitLine = data.second_line;
        }

        debitLineId = debitLine?.id || 0;
        creditLineId = creditLine?.id || 0;

        debit_account_id = debitLine?.account_id || 0;
        credit_account_id = creditLine?.account_id || 0;
        debit_amount = debitLine?.debit_amount || 0;
        credit_amount = creditLine?.credit_amount || 0;

        oldData = {
            title,
            notes,
            reference_id,
            debit_account_id,
            credit_account_id,
            debit_amount,
            credit_amount,
        };

        loading = false;
    };

    load();
</script>

<AppBar>
    <div slot="lead" class="flex gap-2">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li class="crumb">Edit Transaction</li>
        </ol>
    </div>
</AppBar>

{#if loading}
    <Loader />
{:else}
    <TxnForm
        {onSubmit}
        edit={true}
        {title}
        {notes}
        {debit_account_id}
        {credit_account_id}
        {debit_amount}
        {credit_amount}
        {reference_id}
    />
{/if}
