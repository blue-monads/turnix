<script lang="ts">
</script>

<div class="p-2">
    <div class="card p-2">
        <header class="flex justify-center">
            <div class="flex flex-col gap-2">
                <h2 class="h2">Profile</h2>

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
    </div>
</div>
