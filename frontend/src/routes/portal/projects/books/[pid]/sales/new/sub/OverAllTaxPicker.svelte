<script lang="ts">
    import { getModalStore } from "@skeletonlabs/skeleton";
    export let parent: any;

    const store = getModalStore();

    export let sub_total = 0;
    export let overall_tax_amount = 0;
    export let onSet = (data: any) => {};
</script>
