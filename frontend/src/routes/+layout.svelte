<script lang="ts">
  import {
    Modal,
    getModalStore,
    type ModalComponent,
  } from "@skeletonlabs/skeleton";
  import { initializeStores } from "@skeletonlabs/skeleton";
  import Picker from "./portal/project/picker.svelte";
  import Options from "./portal/project/options.svelte";
  import "./styles.css";
  import AccountPicker from "./portal/projects/books/[pid]/txn/AccountPicker.svelte";
  import SalesItemPick from "./portal/projects/books/[pid]/sales/new/sub/SalesItemPick.svelte";
  import ClientPicker from "./portal/projects/books/[pid]/sales/new/sub/ClientPicker.svelte";
    import NewFolderPanel from "./portal/project/files/[pid]/sub/NewFolderPanel.svelte";

  initializeStores();

  const modalRegistry: Record<string, ModalComponent> = {
    project_picker: { ref: Picker },
    project_options: { ref: Options },
    new_folder_panel: { ref: NewFolderPanel },

    books_account_picker: { ref: AccountPicker },
    books_sales_item_pick: { ref: SalesItemPick },
    books_client_picker: { ref: ClientPicker },

  };
</script>

<Modal components={modalRegistry} />

<slot />
