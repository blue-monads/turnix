<script>
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    const storeData = {
        name: "Turnix Store",
        page: 0,
        listings: [
            {
                name: "Calender",
                slug: "simpleCalender",
                info: "This is a simple calender application.",
                tags: ["calender", "simple"],
                url: ""
            },
        ],
    };
</script>

<div class="flex flex-col gap-2 md:p-5">
    <div class="flex items-center max-w-md mx-auto">
        <h2 class="h2">{storeData.name}</h2>
    </div>
    <div class="flex min-h-18 p-4">
        <div class="flex items-center max-w-md mx-auto bg-white rounded-lg">
            <div class="w-full">
                <input
                    type="search"
                    class="w-full p-2 text-gray-800 rounded-full focus:outline-none"
                    placeholder="search"
                />
            </div>
            <div>
                <button
                    class="flex items-center justify-center w-12 h-12 text-white"
                >
                    <SvgIcon
                        name="magnifying-glass"
                        className="w-8 h-8 text-gray-500"
                    />
                </button>
            </div>
        </div>
    </div>

    <div class="flex p-2 justify-center gap-2">
        {#each storeData.listings as list}
            <div class="min-h-12 card flex p-4 gap-2 flex-col md:flex-row">
                <div>
                    <SvgIcon
                        className="w-16 h-16 text-gray-500"
                        name="squares-2x2"
                    />
                </div>

                <div class="flex flex-col min-w-32">
                    <h5 class="h5 text-gray-800">{list.name}</h5>
                    <p class="p">{list.info}</p>
                    <div class="inline-flex gap-1">
                        {#each list.tags as tag }
                            <span class="chip variant-ringed hover:underline">{tag}</span>
                        {/each}

                    </div>                    
                </div>

                <div class="flex md:justify-center items-center">
                    <button
                        class="btn btn-md variant-filled text-xs hover:opacity-45"
                    >
                        Install
                    </button>
                </div>
            </div>
        {/each}
    </div>
</div>
