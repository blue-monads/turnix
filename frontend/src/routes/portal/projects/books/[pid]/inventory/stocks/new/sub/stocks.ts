
export interface StockInData {
    stock: NewStockIn;
    lines: NewStockInLine[];
}

export interface NewStockIn {
    info: string;
    amount: number;
    vendor_id: number;
}

export interface NewStockInLine{
    info: string;
    product_id: number;
    qty: number;
    amount: number;
}