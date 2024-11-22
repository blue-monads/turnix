<script lang="ts">
    import { Tab, TabGroup } from "@skeletonlabs/skeleton";
    import CodeMirror from "svelte-codemirror-editor";
    import { sql, SQLite } from "@codemirror/lang-sql";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { page } from "$app/stores";
    import { javascript } from "@codemirror/lang-javascript";
    import Loader from "$lib/compo/loader/loader.svelte";

    interface Props {
        isReadOnly?: boolean;
    }

    let { isReadOnly = false }: Props = $props();

    let sqlCode = $state("select * from __project__Accounts;");
    let jsCode = $state("const process = (ctx) => {\n}");

    const pid = $page.params["pid"];
    const rootApi = getContext("__api__") as RootAPI;

    let lastData: Record<string, any>[] = $state([]);
    let loading = $state(false);
    let tabMode: "sql" | "formatter" = $state("sql");
    let runError = $state("");
    let inferedColumns: string[] = $state([]);

    const runLoad = async () => {
        loading = true;
        runError = "";
        const resp = await rootApi.runProjectSQL2(pid, {
            qstr: sqlCode,
            data: [],
        });
        if (resp.status !== 200) {
            runError = `${resp.statusText} : ${resp.data}`;
            return;
        }

        lastData = resp.data;
        loading = false;
    };

    $effect(() => {
        if (!lastData.length) {
            inferedColumns = [];
            return;
        }

        const nextColumn = Object.keys(lastData[0]);
        if (nextColumn.includes("id")) {
            inferedColumns = ["id", ...nextColumn.filter((x) => x !== "id")];
        } else {
            inferedColumns = nextColumn;
        }
    });
</script>

<div class="flex flex-col w-full h-[94vh] p-2">
    <div class="w-full p-2 rounded">
        <TabGroup>
            <Tab bind:group={tabMode} name="SQL Query" value={"sql"}>
                <span>SQL</span>
            </Tab>

            <Tab bind:group={tabMode} name="Formatter" value={"formatter"}>
                <span>Formatter</span>
            </Tab>

            <svelte:fragment slot="panel">
                <div class="min-h-10 card overflow-auto resize-y p-2">
                    {#if tabMode === "sql"}
                        <CodeMirror
                            bind:value={sqlCode}
                            readonly={isReadOnly}
                            extensions={[]}
                            lang={sql({
                                dialect: SQLite,
                            })}
                        />
                    {:else}
                        <CodeMirror
                            bind:value={jsCode}
                            readonly={isReadOnly}
                            extensions={[]}
                            lang={javascript()}
                        />
                    {/if}
                </div>
            </svelte:fragment>
        </TabGroup>
    </div>

    <div>
        <p class="text-red-600">{runError}</p>
    </div>

    <div class="flex justify-end p-2">
        <button
            class="inline-flex rounded p.05 variant-filled self-center"
            onclick={runLoad}
        >
            <SvgIcon className="w-4 h-4 mt-1" name="play" />
            Run
        </button>
    </div>
    <div class="w-full h-full p-2 overflow-auto">
        {#if loading}
            <Loader />
        {:else}
            <table class="table table-hover overflow-auto input">
                <thead>
                    <tr>
                        {#each inferedColumns as col}
                            <td class="p-1">
                                {col}
                            </td>
                        {/each}
                    </tr>
                </thead>

                <tbody>
                    {#each lastData as row}
                        <tr>
                            {#each inferedColumns as col}
                                {@const cellval = row[col]}
                                <td class="p-1">
                                    {#if cellval !== undefined}
                                        {row[col]}
                                    {/if}
                                </td>
                            {/each}
                        </tr>
                    {/each}
                </tbody>
            </table>
        {/if}
    </div>
</div>
