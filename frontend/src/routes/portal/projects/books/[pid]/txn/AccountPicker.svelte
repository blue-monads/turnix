<script lang="ts">
  import { AutoTable } from "$lib/compo";

  import { getModalStore } from "@skeletonlabs/skeleton";

  const store = getModalStore();

  let data: object[] = $store[0].meta["data"];
</script>

<AutoTable
  action_key={"id"}
  key_names={[
    ["id", "ID"],
    ["name", "Name"],
    ["info", "Info"],
    ["acc_type", "Account type"],
  ]}
  datas={data}
  color={["acc_type"]}
  actions={[
    {
      Name: "select",

      icon: "plus",
      // @ts-ignore
      Action: async (id) => {
        $store[0].meta["onClick"](id)
        store.close()

      },
    },
  ]}
/>
