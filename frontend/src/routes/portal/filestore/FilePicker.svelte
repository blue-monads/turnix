<script lang="ts">
    import { RadioGroup, RadioItem } from "@skeletonlabs/skeleton";

    let value = 0;
</script>

<div class="w-modal-wide bg-white p-4">
    <div class="flex justify-between">
        <h4 class="h4">File Picker</h4>
        <RadioGroup>
            <RadioItem bind:group={value} name="justify" value={0}
                >Personal</RadioItem
            >
            <RadioItem bind:group={value} name="justify" value={1}
                >Project</RadioItem
            >
        </RadioGroup>
    </div>
</div>
