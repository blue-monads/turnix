<script lang="ts">
    import { AppBar, RadioGroup, RadioItem } from "@skeletonlabs/skeleton";
    import FileListings from "../self/files/FileListings.svelte";

    let value = 0;
</script>

<div class="w-modal-wide bg-white p-4">
    <AppBar>
        <svelte:fragment slot="lead">
            <h4 class="h4">File Picker</h4>
        </svelte:fragment>

        <svelte:fragment slot="trail">
            <RadioGroup>
                <RadioItem bind:group={value} name="justify" value={0}
                    >Personal
                </RadioItem>
                <RadioItem bind:group={value} name="justify" value={1}>
                    Project
                </RadioItem>
            </RadioGroup>
        </svelte:fragment>
    </AppBar>

    <div class="p-2">
        <FileListings
            baseUrl=""
            path=""
            files={[]}
            loading={false}
            selected=""
        />
    </div>
</div>
