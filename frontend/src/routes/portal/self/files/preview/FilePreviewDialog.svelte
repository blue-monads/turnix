<script lang="ts">
    import { getModalStore } from "@skeletonlabs/skeleton";
    import FilePreview from "./FilePreview.svelte";

    const store = getModalStore();

    const data = $store[0].meta;
</script>

<div class="w-modal-wide bg-white p-4 max-h-[90vh] overflow-auto">
    <FilePreview api={data.api} fileId={data.fileId} filename={data.filename} />
</div>
