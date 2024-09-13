
export interface SaleData {
    sale: NewSale;
    lines: NewSaleLine[];
}

export interface NewSale {
    title: string;
    notes: string;
    client_id: number;
    client_name: string;
    total_item_price: number;
    total_item_tax_amount: number;
    total_item_discount_amount: number;
    sub_total: number;
    overall_discount_amount: number;
    overall_tax_amount: number;
    total: number;
    sales_date: string;
}

export interface NewSaleLine {
    info: string;
    product_id: number;
    qty: number;
    amount: number;
    price: number;
    tax_amount: number;
    discount_amount: number;
    total_amount: number;
}