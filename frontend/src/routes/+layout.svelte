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
  import NewPluginPicker from "./portal/project/plugins/NewPluginPicker.svelte";
  import FilePicker from "./portal/filestore/FilePicker.svelte";
  import {  Drawer } from '@skeletonlabs/skeleton';
    import Notification from "./portal/self/inbox/notification.svelte";
    import OverAllTaxPicker from "./portal/projects/books/[pid]/sales/new/sub/OverAllTaxPicker.svelte";
    import OverAllDiscountPicker from "./portal/projects/books/[pid]/sales/new/sub/OverAllDiscountPicker.svelte";

  initializeStores();

  const modalRegistry: Record<string, ModalComponent> = {
    project_picker: { ref: Picker },
    project_options: { ref: Options },
    project_new_plugin_picker: { ref: NewPluginPicker },
    new_folder_panel: { ref: NewFolderPanel },

    books_account_picker: { ref: AccountPicker },
    books_sales_item_pick: { ref: SalesItemPick },
    books_client_picker: { ref: ClientPicker },
    books_sales_overall_tax_picker: { ref: OverAllTaxPicker },
    books_sales_overall_discount_picker: { ref: OverAllDiscountPicker },


    file_picker: { ref: FilePicker },

  };
</script>

<Drawer> 
  <Notification />
</Drawer>

<Modal components={modalRegistry} />

<slot />
