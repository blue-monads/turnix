<script lang="ts">
  import type { TxnLine } from "./txntype";

  export let lineData: TxnLine[] = [];
  export let accountsIndex: Record<number, string>

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
          <td class="p-0" role="gridcell">{txn.txn.id}</td>
          <td class="p-0" role="gridcell">{txn.txn.title}</td>
          <td class="p-0" role="gridcell">{txn.txn.notes}</td>
          <td class="p-0" role="gridcell"></td>
          <td class="p-0" role="gridcell"></td>
          <td class="p-0" role="gridcell"></td>

          <td class="p-0" role="gridcell"></td>
          <td class="p-0" role="gridcell">
            <button class="underline">edit</button>
          </td>
        </tr>

        {#each txn.lines as line}
          <tr>
            <td class="p-0" role="gridcell"></td>
            <td class="p-0" role="gridcell"></td>
            <td class="p-0" role="gridcell"></td>
            <td class="p-0" role="gridcell">{accountsIndex[line.account_id] || ""}</td>
            <td class="p-0" role="gridcell">
              {#if line.credit_amount}
                {line.credit_amount}
              {/if}
            </td>
            <td class="p-0" role="gridcell">
              {#if line.debit_amount}
                {line.debit_amount}
              {/if}
            </td>
            <td class="p-0" role="gridcell"></td>
            <td class="p-0" role="gridcell"></td>
          </tr>
        {/each}
      {/each}
    </tbody>
  </table>
</div>
