<script>
    import { FloatyButton } from "$lib/compo";
    import { AppBar } from "@skeletonlabs/skeleton";

    const users = [
        {
            name: "John Doe",
            bio: "i am john",
        },

        {
            name: "Jane Smith",
            bio: "i am jane",
        },
    ];
</script>

<AppBar>
    <svelte:fragment slot="lead">
        <h4 class="h4">Users</h4>
    </svelte:fragment>
</AppBar>

<div class="p-2">
    <div class="card p-2">
        <dl class="list-dl">
            {#each users as user}
                <div>
                    <span class="badge-icon p-4 variant-soft-secondary">
                        <i class="fa-solid fa-book"></i>
                    </span>

                    <span class="flex-auto"
                        ><dt class="font-bold">{user.name}</dt>
                        <dd class="text-sm opacity-50">
                            {user.bio}
                        </dd></span
                    > <span>⋮</span>
                </div>
            {/each}
        </dl>
    </div>
</div>

<FloatyButton handler={() => {}} />
