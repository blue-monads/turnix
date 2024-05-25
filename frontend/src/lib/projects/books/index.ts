import type { RootAPI } from "$lib/api";
import type { AxiosInstance } from "axios";

export interface Account {
    id: number
    name: string
    info: string
    acc_type: string
    total_debit: number
    total_credit: number
    contact_id: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}


export interface TxnLines {
    transactions: Transaction[]
    lines: Line[]
}

export interface Transaction {
    id: number
    title: string
    notes: string
    linked_sales_id: number
    linked_invoice_id: number
    reference_id: string
    attachments: string
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface Line {
    id: number
    account_id: number
    txn_id: number
    debit_amount: number
    credit_amount: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
}


export interface TxnData {
    title: string
    notes: string
    debit_account_id: number
    credit_account_id: number
    debit_amount: number
    credit_amount: number
    reference_id?: string
    attachments?: any
}


export interface TxnUpdateWithLineOptions {
    txn_data: Record<string, any>;
    first_line_id: number;
    first_line_data: Record<string, any>;
    second_line_id: number;
    second_line_data: Record<string, any>;
}

export interface TransactionWithLine {
    txn: Transaction;
    first_line: Line | null;
    second_line: Line | null;
}

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
        return this.client.get<Account[]>(`books/${pid}/account`)
    }

    addAccount = (pid: string, data: Partial<Account>) => {
        return this.client.post(`books/${pid}/account`, data)
    }

    getAccount = (pid: string, aid: string) => {
        return this.client.get<Account>(`books/${pid}/account/${aid}`)
    }

    updateAccount = (pid: string, aid: string, data: Partial<Account>) => {
        return this.client.post(`books/${pid}/account/${aid}`, data)
    }

    deleteAccount = (pid: string, aid: string) => {
        return this.client.delete(`books/${pid}/account/${aid}`)
    }


    // transactions


    listTxn = (pid: string) => {
        return this.client.get<Transaction[]>(`books/${pid}/txn`)
    }

    listTxnWithLines = (pid: string, offset?: string) => {
        return this.client.get<TxnLines>(`books/${pid}/txn/line/list?offset=${offset || 0}`)
    }

    listAccTxnWithLines = (pid: string, aid: string, offset?: string) => {
        return this.client.get<TxnLines>(`books/${pid}/txn/line/${aid}/list?offset=${offset || 0}`)
    }


    addTxn = (pid: string, data: TxnData) => {
        return this.client.post(`books/${pid}/txn`, data)
    }
    getTxn = (pid: string, tid: string) => {
        return this.client.get<Transaction>(`books/${pid}/txn/${tid}`)
    }

    getTxnWithLines = (pid: string, tid: string) => {
        return this.client.get<TransactionWithLine>(`books/${pid}/txn/${tid}/line`)
    }

    updateTxn = (pid: string, tid: string, data: Partial<Transaction>) => {
        return this.client.post(`books/${pid}/txn/${tid}`, data)
    }
    updateTxnWithLine = (pid: string, tid: string, data: TxnUpdateWithLineOptions) => {
        return this.client.post(`books/${pid}/txn/${tid}/line`, data)
    }

    deleteTxn = (pid: string, tid: string) => {
        return this.client.delete(`books/${pid}/txn/${tid}`)
    }

}


