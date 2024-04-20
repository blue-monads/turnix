<script lang="ts">
	import { Modal,getModalStore, type ModalComponent } from "@skeletonlabs/skeleton";
	import { initializeStores } from "@skeletonlabs/skeleton";
	import Picker from "./portal/project/picker.svelte";
    import Options from "./portal/project/options.svelte";
	import "./styles.css";

	initializeStores();

	const modalRegistry: Record<string, ModalComponent> = {
		project_picker: { ref: Picker },
		project_options: {ref: Options}

	};
</script>

<Modal components={modalRegistry} />

<slot />
