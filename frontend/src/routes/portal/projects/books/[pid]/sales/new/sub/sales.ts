import { writable } from "svelte/store";


export interface SaleLine {
    info: string;
    product_id: number;
    qty: number;
    amount: number;
}


