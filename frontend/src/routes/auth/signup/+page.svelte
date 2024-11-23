<script lang="ts">
    let name = $state("");
    let email = $state("");
    let password = $state("");
    let repassword = $state("");
</script>

<div class="flex justify-center items-center">
    <div class="card bg-initial card-hover overflow-hidden variant-ringed">
        <header></header>

        <div class="p-4 space-y-4">
            <h3 class="h3" data-toc-ignore="">Login</h3>
            <article class="flex flex-col">
                <label class="label">
                    <span>Name:</span>
                    <input
                        bind:value={name}
                        class="input p-2"
                        type="text"
                        placeholder="name"
                    />
                </label>

                <label class="label">
                    <span>Email:</span>
                    <input
                        bind:value={email}
                        class="input p-2"
                        type="email"
                        placeholder="email"
                    />
                </label>

                <label class="label">
                    <span>Password:</span>
                    <input
                        bind:value={password}
                        class="input p-2"
                        type="password"
                        placeholder="password"
                    />
                </label>

                <label class="label">
                    <span>Re Password:</span>
                    <input
                        bind:value={repassword}
                        class="input p-2"
                        type="password"
                        placeholder="password"
                    />
                </label>
            </article>
        </div>
        <hr class="opacity-50" />
        <footer class="p-4 flex justify-start items-center space-x-4">
            <button
                type="button"
                disabled={!name ||
                    !email ||
                    !password ||
                    password === repassword}
                class="btn variant-filled-primary"
            >
                <span>Submit</span>
            </button>
        </footer>
    </div>
</div>
