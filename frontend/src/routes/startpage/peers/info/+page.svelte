<script lang="ts">
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";

    const peerInfo = $state({
        id: "xx-yy-zz",
        added: "2022-12-31",
        lastSeen: "2022-12-31",
        status: "online",
        fullPublicKey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        peerType: "guest", // admin
        preConnect: false,
        adminKey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        multiAddrs: [
            "/ip6/2604:a880:1:20::204:4001/udp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
        ],

        staticAddrs: [
            {
                ip: "44.55.66.77",
                port: 1234,
                protocol: "tcp",
            },
        ],

        staticRelays: [
            {
                ip: "44.55.66.77",
                port: 1234,
                protocol: "tcp",
                relayToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            },
        ],
    });
</script>

<div class="flex p-2 flex-col gap-2 overflow-auto max-h-screen">
    <div class="card p-4 flex flex-col gap-2">
        <header class="card-header">
            <h4 class="h4">Peer Info</h4>
        </header>

        <section class="p-4 flex flex-col gap-4">
            <label class="label flex flex-col">
                <span>Id</span>
                <input
                    bind:value={peerInfo.id}
                    class="input p-0.5 w-64"
                    disabled
                    type="text"
                    placeholder="name"
                />
            </label>

            <div class="label flex flex-col">
                <span>Status</span>
                <div>
                    <span class="chip variant-filled">{peerInfo.status}</span>
                </div>
            </div>

            <!-- last seen -->

            <div class="label flex flex-col">
                <span>Last seen</span>
                <div>
                    <span class="chip variant-filled">{peerInfo.lastSeen}</span>
                </div>
            </div>

            <!-- peer type -->

            <div class="label flex flex-col">
                <span>Peer type</span>
                <select bind:value={peerInfo.peerType} class="select">
                    <option value="guest">Guest</option>
                    <option value="admin">Admin</option>
                </select>
            </div>

            <!-- fullpublickey -->

            <div class="label flex flex-col">
                <span>Full public key</span>
                <textarea class="input textarea" disabled
                    >{peerInfo.fullPublicKey}</textarea
                >
            </div>

            <div class="label flex flex-col">
                <span>Admin key</span>
                <textarea class="input textarea" disabled
                    >{peerInfo.adminKey}</textarea
                >
            </div>

            <!-- preconnect -->
            <div class="label flex flex-col">
                <span>Pre connect</span>

                <input
                    value={peerInfo.preConnect}
                    type="checkbox"
                    class="checkbox"
                />
            </div>

            <!-- multiaddrs table -->

            <div class="label flex flex-col">
                <span>Multiaddrs</span>
                <table class="table table-hover overflow-auto input">
                    <thead>
                        <tr>
                            <th>Addr</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each peerInfo.multiAddrs as addr}
                            <tr>
                                <td>
                                    <span class="chip variant-ghost-secondary text-lg">
                                        {addr}

                                    </span>

                                
                                </td>

                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>

            <!-- staticaddrs table -->

            <div class="label flex flex-col">
                <span>Static addrs</span>

                <table class="table table-hover overflow-auto input">
                    <thead>
                        <tr>
                            <th>Protocol</th>
                            <th>IP</th>
                            <th>Port</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each peerInfo.staticAddrs as addr}
                            <tr>
                                <td>{addr.protocol}</td>
                                <td>{addr.ip}</td>
                                <td>{addr.port}</td>
                            </tr>
                        {/each}
                    </tbody>
                </table>

                <div class="label flex justify-start p-1">
                    <button class="btn btn-sm variant-filled">
                        <SvgIcon name="plus" className="w-4 h-4" />
                        add
                    </button>
                </div>

            </div>

            <!-- static relays table -->

            <div class="label flex flex-col">
                <span>Static relays</span>
                <table class="table table-hover overflow-auto input">
                    <thead>
                        <tr>
                            <th>Protocol</th>
                            <th>IP</th>
                            <th>Port</th>
                            <th>Relay token</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each peerInfo.staticRelays as relay}
                            <tr>
                                <td>{relay.protocol}</td>
                                <td>{relay.ip}</td>
                                <td>{relay.port}</td>
                                <td>{relay.relayToken}</td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
                <div class="label flex justify-start p-1">
                    <button class="btn btn-sm variant-filled">
                        <SvgIcon name="plus" className="w-4 h-4" />
                        add
                    </button>
                </div>
            </div>

            <footer class="card-footer flex justify-end gap-2">
                <button class="btn btn-sm variant-filled-primary">
                    shell
                </button>
                <button class="btn btn-sm variant-filled-secondary">
                    share
                </button>
                <button class="btn btn-sm variant-filled-tertiary mr-8">
                    chat
                </button>

                <button class="btn btn-sm variant-filled"> save </button>

                <button class="btn btn-sm variant-filled-error">
                    delete
                </button>
            </footer>
        </section>
    </div>
</div>
