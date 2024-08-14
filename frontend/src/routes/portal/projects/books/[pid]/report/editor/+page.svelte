<script lang="ts">
    import {
        AppBar,
        getModalStore,
        Tab,
        TabGroup,
    } from "@skeletonlabs/skeleton";

    import CodeMirror from "svelte-codemirror-editor";
    import { sql, SQLite } from "@codemirror/lang-sql";

    import * as sampleCode from "./sampleCode";

    let tabSet: number = 0;
    let sqlCode = sampleCode.sqlCode;
    let htmlCode = sampleCode.htmlCode;
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <ol class="breadcrumb">
            <li class="crumb">
                <a class="anchor" href="/z/pages/portal/projects/books">Books</a
                >
            </li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Reports</li>
            <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            <li>Editor</li>
        </ol>
    </svelte:fragment>
</AppBar>

<div class=" flex flex-col md:flex-row w-full h-full max-h-[90vh]">
    <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full overflow-auto">
        <TabGroup>
            <Tab bind:group={tabSet} name="tab1" value={0}>
                <span>SQL Query</span>
            </Tab>
            <Tab bind:group={tabSet} name="tab2" value={1}>
                <span>UI</span>
            </Tab>

            <svelte:fragment slot="panel">
                {#if tabSet === 0}
                    <CodeMirror
                        bind:value={sqlCode}
                        lang={sql({
                            dialect: SQLite,
                        })}
                    />
                {:else if tabSet === 1}
                    <CodeMirror bind:value={htmlCode} />
                {/if}
            </svelte:fragment>
        </TabGroup>
    </div>

    <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full">
        <h4 class="h4 p-1 uppercase">Preview</h4>

        <div class="card p-2 h-full w-full">

            <iframe title="preview" srcdoc="<h1>Hello World</h1>" 
                width="100%" height="100%">
            </iframe>


        </div>
    </div>
</div>
