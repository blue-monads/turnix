<script lang="ts">
  export let value = "";
  export let onChange = null;
  export let seperator = ",";

  $: _value = value.split(seperator);

  let new_val = "";

  const push = () => {
    if (!onChange) return;
    onChange(value);
  };

  const changeInner = (index) => (ev) => {
    const values = [..._value];
    values[index] = ev.target["value"];
    value = values.toString();
    push();
  };

  const removeInner = (index) => () => {
    const values = _value.filter((_value, idx) => idx !== index);
    value = values.toString();
    push();
  };
</script>

<div class="flex flex-col p-1 gap-1">
  {#each _value as val, idx}
    <div class="flex">
      <input
        type="text"
        class="w-full p-1 border"
        value={val}
        on:change={changeInner(idx)}
      />

      <button
        class="bg-gray-700 rounded text-white p-2"
        on:click={removeInner(idx)}>-</button
      >
    </div>
  {/each}

  <div class="flex mt-1">
    <input type="text" class="p-1 w-full border" bind:value={new_val} />
    <button
      class="bg-gray-700 rounded text-white p-1"
      on:click={() => {
        if (!new_val) return;
        value = value + "," + new_val;
        new_val = "";
        push();
      }}>+</button
    >
  </div>
</div>
