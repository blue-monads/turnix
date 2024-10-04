<script lang="ts">
    import { getModalStore } from "@skeletonlabs/skeleton";

    const store = getModalStore();

    interface Props {
        parent: any;
        sub_total?: number;
        overall_tax_amount?: number;
        onSet?: any;
    }

    let {
        parent,
        sub_total = 0,
        overall_tax_amount = 0,
        onSet = (data: any) => {}
    }: Props = $props();
</script>

<div class="card rounded w-modal-slim {parent.backgroundColor}">TODO</div>
