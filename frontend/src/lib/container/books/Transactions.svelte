<script lang="ts">
  import { formatCurrency } from "$lib/utils";
  import type { TxnLine } from "./txntype";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import { createEventDispatcher, getContext, onMount } from "svelte";

  const dispatcher = createEventDispatcher();


  interface Props {
    lineData?: TxnLine[];
    accountsIndex: Record<number, string>;
    pid: string;
  }

  let { lineData = [], accountsIndex, pid }: Props = $props();

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
          <td role="gridcell" />
          <td role="gridcell" />
          <td role="gridcell" />

          <td role="gridcell" />
          <td role="gridcell">
            <a
              class="underline"
              href={`/z/pages/portal/projects/books/${pid}/txn/edit?tid=${txn.txn.id}`}
            >
              edit
            </a>
          </td>
        </tr>

        {#each txn.lines as line}
          <tr>
            <td role="gridcell" />
            <td role="gridcell" />
            <td role="gridcell" />
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
            <td role="gridcell" />
            <td role="gridcell" />
          </tr>
        {/each}
      {/each}
    </tbody>
  </table>

  <div class="flex justify-end gap-2 p-1">
    <div class="flex gap-2">
      <button
        class="btn btn-sm bg-gray-100"

        
        onclick={() => {
          console.log("@prev");
          dispatcher("prev_page");
        }}
      >
        <SvgIcon className="h-4 w-4" name="chevron-left" />
      </button>

      <button
        class="btn btn-sm bg-gray-100"
        onclick={() => {
          console.log("@next");
          dispatcher("next_page");
        }}
      >
        <SvgIcon className="h-4 w-4" name="chevron-right" />
      </button>
    </div>
  </div>
</div>
