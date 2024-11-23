<script lang="ts">
    import { preventDefault } from 'svelte/legacy';

    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { getContext } from "svelte";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";

    const rootApi = getContext("__api__") as RootAPI;

    let name = $state("");
    let bio = $state("");
    let utype = $state("user");
    let email = $state("");
    let phone = $state("");
    let password = $state("");

    const addUser = async () => {

        console.log(name, bio, utype, email, phone, password);

        const resp = await rootApi.addSelfUser({
            name,
            bio,
            utype,
            email,
            phone,
            password,
        });
        if (resp.status !== 200) {
            return;
        }

        goto(`/z/pages/portal/self/users`);
    };
</script>

<div class="p-2">
    <div class="card p-4">
        <header class="flex justify-center">
            <div class="flex flex-col gap-2">
                <h3 class="h3">Create New User</h3>
            </div>
        </header>

        <section class="">
            <form
                class="flex flex-col gap-4"
                onsubmit={preventDefault(addUser)}
            >
                <label class="label"
                    ><span>Name</span>
                    <input
                        class="input p-1"
                        type="text"
                        required
                        bind:value={name}
                    /></label
                >

                <label class="label">
                    <span>Email</span>
                    <input 
                    class="input p-1" 
                    type="email" 
                    bind:value={email} 
                    required                                        
                    
                    />
                </label>

                <label class="label">
                    <span>Phone</span>
                    <input 
                    class="input p-1" 
                    type="text" 
                    bind:value={phone} 
                    required
                    
                    />
                </label>

                <label class="label">
                    <span>Password</span>
                    <input
                        class="input p-1"
                        type="password"
                        bind:value={password}
                    />
                </label>

                <label class="label">
                    <span>User Type</span>
                    <select class="select" bind:value={utype}>
                        <option value="real">User</option>
                        <option value="bot">Bot</option>
                        <option value="super">Super User</option>
                    </select>
                </label>

                <label class="label">
                    <span>Bio</span>
                    <textarea class="textarea p-1" bind:value={bio} rows={4} />
                </label>

                <div class="flex justify-end pt-8 gap-2">
                    <button type="submit" class="btn btn-sm variant-filled">
        
                        Create
                    </button>
        
                </div>

            </form>
            


        </section>



    </div>
</div>
