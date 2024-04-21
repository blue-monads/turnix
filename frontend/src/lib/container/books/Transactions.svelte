<script lang="ts">
  import { formatCurrency } from "$lib/utils";
  import type { TxnLine } from "./txntype";

  export let lineData: TxnLine[] = [];
  export let accountsIndex: Record<number, string>;

  console.log("@lines", lineData);
</script>

<div class="card m-2 overflow-x-auto">
  <table class="table p-0 table-compact">
    <thead class="table-head">
      <tr>
        <th role="columnheader">#</th>
        <th role="columnheader">Title</th>
        <th role="columnheader">Notes</th>

        <th role="columnheader">Account</th>
        <th role="columnheader">Debit</th>
        <th role="columnheader">Credit</th>

        <th role="columnheader">Attachments</th>
        <th role="columnheader">Actions</th>
      </tr>
    </thead>
    <tbody class="table-body p-0 overflow-y-scroll w-full max-h-screen">
      {#each lineData as txn}
        <tr>
          <td role="gridcell">{txn.txn.id}</td>
          <td role="gridcell">{txn.txn.title}</td>
          <td role="gridcell">{txn.txn.notes}</td>
          <td role="gridcell"></td>
          <td role="gridcell"></td>
          <td role="gridcell"></td>

          <td role="gridcell"></td>
          <td role="gridcell">
            <button class="underline">edit</button>
          </td>
        </tr>

        {#each txn.lines as line}
          <tr>
            <td role="gridcell"></td>
            <td role="gridcell"></td>
            <td role="gridcell"></td>
            <td role="gridcell">
              <span class="chip variant-filled">
                {accountsIndex[line.account_id] || ""}</span
              >
            </td>
            <td role="gridcell">
              {#if line.credit_amount}
                <span class="chip variant-ghost">
                  {formatCurrency(line.credit_amount)}
                </span>
              {/if}
            </td>
            <td role="gridcell">
              {#if line.debit_amount}
                <span class="chip variant-ghost">
                  {formatCurrency(line.debit_amount)}
                </span>
              {/if}
            </td>
            <td role="gridcell"></td>
            <td role="gridcell"></td>
          </tr>
        {/each}
      {/each}
    </tbody>
  </table>
</div>
