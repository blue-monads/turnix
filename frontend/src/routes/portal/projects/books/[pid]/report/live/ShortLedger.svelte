<script lang="ts">
    interface ReportLedgerShort {
        account_id: number;
        account_name: string;
        total_debit: number;
        total_credit: number;
        acc_type: string;
    }

    export let data: ReportLedgerShort[];
</script>

<div class="table-container h-[calc(100vh-14rem)] overflow-y-auto">
    <table class="table table-hover h-fit">
        <thead>
            <tr>
                <th>Id</th>
                <th>Account</th>
                <th>Account Type</th>
                <th>Debit Total</th>
                <th>Credit Total</th>
            </tr>
        </thead>
        <tbody>
            {#each data as item, i}
                <tr>
                    <th>{item.account_id}</th>
                    <td>{item.account_name}</td>
                    <td>{item.acc_type}</td>
                    <td>{item.total_debit}</td>
                    <td>{item.total_credit}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
