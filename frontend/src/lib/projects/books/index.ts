import type { RootAPI } from "$lib/api";
import type { AxiosInstance } from "axios";


export const NewBookAPI = (api: RootAPI) => {
    return new BooksAPI(api.client)
}



export class BooksAPI {
    client: AxiosInstance

    constructor(client: AxiosInstance) {
        this.client = client
    }


    // accounts

    listAccount = (pid: string) => {
        return this.client.get(`/${pid}/account`)
    }

    addAccount = (pid: string, data: any) => {
        return this.client.post(`/${pid}/account`, data)
    }

    getAccount = (pid: string, aid: string) => {
        return this.client.get(`/${pid}/account/${aid}`)
    }

    updateAccount = (pid: string, aid: string, data: object) => {
        return this.client.post(`/${pid}/account/${aid}`, data)
    }

    deleteAccount = (pid: string, aid: string) => {
        return this.client.delete(`/${pid}/account/${aid}`)
    }


    // transactions

    listTxn = (pid: string) => {
        return this.client.get(`/${pid}/txn`)
    }
    addTxn = (pid: string, data: object) => {
        return this.client.post(`/${pid}/txn`, data )
    }
    getTxn = (pid: string, tid: string) => {
        return this.client.get(`/${pid}/txn/${tid}`)
    }
    updateTxn = (pid: string, tid: string) => {
        return this.client.post(`/${pid}/txn/${tid}`)
    }
    deleteTxn = (pid: string, tid: string) => {
        return this.client.delete(`/${pid}/txn/${tid}`)
    }

}


