<div class="flex p-2 flex-col gap-2 overflow-auto max-h-screen">
    <div class="card p-4 flex flex-col gap-2">
        <header class="card-header">
            <h4 class="h4">Setting</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <label class="label flex flex-col">
                <span>Port</span>
                <input
                    class="input p-0.5 w-64"
                    type="text"
                    value="8000"
                    placeholder="name"
                />

                <span class="italic text-xs">Zero means random</span>
            </label>

            <div class="label flex flex-col">
                <span>P2P Proxy Any PORT</span>

                <input type="checkbox" class="checkbox" />
            </div>

            <div class="label flex flex-col">
                <span>Save config on file</span>

                <input type="checkbox" class="checkbox" />
            </div>

            <div class="label flex flex-col">
                <span>Auto start on boot</span>
                <input type="checkbox" class="checkbox" />
            </div>

            <div class="label flex flex-col">
                <span>Secret Key</span>
                <input
                    class="input p-0.5 w-72"
                    type="password"
                    value="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
                    placeholder="name"
                />
            </div>

            <div class="label flex flex-col">
                <span>Peer ID</span>
                <textarea
                    class="input textarea"
                    value="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
                    disabled
                    rows="3"
                />
            </div>

            <div class="label flex flex-col">
                <span>Full Public Key</span>
                <textarea
                    class="input textarea"
                    value="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
                    disabled
                    rows="3"
                />
            </div>
        </section>
        <footer class="card-footer flex justify-end gap-2">
            <button class="btn btn-sm variant-ghost-primary"> other instances </button>            
            <button class="btn btn-sm variant-filled"> save </button>
        </footer>
    </div>
</div>
