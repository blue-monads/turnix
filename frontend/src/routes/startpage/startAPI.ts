import axios from "axios"

export interface Status {
    is_running: boolean;
    port?: string;
    status: string;
}



export const buildApi = async  (base: string) => {

    console.log("buildApi", base);

    const client = axios.create({
        baseURL: base,
        timeout: 10000,
        headers: {
            "NodeCTRLKey": "fixme"
        }
    });


    const getStatus = async () => {
        return await client.get<Status>("/z/eapi/status")
    }

    const startInstance = async () => {
        return await client.post("/z/eapi/start", {})
    }


    return { getStatus, startInstance }
}

export type StartAPI = Awaited< ReturnType<typeof buildApi>>