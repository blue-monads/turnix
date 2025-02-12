<script lang="ts">
  import { RootAPI } from "$lib/api";
  import { getContext } from "svelte";
  import {
    AppBar,
    getModalStore,
    type ModalSettings,
  } from "@skeletonlabs/skeleton";

  const api = getContext("__api__") as RootAPI;

  const store = getModalStore();
</script>

<svelte:head>
  <title>Projects</title>
</svelte:head>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4">Home</h4>
  </svelte:fragment>
</AppBar>

<div class="flex flex-col gap-8 mt-10">
  <div class="flex min-h-24 p-4">
    <div class="flex items-center max-w-md mx-auto bg-white rounded-lg">
      <div class="w-full">
        <input
          type="search"
          class="w-full px-4 py-1 text-gray-800 rounded-full focus:outline-none"
          placeholder="search"
        />
      </div>
      <div>
        <!-- svelte-ignore a11y_consider_explicit_label -->
        <button
        
          class="flex items-center btn variant-filled justify-center w-12 h-12 text-white rounded-r-lg"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
        </button>
      </div>
    </div>
  </div>

  <div
    class="flex items-center justify-center min-h-24 flex-col gap-4 variant-glass"
  >
    <h2 class="h2">Welcome to the Portal</h2>
    <p class="p text-xl">
      Browse your projects or create a new one to get started.
    </p>
  </div>

  <div class="flex justify-around">
    <a href="/z/pages/portal/project" class="btn variant-filled">Explore</a>
  </div>
</div>
