<script lang="ts">
  import type { RootAPI } from "$lib/api";
  import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
  import { AppBar } from "@skeletonlabs/skeleton";
  import { getContext } from "svelte";

  const api = getContext("__api__") as RootAPI;

  let messages: Record<string, any>[] = [];

  const load = async () => {
    const resp = await api.listUserMessages();
    if (resp.status !== 200) {
      return;
    }

    messages = resp.data;
  };

  load();
</script>

<AppBar>
  <svelte:fragment slot="lead">
    <h4 class="h4">Inbox</h4>
  </svelte:fragment>

  <svelte:fragment slot="trail">
    <button class="btn btn-sm variant-filled">
      <SvgIcon className="w-4 h-4" name="ellipsis-horizontal" />
      Clear
    </button>
    <button 
    
    on:click={load}
    class="btn btn-sm variant-filled-secondary">
      <SvgIcon className="w-4 h-4" name="arrow-path" />
      refresh
    </button>



  </svelte:fragment>



</AppBar>

<div class="flex flex-col">
  <div class="overflow-auto">
    <div class="inline-block min-w-full p-2 "> 
        <table class="min-w-full border-gray-300 border"> 
          <thead class="bg-gray-200 border-b">
            <tr>
              <th
                scope="col"
                class="text-sm font-medium text-gray-900 p-2 text-left"
              >
                #
              </th>
              <th
                scope="col"
                class="text-sm font-medium text-gray-900 p-2 text-left"
              >
                Title
              </th>

              <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-2 text-left"
            >
              Type
            </th>


              <th
                scope="col"
                class="text-sm font-medium text-gray-900 p-2 text-left"
              >
                Content
              </th>
              <th
                scope="col"
                class="text-sm font-medium text-gray-900 p-2 text-left"
              >
                
              </th>
            </tr>
          </thead>
          <tbody>
            {#each messages as message}
            <tr
            class="bg-white border-b transition duration-300 ease-in-out hover:bg-gray-100"
          >
            <td
              class="p-2 whitespace-nowrap text-sm font-medium text-gray-900"
              >1</td
            >
            <td
              class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"
            >
              { message.title || ""}
            </td>

            <td
            class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"
          >
          <span class="chip variant-filled-secondary">
            { message.type || ""}
          </span>
            
          </td>

            <td
              class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"
            >
          { message.contents || ""}
            </td>
            <td
              class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"
            >
              
            </td>
          </tr>
            {/each}

          </tbody>
        </table>
        <div class="flex justify-between py-2">
          <button class="btn btn-sm variant-ghost-secondary">
            <SvgIcon className="h-4 w-4" name="arrow-left" />
          </button>

          <button class="btn btn-sm variant-ghost-secondary">
            <SvgIcon className="h-4 w-4" name="arrow-right" />
          </button>

        </div>
    </div>
  </div>
</div>
