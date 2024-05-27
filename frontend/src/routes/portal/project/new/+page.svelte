<script lang="ts">
  import { gotoPorjects } from "$lib/nav";
  import { AutoForm, PageLayout, AutoTable } from "$lib/compo";
  import { params } from "$lib/params";
  import { getContext } from "svelte";
  import Loader from "$lib/compo/loader/loader.svelte";
  import type { Schema, Field } from "$lib/compo/autoform/form";
  import type { RootAPI } from "$lib/api";

  const api = getContext("__api__") as RootAPI;
  let loading = true;

  let message = "";

  let schemaFields: Field[] = [];
  const load = async () => {
    const resp = await api.getProjectTypeForm($params["ptype"]);
    if (resp.status !== 200) {
      return;
    }

    schemaFields = resp.data || [];
    loading = false;
  };

  load();
</script>

{#if loading}
  <Loader />
{:else}
  <div class="p-2">
    <div class="card p-2">
      <header class="flex justify-center">
        <div class="flex flex-col gap-2">
          <h3 class="h3">New Project</h3>
        </div>
      </header>

      <section class="flex flex-col gap-4">
        <label class="label">
          <span>Name</span>
          <input class="input p-1" type="text" placeholder="TestProject1" />
        </label>

        <label class="label">
          <span>Project Type</span>
          <select class="select" disabled>
            <option value={$params["ptype"]}>{$params["ptype"]}</option>
          </select>
        </label>

        <label class="label">
          <span>Info</span>
          <textarea
            class="textarea"
            rows="4"
            placeholder="This is a project description"
          />
        </label>
      </section>

      <section class="mt-2">
        <AutoForm
          data={{}}
          noSubmit={true}
          {message}
          schema={{
            fields: schemaFields,
            name: "",
            required_fields: ["name"],
          }}
        />
      </section>

      <footer class="flex justify-end p-2 gap-2">
        <button class="btn variant-filled" > Update </button>
      </footer>
    </div>
  </div>
{/if}
