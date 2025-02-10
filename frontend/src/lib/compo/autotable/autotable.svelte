<script lang="ts">
  import { onMount } from 'svelte';
  import SvgIcon from "../icons/SvgIcon.svelte";

  interface Props {
    actions?: {
      Action: (id: any, data: any) => Promise<void> | void;
      Class?: string;
      icon?: string;
      drop?: boolean;
      Name: string;
    }[];
    key_names?: any[];
    datas?: any[];
    action_key?: string;
    color?: any[];
    show_drop?: boolean;
    hashSeed?: number;
  }

  let {
    actions = [],
    key_names = [],
    datas = [],
    action_key = "",
    color = [],
    show_drop = false,
    hashSeed = 65
  }: Props = $props();

  let extern_actions: any[] = $state([]);
  let drop_actions: any[] = $state([]);
  let activeDropdownRowIndex: number | null = $state(null);
  let dropdownTriggers: HTMLButtonElement[] = [];
  let dropdownPositions: {top: number, left: number}[] = $state([]);

  if (!show_drop) {
    extern_actions = actions;
  } else {
    extern_actions = actions.filter((v: any) => !v["drop"]);
    drop_actions = actions.filter((v: any) => !!v["drop"]);
  }

  const hashCode = (str: string) => {
    let hash = hashSeed
    for (var i = 0; i < str.length; i++) {
      hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return hash;
  };

  const color_it = (key: string, str: string) => {
    if (!color.includes(key)) {
      return "";
    }
    return `background: hsl(${hashCode(str) % 360}, 100%, 80%)`;
  };

  const toggleDropdown = (index: number) => {
    if (activeDropdownRowIndex === index) {
      activeDropdownRowIndex = null;
      return;
    }
    
    activeDropdownRowIndex = index;
    
    // Calculate dropdown position
    if (dropdownTriggers[index]) {
      const rect = dropdownTriggers[index].getBoundingClientRect();
      dropdownPositions[index] = {
        top: rect.bottom + window.scrollY,
        left: rect.left + window.scrollX
      };
    }
  };

  // Close dropdown when clicking outside
  const handleClickOutside = (event: MouseEvent) => {
    if (activeDropdownRowIndex !== null) {
      const dropdown = event.target as HTMLElement;
      if (!dropdown.closest('.dropdown-container')) {
        activeDropdownRowIndex = null;
      }
    }
  };

  onMount(() => {
    document.addEventListener('click', handleClickOutside);
    return () => {
      document.removeEventListener('click', handleClickOutside);
    };
  });
</script>

<div class="p-2 overflow-auto">
  <table class="table-auto border-collapse w-full card">
    <thead>
      <tr
        class="rounded-lg text-sm font-medium text-gray-700 text-left"
        style="font-size: 0.9674rem"
      >
        {#each key_names as key_name}
          <th class="px-2 py-1" style="background-color:#f8f8f8"
            >{key_name[1]}</th
          >
        {/each}
        <th class="px-2 py-2" style="background-color:#f8f8f8"> Actions </th>
      </tr>
    </thead>
    <tbody class="text-sm font-normal text-gray-700">
      {#each datas as data, rowIndex}
        <tr
          class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
        >
          {#each key_names as key_name}
            <td class="px-3 py-1">
              <span
                class="p-1 rounded-lg"
                style={color_it(key_name[0], data[key_name[0]] || "")}
                >{data[key_name[0]] || ""}</span
              >
            </td>
          {/each}

          <td class="px-3 py-1 dropdown-container">
            <div class="flex flex-row items-center">
              {#each extern_actions as action}
                {@const icon = action["icon"]}
                <button
                  onclick={() => action.Action(data[action_key], data)}
                  class="flex m-1 text-white transform hover:scale-110 btn btn-sm {action.Class ||
                    'bg-blue-400'}"
                >
                  {#if icon}
                    <SvgIcon name={icon} className="h-5 w-5" />
                  {/if}

                  {action.Name}</button
                >
              {/each}

              {#if show_drop}
                <div>
                  <button 
                    bind:this={dropdownTriggers[rowIndex]}
                    class="btn btn-sm bg-blue-400 flex items-center" 
                    onclick={() => toggleDropdown(rowIndex)}
                  >
                    {#if activeDropdownRowIndex === rowIndex}
                      <SvgIcon name="chevron-up" className="h-5 w-5 mr-1" />
                    {:else}
                      <SvgIcon name="chevron-down" className="h-5 w-5 mr-1" />
                    {/if}
                    options
                  </button>

                  {#if activeDropdownRowIndex === rowIndex && dropdownPositions[rowIndex]}
                    <div 
                      class="fixed z-50 bg-white border rounded-md shadow-lg" 
                      style="top: {dropdownPositions[rowIndex].top}px; left: {dropdownPositions[rowIndex].left}px;"
                    >
                      {#each drop_actions as action}
                        <button
                          onclick={() => {
                            action.Action(data[action_key], data);
                            activeDropdownRowIndex = null;
                          }}
                          class="flex w-full items-center px-4 py-2 text-sm text-gray-700 hover:bg-blue-500 hover:text-white"
                        >
                          {#if action["icon"]}
                            <SvgIcon name={action["icon"]} className="h-5 w-5 mr-2" />
                          {/if}
                          <span>{action.Name}</span>
                        </button>
                      {/each}
                    </div>
                  {/if}
                </div>
              {/if}
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>