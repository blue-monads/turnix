import type { SimpleRATApi } from "../lib/SimpleRATApi";
 
export const webspcketTest = async (api: SimpleRATApi, pid: string, did: number) => {

    let url = api.getNewRoomUrl(pid, did)

    if (url.startsWith("https")) {
        url = url.replace("https", "wss")
    } 

    if (url.startsWith("http")) {
        url = url.replace("http", "ws")
    }

    const ws = new WebSocket(url)

    ws.onopen = () => {
        console.log("Connected")
    }

    ws.onmessage = (e) => {
        console.log("Message", e.data)
    }

    ws.onclose = () => {
        console.log("Closed")
    }

    ws.onerror = (e) => {
        console.log("Error", e)
    }




}