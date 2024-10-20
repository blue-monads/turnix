import type { Line, TxnLinData, TxnLine } from "$lib/container/books/txntype"





export const formatResponse = (data: TxnLinData): { maxId: number, txns: TxnLine[] } => {

    let index: Record<number, Line[]> = {}
    let maxId = 0

    data.lines.forEach((line) => {
        let list = index[line.txn_id]
        if (!list) {
            list = []
            index[line.txn_id] = list
        }

        list.push(line)

    })


    const txns = data.transactions.map((txn) => {
        maxId = Math.max(maxId, txn.id)

        return {
            lines: index[txn.id] || [],
            txn
        }
    })

    return {
        maxId,
        txns
    }
}