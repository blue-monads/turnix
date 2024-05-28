<script lang="ts">
  export let modified = false;
  export let data = {};
  export let onChange: Function | undefined = undefined;
  export let sensitive = false;

  $: _data = { ...(typeof data === "string" ? JSON.parse(data) : data) };

  export const getData = () => ({ ..._data });

  let current_active_key = "";
  let new_key = "";
  let new_value = "";
  let run = false;

  const value_set = (key: string, newvalue: any) => {
    modified = true;
    run = false;
    _data = { ..._data, [key]: newvalue };
  };

  $: if (modified && onChange && !run) {
    onChange(_data);
    run = true;
  }
</script>

<div class="border p-2 shadow rounded">
  <table class="w-full text-sm text-left text-gray-500">
    <thead class="text-xs text-gray-700 uppercase bg-gray-50">
      <tr>
        <th scope="col" class="px-6 py-3"> Key </th>
        <th scope="col" class="px-6 py-3"> Value </th>
        <th scope="col" class="px-6 py-3">
          <span class="sr-only">delete</span>
        </th>
      </tr>
    </thead>
    <tbody>
      {#key _data}
        {#each Object.entries(_data) as [key, val]}
          <tr class="bg-white border-b hover:bg-gray-50">
            <th
              scope="row"
              class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
            >
              <div>
                {key}
              </div>
            </th>
            <td class="px-6 py-4">
              {#if current_active_key === key}
                <input
                  type="text"
                  class="border border-slate-500 rounded-sm w-full"
                  value={val + ""}
                  on:change={(ev) => value_set(key, ev.target["value"])}
                />
              {:else}
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->

                <input
                  value={val}
                  type={sensitive ? "password" : "text"}
                  disabled={true}
                />
              {/if}
            </td>
            <td class="px-6 py-4 text-right">
              {#if current_active_key !== key}
                <button
                  on:click={() => {
                    current_active_key = key;
                  }}
                  class="font-medium text-blue-600 hover:underline"
                  >edit
                </button>
              {/if}
              <button
                class="font-medium text-blue-600 hover:underline"
                on:click={() => {
                  delete _data[key];
                  _data = { ..._data };
                  modified = true;
                  run = false;
                }}
              >
                delete
              </button>
            </td>
          </tr>
        {/each}

        <tr class="bg-gray-50 border-b hover:bg-gray-100">
          <th scope="row" class="font-medium text-gray-900 whitespace-nowrap">
            <input
              type="text"
              class="border border-slate-500 rounded-sm w-full"
              bind:value={new_key}
            />
          </th>
          <td class="">
            <input
              type="text"
              class="border border-slate-500 rounded-sm w-full"
              bind:value={new_value}
            />
          </td>
          <td class="text-right">
            <button
              class="font-medium text-blue-600 hover:underline"
              on:click={() => {
                if (!new_key) {
                  return;
                }

                modified = true;
                run = false;
                _data = { ..._data, [new_key]: new_value };
                new_key = "";
                new_value = "";
              }}
            >
              add
            </button>
          </td>
        </tr>
      {/key}
    </tbody>
  </table>
</div>
