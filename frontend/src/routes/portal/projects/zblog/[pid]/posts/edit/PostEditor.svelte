<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";
    import CodeMirror from "svelte-codemirror-editor";
    import {markdown} from "@codemirror/lang-markdown"

    interface Props {
        content: string;
    }

    let {
        content = $bindable(""),
    } = $props() as Props;
</script>


<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto card p-4">
    <CodeMirror
    bind:value={content}
    lang={markdown()}
    extensions={[]}


/>

</div>

