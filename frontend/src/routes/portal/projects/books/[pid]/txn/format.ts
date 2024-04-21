import type { Line, TxnLinData, TxnLine } from "$lib/container/books/txntype"





export const formatResponse = (data: TxnLinData): TxnLine[] => {

    let index: Record<number, Line[]> = {}

    data.lines.forEach((line) => {
        let list = index[line.txn_id]
        if (!list) {
            list = []
            index[line.txn_id] = list
        }

        list.push(line)

    })


    return data.transactions.map((txn) => {
        return {
            lines: index[txn.id] || [],
            txn
        }
    })
}