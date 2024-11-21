(() => {

    let pendingMessages = {}
    let messageId = 0;

    window.addEventListener("message", (ev) => {
        console.log("message", ev);
        if (ev.data !== "transfer_port") {
            console.log("wrong message", ev);
            return;
        }
        const port = ev.ports[0];
        port.onmessage = handleMessage;
        window["__parent_port__"] = port;
    }, false);

    const handleMessage = (ev) => {
        const data = ev.data;
        const messageId = data.msgId;
        const resolve = pendingMessages[messageId];
        delete pendingMessages[messageId];
        resolve(data.data);
        console.log("data", data);
    }

    const sendMessage = async (data) => {
        console.log("sendMessage", data);
        const msgId = messageId + 1;
        messageId = msgId;
        const p = new Promise((resolve) => {
            pendingMessages[msgId] = resolve;
        });
        data["msgId"] = msgId;
        const port = window["__parent_port__"]
        port?.postMessage(data);
        return p;
    }

    window["hostIPC"] = {
        sendMessage,
    }
    
})();


