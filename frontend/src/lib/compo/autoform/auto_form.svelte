<script lang="ts">
  import KvEditor from "./_kv_editor.svelte";
  import type { Schema } from "./form";
  import MultiText from "./_multi_text.svelte";
  import Action from "./_action.svelte";

  import { createEventDispatcher } from "svelte";
  import SvgIcon from "../icons/SvgIcon.svelte";

  export let schema: Schema;
  export let data: Record<string, any> = {};
  export let modified = false;
  export let message = "";
  export let onSave: ((data: any) => Promise<void>) | undefined = undefined;
  export let noSubmit = false;

  const dispatch = createEventDispatcher();

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

    dispatch("modified", mod_data);
  };

  const newSlug = (field: string) => {
    const s = "fixme";
    mod_data = { ...mod_data, [field]: s };
    return s;
  };
</script>

<div class="h-full w-full p-4 overflow-auto">
  <div class="p-5 w-full card">
    <h3 class="h3">{schema.name}</h3>
    <p class="text-red-500">{message || ""}</p>

    {#each schema.fields as field, idx}
      <div class="flex-col flex py-3">
        <label for={`field-${idx}`} class="pb-2 label">{field.name}</label>

        {#if field.ftype === "TEXT"}
          <input
            id="field-{idx}"
            type="text"
            list="field-{idx}-datalist"
            value={get(field.key_name)}
            on:change={set(field.key_name)}
            disabled={field.disabled}
            class="p-2 input"
          />

          <datalist id="field-{idx}-datalist">
            {#each field.options || [] as opt}
              <option value={opt}>{opt}</option>
            {/each}
          </datalist>
        {:else if field.ftype === "EMAIL"}
          <input
            id="field-{idx}"
            type="email"
            list="field-{idx}-datalist"
            value={get(field.key_name)}
            on:change={set(field.key_name)}
            disabled={field.disabled}
            class="p-2 input"
          />
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
                class="p-2 input"
              />

              <datalist id="field-{idx}-datalist">
                {#each field.options || [] as opt}
                  <option value={opt}>{opt}</option>
                {/each}
              </datalist>
            {:else}
              <select
                class="p-1 input"
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
                ><SvgIcon
                  name={_open_selects[idx] ? "lock-open" : "lock-closed"}
                  className="w-6 h-6"
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
            class="p-2 textarea"
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
          <div class="flex w-full justify-start">
            <input
              type="checkbox"
              id={`field-${idx}`}
              value={get(field.key_name) || false}
              on:change={setBool(field.key_name)}
              class="p-2"
            />
          </div>
        {:else if field.ftype === "KEY_VALUE_TEXT"}
          <KvEditor
            data={data[field.key_name] || {}}
            onChange={(data) => {
              setValue(field.key_name)(JSON.stringify(data));
            }}
          />
        {:else}
          <div>Not impl</div>
        {/if}
      </div>
    {/each}

    {#if modified && !noSubmit}
      <div class="flex justify-end py-3">
        <Action
          name="Save"
          onClick={() => {
            onSave && onSave(mod_data);
          }}
        />
      </div>
    {/if}
  </div>
</div>
