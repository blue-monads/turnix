<script lang="ts">
    import AutoForm from "../../../../../../lib/compo/autoform/auto_form.svelte";
    import { params } from "$lib/params";
    import type { API } from "$lib/api";
    import { getContext } from "svelte";

    let pid = $params["pid"];
    let mid = $params["mid"];
    let message = "";


    const api = getContext("__api__") as API;

    const load = async () => {
        

    }



</script>


<AutoForm
    data={{}}
    {message}
    schema={{
        fields: [
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "contents",
                name: "Contents",
            },
            {
                ftype: "SELECT",
                key_name: "status",
                name: "Status",
                options: ["submited", "processed", "denied"],
            },
            {
                ftype: "KEY_VALUE_TEXT",
                key_name: "extrameta",
                name: "Extra Meta",
            },
        ],
        name: "Add Queue message",
        required_fields: ["name"],
    }}
    onSave={async (newdata) => {


    }}
/>
