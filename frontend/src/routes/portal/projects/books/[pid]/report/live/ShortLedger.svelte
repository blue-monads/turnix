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
