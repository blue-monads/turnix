<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";

    export let codeValue = `const handle = (ctx) => {

    }`;




</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books"
                    >Projects</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Hooks</li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>New</li>
        </ol>
    </svelte:fragment>
</AppBar>

<div class="flex p-2 flex-col gap-2">
    <div class="card p-4 flex flex-col gap-2">
        <header class="card-header">
            <h4 class="h4">New Hook</h4>
        </header>

        <section class="p-4">
            <label class="label">
                <span>Name</span>
                <input class="input" type="text" placeholder="Input" />
            </label>

            <label class="label">
                <span>Event Type</span>
                <select class="select">
                    <option value="1">after_new_record</option>
                    <option value="2">before_new_record</option>
                </select>
            </label>

            <label class="label">
                <span>Hook Type</span>
                <select class="select">
                    <option value="script">Script</option>
                    <option value="webhook">Webhook</option>
                </select>
            </label>

            <label class="label">
                <span>Script Code</span>
                <textarea bind:value={codeValue} class="textarea" rows="10" />
            </label>
        </section>

        <footer class="card-footer flex justify-end">
            <button class="btn btn-sm variant-filled"> save </button>
        </footer>
    </div>
</div>
