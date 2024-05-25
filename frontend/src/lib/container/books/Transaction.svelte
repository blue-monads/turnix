<script lang="ts">
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import { FileDropzone } from "@skeletonlabs/skeleton";

  let files: FileList;

  export let edit = false;

  export let title = "";
  export let notes = "";
  export let debit_account_id = 0;
  export let credit_account_id = 0;
  export let debit_amount = 100;
  export let credit_amount = 100;
  export let reference_id = "";
  export let attachments: Record<string, string> = {};

  export let onSubmit: Function;
  export let onPickAccount: Function;

  export let accountsIndex: Record<string, string> = {};

  const pickDebitAccount = async () => {
    const nextaccount = await onPickAccount();
    if (nextaccount) {
      debit_account_id = nextaccount;
    }
  };

  const pickCreditAccount = async () => {
    const nextaccount = await onPickAccount();
    if (nextaccount) {
      credit_account_id = nextaccount;
    }
  };
</script>

<div class="p-2">
  <form
    class="card"
    on:submit|preventDefault={() => {
      onSubmit({
        title,
        notes,
        debit_account_id,
        credit_account_id,
        debit_amount,
        credit_amount,
        reference_id,
        attachments,
      });
    }}
  >
    <header class="card-header">
      <h4 class="h4">
        {#if edit}
          Edit Transaction
        {:else}
          Add Transaction
        {/if}
      </h4>
    </header>

    <section class="p-4 flex flex-col gap-4">
      <label class="label">
        <span>Title</span>
        <input
          name="title"
          bind:value={title}
          class="input p-1"
          type="text"
          placeholder="Opening balance"
        />
      </label>

      <div class="label border p-2 flex flex-col">
        <span>Lines</span>
        <div class="flex flex-row justify-start gap-4 items-center">
          <label class="label">
            <span>Debit Account</span>
            <input
              value={accountsIndex[debit_account_id] || ""}
              class="input p-1"
              type="text"
              disabled
              placeholder="Petty cash"
            />
          </label>

          <button
            class="pt-4 cursor-pointer"
            type="button"
            on:click={pickDebitAccount}
          >
            <SvgIcon name="arrow-up-right" className="w-5 h-5" />
          </button>

          <label class="label">
            <span>Amount</span>
            <input
              name="debit amount"
              bind:value={debit_amount}
              class="input p-1"
              type="number"
            />
          </label>
        </div>

        <div class="flex flex-row justify-start gap-4">
          <label class="label">
            <span>Credit Account</span>
            <input
              class="input p-1"
              value={accountsIndex[credit_account_id] || ""}
              type="text"
              disabled
              placeholder="Bank XYZ"
            />
          </label>

          <button
            class="pt-4 cursor-pointer"
            type="button"
            on:click={pickCreditAccount}
          >
            <SvgIcon name="arrow-up-right" className="w-5 h-5" />
          </button>

          <label class="label">
            <span>Amount</span>
            <input
              name="credit amount"
              bind:value={credit_amount}
              class="input p-1"
              type="number"
            />
          </label>
        </div>
      </div>

      <label class="label">
        <span>Notes</span>
        <textarea
          bind:value={notes}
          class="textarea p-1"
          rows="4"
          placeholder={"Extra information about txn"}
        />
      </label>

      <!-- svelte-ignore a11y-label-has-associated-control -->
      <label class="label">
        <span>Attachments</span>
        <FileDropzone bind:files name="files">
          <div slot="lead" class="flex justify-center">
            <SvgIcon name="arrow-up-tray" className="w-6 h-6" />
          </div>

          <svelte:fragment slot="message">
            <strong>Upload a file</strong> or drag and drop
          </svelte:fragment>
          <svelte:fragment slot="meta">PNG, JPG or PDF</svelte:fragment>
        </FileDropzone>
      </label>

      <label class="label">
        <span>Reference ID</span>
        <input class="input p-1" type="text" placeholder="REF-123" />
      </label>
    </section>
    <footer class="card-footer flex justify-end">
      <button type="submit" class="btn variant-filled"> save </button>
    </footer>
  </form>
</div>
