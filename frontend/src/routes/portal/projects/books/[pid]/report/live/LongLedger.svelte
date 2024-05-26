<script lang="ts">
    interface ReportLongLedger {
        txn_id: number;
        title: string;
        account_name: string;
        acc_type: string;
        debit_amount: number;
        credit_amount: number;
        total_debit: number;
        total_credit: number;
        account_id: number;
    }

    interface LedgerGroup {
        account_id: number;
        account_name: string;
        account_type: string;
        total_debit: number;
        total_credit: number;
        lines: ReportLongLedger[];
    }

    export let data: ReportLongLedger[];

    const format = (rawData: ReportLongLedger[]) => {
        let groupedLedger: Record<number, LedgerGroup> = {};

        rawData.forEach((item) => {
            let old = groupedLedger[item.account_id];
            if (!old) {
                old = {
                    account_id: item.account_id,
                    account_name: item.account_name,
                    account_type: item.acc_type,
                    lines: [],
                    total_credit: item.total_credit,
                    total_debit: item.total_debit,
                };

                groupedLedger[item.account_id] = old;
            }

            old.lines.push(item);
        });

        console.log("@after_format", groupedLedger);

        return groupedLedger;
    };

    $: grouped = format(data);

    $: {
        console.log("@data/m", data);
        grouped = format(data);
        console.log("@grouped", grouped);
    }
</script>

<div class="table-container h-[calc(100vh-14rem)] overflow-y-auto">
    <table class="table table-hover h-fit">
        <thead>
            <tr>
                <th>Txn</th>
                <th>Title</th>

                <th>Debit Amount</th>
                <th>Credit Amount</th>
            </tr>
        </thead>
        <tbody>
            {#each Object.values(grouped) as group, i}
                <tr>
                    <td colspan="5">
                        <span class="font-semibold text-lg"
                            >{group.account_name}</span
                        >
                    </td>
                </tr>

                {#each group.lines as row}
                    <tr>
                        <td>{row.txn_id}</td>
                        <td>{row.title}</td>

                        <td>{row.debit_amount}</td>
                        <td>{row.credit_amount}</td>
                    </tr>
                {/each}

                <tr>
                    <th colspan="2">Total</th>
                    <td>{group.total_debit}</td>
                    <td>{group.total_credit}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
