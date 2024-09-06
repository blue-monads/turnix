
<script lang="ts">

    let message = "";

</script>

<div class="p-2">
    <div class="card p-2">
        <header class="flex flex-col items-center justify-center">

            <div class="flex flex-col gap-2 p-2">
                <div class="ml-4 w-16 h-16 rounded-full bg-gray-500"></div>
                <h3 class="h3">John Doe</h3>
                <span class="text-sm opacity-50">john@gmail.com</span>

                <span class="chip variant-filled">bot</span>

            </div>

            <p>I am a nobody lupurem urbis quidem, quid enim est, quod est, utrumque, sic facit. Quid enim est? Quid est? Quid est enim? Quid est? Quid est enim? Quid est?</p>



        </header>
        <section class="flex flex-col gap-4">
                        
        
            <label class="label"
                ><span class="font-semibold">Send message</span>
                <textarea bind:value={message} class="textarea" rows="4" placeholder="I am a nobody"
                ></textarea></label
            >

            <p class="text-xs italic">User might not be able to see this message, based on his/her notification settings</p>

            <div class="flex justify-end pt-8 gap-2">
                <button disabled={message.length === 0} type="submit" class="btn btn-sm variant-filled">
                    Send
                </button>

            </div>


        </section>
        
    </div>
</div>
