<script lang="ts">
    import type { RootAPI } from "$lib/api";
    import { getContext } from "svelte";

    const api = getContext("__api__") as RootAPI;

    let messages = [];

    const load = () => {
        

    }

</script>

<div>Messages</div>
