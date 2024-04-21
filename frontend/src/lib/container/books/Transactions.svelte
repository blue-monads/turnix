<script lang="ts">
  const sourceData = [
    {
      id: 1,
      title: "Open petty cash",
      notes: "New petty cash account bashed on @johns suggestion",
      links: "",
      lines: [
        {
          account_id: 12,
          debit_amount: 2000,
          credit_amount: 0,
        },
        {
          account_id: 19,
          credit_amount: 2000,
          debit_amount: 0,
        },
      ],
    },
    {
      id: 2,
      title: "Open petty cash",
      notes: "New petty cash account bashed on @johns suggestion",
      links: "",
      lines: [
        {
          account_id: 12,
          debit_amount: 2000,
          credit_amount: 0,
        },
        {
          account_id: 19,
          credit_amount: 2000,
          debit_amount: 0,
        },
      ],
    },
    {
      id: 3,
      title: "Open petty cash",
      notes: "New petty cash account bashed on @johns suggestion",
      links: "",
      lines: [
        {
          account_id: 12,
          debit_amount: 2000,
          credit_amount: 0,
        },
        {
          account_id: 19,
          credit_amount: 2000,
          debit_amount: 0,
        },
      ],
    },
  ];
</script>

<div class="card m-2">
  <table class="table p-0">
    <thead class="table-head">
      <tr>
        <th role="columnheader">#</th>
        <th role="columnheader">Title</th>
        <th role="columnheader">Notes</th>

        <th role="columnheader">Account</th>
        <th role="columnheader">Debit</th>
        <th role="columnheader">Credit</th>
        <th role="columnheader">Links</th>
        <th role="columnheader">Attachments</th>
        <th role="columnheader">Actions</th>
      </tr>
    </thead>
    <tbody class="table-body p-0">
      {#each sourceData as txn}
        <tr>
          <td class="p-0" role="gridcell">{txn.id}</td>
          <td class="p-0" role="gridcell">{txn.title}</td>
          <td class="p-0" role="gridcell">{txn.notes}</td>
          <td class="p-0" role="gridcell"></td>
          <td class="p-0" role="gridcell"></td>
          <td class="p-0" role="gridcell"></td>
          <td class="p-0" role="gridcell">{txn.links}</td>
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
            <td class="p-0" role="gridcell">{line.account_id}</td>
            <td class="p-0" role="gridcell">{line.credit_amount}</td>
            <td class="p-0" role="gridcell">{line.debit_amount}</td>
            <td class="p-0" role="gridcell"></td>
            <td class="p-0" role="gridcell"></td>
            <td class="p-0" role="gridcell"></td>
          </tr>
        {/each}
      {/each}
    </tbody>
  </table>
</div>
