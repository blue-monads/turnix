<script>
    import { login } from "$lib/api/auth";

    let email = $state("dev@example.com");
    let password = $state("dev123");

    const loginHandle = async () => {
        const resp = await login({
            email, 
            password
        }) 
        if (resp.status !== 200) {
            return
        }

        localStorage.setItem("access_token", resp.data["access_token"] )

        window.location.pathname = "/z/pages/portal"
    }

</script>

<div class="flex justify-center items-center">
    <div class="card bg-initial overflow-hidden hover:variant-ringed">
        <header></header>

        <div class="p-4 space-y-4">
            <h3 class="h3" data-toc-ignore="">Login</h3>
            <article class="flex flex-col">
                <label class="label">
                    <span>Username:</span>
                    <input
                        class="input p-2"
                        type="text"
                        bind:value={email}
                        placeholder="Email"
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
            </article>
        </div>
        <hr class="opacity-50" />
        <footer class="p-4 flex justify-start items-center space-x-4">
            <button onclick={loginHandle} type="button" class="btn variant-filled-primary">
                <span>Submit</span>
            </button>
        </footer>
    </div>
</div>
