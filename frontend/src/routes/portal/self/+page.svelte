<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";

</script>


<AppBar>
	<svelte:fragment slot="lead">
        <h4 class="h4">Profile</h4>
    </svelte:fragment>

	<svelte:fragment slot="trail">
        <a href="/z/pages/portal/self/users" class="btn variant-filled">Users</a>
        <a href="/z/pages/portal/self/devices" class="btn variant-filled">Devices</a>
    </svelte:fragment>
</AppBar>


<div class="p-2">    
    <div class="card p-2">
        <header class="flex justify-center">
            <div class="flex flex-col gap-2">
                <div class="w-16 h-16 rounded-full bg-gray-500"></div>
            </div>
        </header>

        <section class="flex flex-col gap-4">
            <label class="label">
                <span>Name</span>
                <input class="input p-1" type="text" placeholder="John" />
            </label>

            <label class="label">
                <span>Email</span>
                <input
                    class="input p-1"
                    type="email"
                    value="xyz@example.com"
                    disabled
                    placeholder="user@example.com"
                />
            </label>

            <label class="label">
                <span>User Type</span>
                <select class="select" disabled>
                    <option value="5">user</option>
                </select>
            </label>

            <label class="label">
                <span>Bio</span>
                <textarea
                    class="textarea"
                    rows="4"
                    placeholder="I am a nobody"
                />
            </label>
        </section>

        <footer class="flex justify-end p-2 gap-2">
            <a href="/z/pages/portal/self/reset-password" class="btn variant-ghost-secondary" >
                Change Password
            </a>
            
            <a href="/z/pages/portal/self/reset-email" class="btn variant-ghost-tertiary" >
                Reset email
            </a>

            <button class="btn variant-filled" disabled>
                Update
            </button>


        </footer>

    </div>
</div>
