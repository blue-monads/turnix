<script lang="ts">
  import { onMount } from "svelte";
  import { scale } from "svelte/transition";
  interface Props {
    children?: import('svelte').Snippet;
  }

  let { children }: Props = $props();

  let show = $state(false); // menu state
  let menuRef: any = $state(null); // menu wrapper DOM reference

  onMount(() => {
    const handleOutsideClick = (event: Event) => {
      if (show && !menuRef.contains(event.target)) {
        show = false;
      }
    };

    const handleEscape = (event: KeyboardEvent) => {
      if (show && event.key === "Escape") {
        show = false;
      }
    };

    // add events when element is added to the DOM
    document.addEventListener("click", handleOutsideClick, false);
    document.addEventListener("keyup", handleEscape, false);

    // remove events when element is removed from the DOM
    return () => {
      document.removeEventListener("click", handleOutsideClick, false);
      document.removeEventListener("keyup", handleEscape, false);
    };
  });
</script>

<div class="relative" bind:this={menuRef}>
  <div>
    <button
      onclick={() => (show = !show)}
      class="menu focus:outline-none focus:shadow-solid p-2 shadow border rounded"
    >
      Options
    </button>

    {#if show}
      <div
        in:scale={{ duration: 100, start: 0.95 }}
        out:scale={{ duration: 75, start: 0.95 }}
        class="origin-top-right absolute right-0 w-48 py-2 mt-1 bg-white z-50
            rounded shadow-md flex flex-col"
      >
        {@render children?.()}
      </div>
    {/if}
  </div>
</div>
