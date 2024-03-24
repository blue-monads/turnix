<script lang="ts">
  import KvEditor from "./_kv_editor.svelte";
  import type { Schema } from "./form";
  import MultiText from "./_multi_text.svelte";
  import Action from "./_action.svelte";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let schema: Schema;
  export let data: Record<string, any> = {};
  export let modified = false;
  export let message = "";
  export let onSave: (data: any) => Promise<void>;

  let mod_data = {};
  $: _open_selects = {};

  $: console.log(
    `FORM_DEBUG => ${schema.name}`,
    "DATA",
    mod_data,
    "ORIGINAL_DATA",
    data,
    "SCHEMA",
    schema,
  );

  const get = (name: string): any => data[name] || "";
  const set = (name: string) => (ev: any) => {
    setValue(name)(ev.target.value);
  };

  const setNumber = (name: string) => (ev: any) => {
    setValue(name)(Number(ev.target.value));
  };

  const setBool = (name: string) => (ev: any) => {
    setValue(name)(Boolean(ev.target.checked));
  };

  const setValue = (name: string) => (val: any) => {
    mod_data = { ...mod_data, [name]: val };
    modified = true;
  };

  const newSlug = (field: string) => {
    const s = "fixme";
    mod_data = { ...mod_data, [field]: s };
    return s;
  };
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full">
    <div class="text-2xl text-indigo-900">{schema.name}</div>
    <p class="text-red-500">{message || ""}</p>

    {#each schema.fields as field, idx}
      <div class="flex-col flex py-3">
        <label for={`field-${idx}`} class="pb-2 text-gray-700 font-semibold"
          >{field.name}</label
        >

        {#if field.ftype === "TEXT"}
          <input
            id="field-{idx}"
            type="text"
            list="field-{idx}-datalist"
            value={get(field.key_name)}
            on:change={set(field.key_name)}
            disabled={field.disabled}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />

          <datalist id="field-{idx}-datalist">
            {#each field.options || [] as opt}
              <option value={opt}>{opt}</option>
            {/each}
          </datalist>
        {:else if field.ftype === "SELECT"}
          <div class="flex justify-between w-full">
            {#if _open_selects[idx]}
              <input
                id="field-{idx}"
                type="text"
                list="field-{idx}-datalist"
                value={get(field.key_name)}
                on:change={set(field.key_name)}
                disabled={field.disabled}
                class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200 w-full"
              />

              <datalist id="field-{idx}-datalist">
                {#each field.options || [] as opt}
                  <option value={opt}>{opt}</option>
                {/each}
              </datalist>
            {:else}
              <select
                class="p-1 rounded border w-full"
                id="field-{idx}"
                value={get(field.key_name)}
                on:change={set(field.key_name)}
              >
                {#each field.options || [] as opt}
                  <option value={opt}>{opt}</option>
                {/each}
              </select>
            {/if}

            <div class="w-10 p-1 text-gray-700">
              <button
                on:click={() => {
                  _open_selects[idx] = !_open_selects[idx];
                  _open_selects = _open_selects;
                }}
                ><Icon
                  name={_open_selects[idx] ? "lock-open" : "lock-closed"}
                  class="w-6 h-6"
                /></button
              >
            </div>
          </div>
        {:else if field.ftype === "TEXT_SLUG"}
          <input
            id="field-{idx}"
            type="text"
            list="field-{idx}-datalist"
            value={get(field.key_name) ||
              (field["slug_gen"] ? field.slug_gen() : newSlug(field.key_name))}
            on:change={set(field.key_name)}
            disabled={field.disabled}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        {:else if field.ftype === "MULTI_TEXT"}
          <MultiText
            onChange={setValue(field.key_name)}
            value={get(field.key_name)}
          />
        {:else if field.ftype === "LONG_TEXT" || field.ftype === "TEXT_POLICY"}
          <textarea
            id={`field-${idx}`}
            value={get(field.key_name)}
            on:change={set(field.key_name)}
            disabled={field.disabled}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        {:else if field.ftype === "INT"}
          <input
            type="number"
            id={`field-${idx}`}
            value={get(field.key_name)}
            on:change={setNumber(field.key_name)}
            disabled={field.disabled}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        {:else if field.ftype === "BOOL"}
          <input
            type="checkbox"
            id={`field-${idx}`}
            value={get(field.key_name) || false}
            on:change={setBool(field.key_name)}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        {:else if field.ftype === "KEY_VALUE_TEXT"}
          <KvEditor
            data={data[field.key_name] || {}}
            onChange={setValue(field.key_name)}
          />
        {:else}
          <div>Not impl</div>
        {/if}
      </div>
    {/each}

    {#if modified}
      <div class="flex justify-end py-3">
        <Action name="Save" onClick={() => onSave(mod_data)} />
      </div>
    {/if}
  </div>
</div>
