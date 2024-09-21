
<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";

    let isInstanceRunning = false;
    let workingDir = "/home/xyz/Desktop/mnop";
    
</script>

<div class="flex items-center justify-center min-h-screen flex-col gap-4 pt-20">
    {#if isInstanceRunning}
        <h2 class="h2">Turnix is running.</h2>
        <span class="text-indigo-500 chip">{workingDir} </span>
        <p class="p text-xl">click explore to start using it. Or click on options.</p>

        <div class="flex gap-2">
            <button class="btn variant-soft-primary">Explore</button>   
            <button class="btn variant-soft-secondary">Options</button>
        </div>        

    {:else}
        <h2 class="h2">Turnix is not running</h2>
        <p class="p text-xl">
            Click start to run a instance in current directory.
        </p>
        <p class="p text-xl">Or click on options</p>

        <div class="flex gap-2">
            <button class="btn variant-filled">Start</button>
            <button class="btn variant-soft-secondary">Options</button>
        </div>        


    {/if}
</div>
