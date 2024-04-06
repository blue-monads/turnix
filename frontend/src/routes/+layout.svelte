<script lang="ts">
	import {
		Modal,
		getModalStore,
		type ModalComponent,
	} from "@skeletonlabs/skeleton";
	import { initializeStores } from "@skeletonlabs/skeleton";
	import Picker from "./portal/projects/picker.svelte";

	import "./styles.css";
	initializeStores();

	const modalRegistry: Record<string, ModalComponent> = {
		picker: { ref: Picker },
	};
</script>

<Modal components={modalRegistry} />

<slot />
