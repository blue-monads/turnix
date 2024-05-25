<script lang="ts">
  export let name = "";
  export let acc_type = "revenue";
  export let info = "";
  export let edit = false;

  export let onChange: Function;

  const ACCOUNT_TYPES = {
    expenses: "Expenses",
    revenue: "Revenue",
    assets: "Assets",
    liabilities: "Liabilities",
    equity: "Equity",
  };
</script>

<form
  class="p-2"
  on:submit|preventDefault={() => {
    onChange({
      name,
      acc_type,
      info,
    });
  }}
>
  <div class="card">
    <header class="card-header">
      <h4 class="h4">
        {#if edit}
          Edit Account
        {:else}
          Add Account
        {/if}
      </h4>
    </header>

    <section class="p-4 flex flex-col gap-4">
      <label class="label">
        <span>Name</span>
        <input
          bind:value={name}
          class="input p-1"
          type="text"
          placeholder="Input"
        />
      </label>

      <label class="label">
        <span>Type</span>
        <select bind:value={acc_type} class="select">
          {#each Object.entries(ACCOUNT_TYPES) as acc}
            <option value={acc[0]}>{acc[1]}</option>
          {/each}
        </select>
      </label>

      <label class="label">
        <span>Info</span>
        <textarea
          bind:value={info}
          class="textarea p-1"
          rows="4"
          placeholder={"information about account"}
        />
      </label>
    </section>
    <footer class="card-footer flex justify-end">
      <button type="submit" class="btn variant-filled"> save </button>
    </footer>
  </div>
</form>
