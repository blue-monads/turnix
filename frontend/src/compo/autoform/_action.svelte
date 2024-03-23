<script lang="ts">
  export let onClick: Function;
  export let name = "";
  let loading = false;

  const handler = async () => {
    loading = true;
    const r = onClick();
    if (r instanceof Promise) {
      await r;
    }
    loading = false;
  };
</script>

<button
  on:click={handler}
  class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-blue-400"
>
  <svg
    class="h-5 w-5 {loading ? 'animate-spin' : ''}"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    ><path
      d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"
    /><polyline points="17 21 17 13 7 13 7 21" /><polyline
      points="7 3 7 8 15 8"
    /></svg
  >

  {name}
</button>
