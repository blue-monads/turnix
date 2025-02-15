
export interface TxnLinData {
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


export interface TxnLine {
    txn: Transaction
    lines: Line[]
}
