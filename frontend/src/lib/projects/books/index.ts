import type { RootAPI } from "$lib/api";
import type { AxiosInstance } from "axios";


export const NewBookAPI = (api: RootAPI) => {
    return new BooksAPI(api.projectClient)
}



export class BooksAPI {
    client: AxiosInstance

    constructor(client: AxiosInstance) {
        this.client = client
    }

    // /z/api/v1/project/:pid


    // accounts

    listAccount = (pid: string) => {
        return this.client.get(`books/${pid}/account`)
    }

    addAccount = (pid: string, data: any) => {
        return this.client.post(`books/${pid}/account`, data)
    }

    getAccount = (pid: string, aid: string) => {
        return this.client.get(`books/${pid}/account/${aid}`)
    }

    updateAccount = (pid: string, aid: string, data: object) => {
        return this.client.post(`books/${pid}/account/${aid}`, data)
    }

    deleteAccount = (pid: string, aid: string) => {
        return this.client.delete(`books/${pid}/account/${aid}`)
    }


    // transactions


    listTxn = (pid: string) => {
        return this.client.get(`books/${pid}/txn`)
    }

    listTxnWithLines = (pid: string, offset?: string) => {
        return this.client.get(`books/${pid}/txn/line/list?offset=${offset || 0}`)
    }

    listAccTxnWithLines = (pid: string, aid: string, offset?: string) => {
        return this.client.get(`books/${pid}/txn/line/${aid}/list?offset=${offset || 0}`)
    }


    addTxn = (pid: string, data: object) => {
        return this.client.post(`books/${pid}/txn`, data )
    }
    getTxn = (pid: string, tid: string) => {
        return this.client.get(`books/${pid}/txn/${tid}`)
    }
    updateTxn = (pid: string, tid: string) => {
        return this.client.post(`books/${pid}/txn/${tid}`)
    }
    deleteTxn = (pid: string, tid: string) => {
        return this.client.delete(`books/${pid}/txn/${tid}`)
    }

}


