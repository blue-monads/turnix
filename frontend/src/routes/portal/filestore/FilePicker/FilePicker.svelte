<script lang="ts">
    import { RadioGroup, RadioItem } from "@skeletonlabs/skeleton";
    import FileListings from "../../self/files/FileListings.svelte";

    let value = $state(0);
</script>

<div class="w-full bg-white">

    <div class="flex justify-between bg-slate-200 items-center px-2 rounded">
        <div class="flex items-center">
            <h4 class="h4 py-2">File Picker</h4>
        </div>

        <div class="flex items-center">
            <RadioGroup>
                <RadioItem bind:group={value} name="justify" value={0}
                    padding="p-0.5 px-2 text-sm"
                    >Personal
                </RadioItem>
                <RadioItem bind:group={value} name="justify" value={1}
                    padding="p-0.5 px-2 text-sm"
                >
                    Project
                </RadioItem>
            </RadioGroup>
        </div>


    </div>


    <div class="py-2 text-xs">
        <FileListings
            baseUrl=""
            path=""
            files={[]}
            loading={false}
            selected=""
        />
    </div>
</div>
