<script lang="ts">
  import { run } from 'svelte/legacy';

  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import { NewBookAPI } from "$lib/projects/books";
  import { getContext } from "svelte";
  import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
  import type { RootAPI } from "$lib/api";
  import { page } from "$app/stores";
  import LongLedger from "./LongLedger.svelte";
  import ShortLedger from "./ShortLedger.svelte";

  const pid = $page.params["pid"];

  const store = getModalStore();
  const api = NewBookAPI(getContext("__api__") as RootAPI);

  let generating = $state(false);
  let template_id = 0;
  let report_type = $state("short_ledger");

  let current_report_type = $state(report_type);
  let data: any = $state();

  const load = async () => {};

  load();

  const onGenerate = async () => {
    generating = true;

    current_report_type = report_type;
    const resp = await api.generateLiveTxn(pid, {
      report_type,
      template_id,
    });

    if (resp.status !== 200) {
      generating = false;

      return;
    }

    data = resp.data;
    console.log("@data", data);
    generating = false;
  };

  let report_type_invalid = $derived(report_type === "custom" && template_id === 0);

  run(() => {
    console.log("@data/2", data);
  });
</script>

<AppBar>
  <svelte:fragment slot="lead">
    <ol class="breadcrumb">
      <li class="crumb">
        <a class="anchor" href="/z/pages/portal/projects/books">Books</a>
      </li>
      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li>Reports</li>
      <li class="crumb-separator" aria-hidden>&rsaquo;</li>
      <li>Live</li>
    </ol>
  </svelte:fragment>
</AppBar>

<div class="flex flex-col p-2 gap-2">
  <div class="card p-2">
    <header class="card-header">
      <h4 class="h4">Live Reports</h4>
    </header>

    <section
      class="pl-4 py-2 flex justify-between items-center flex-wrap gap-2"
    >
      <div class="flex gap-2 items-center flex-wrap">
        <label class="label">
          <span>Report Type</span>

          <select bind:value={report_type} class="select">
            <option value="short_ledger">Ledger short</option>
            <option value="long_ledger">Ledger long</option>
            <option value="custom">Custom</option>
          </select>
        </label>

        <label class="label">
          <span>From Date</span>
          <input class="input" type="datetime-local" />
        </label>

        <label class="label">
          <span>To Date</span>
          <input class="input" type="datetime-local" />
        </label>
      </div>

      <div>
        <button
          disabled={generating || report_type_invalid}
          onclick={onGenerate}
          class="btn btn-sm variant-filled disabled:variant-ghost disabled:text-slate-800"
        >
          <SvgIcon className="h-4 w-4" name="bolt" />

          Generate
        </button>
      </div>
    </section>
  </div>

  <div class="p-2 card">
    {#if data}
      {#if current_report_type === "long_ledger"}
        {#key data}
          <LongLedger {data} />
        {/key}
      {:else if current_report_type === "short_ledger"}
        {#key data}
          <ShortLedger {data} />
        {/key}
      {/if}
    {/if}
  </div>
</div>
