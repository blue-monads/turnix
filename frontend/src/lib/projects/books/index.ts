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

export interface Catagory {
    id: number
    name: string
    info: string
    image: string
    parent_id: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface Product {
    id: number
    name: string
    info: string
    variant_id: string
    price: number
    parent_id: number
    image: string
    stock_count: number
    epoch: number
    alt_images: string
    created_by: number
    catagory_id: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface Contact {
    id: number
    name: string
    ctype: string
    info: string
    email: string
    phone: string
    phone2: string
    phone3: string
    address: string
    address2: string
    address3: string
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface Sale {
    id: number
    title: string
    client_id: number
    client_name: string
    notes: string
    attachments: string
    total_item_price: number
    total_item_tax_amount: number
    total_item_discount_amount: number
    sub_total: number
    overall_discount_amount: number
    overall_tax_amount: number
    total: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface SaleLine {
    id: number
    info: string
    qty: number
    sale_id: number
    product_id: number
    price: number
    tax_amount: number
    discount_amount: number
    total_amount: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface SalesData {
    sale: Sale
    lines: SaleLine[]
}

export interface Tax {
    id: number
    name: string
    ttype: string
    info: string
    rate: number
    strict: boolean
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
    is_deleted: boolean
}

export interface ProductTax {
    id: number
    catagory_id: number
    product_id: number
    tax_id: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
}


export interface StockIn {
    id: number
    info: string
    amount: number
    vendor_id: number
}

export interface StockInLine {
    id: number
    info: string
    product_id: number
    qty: number
    amount: number
}


export interface EstimateData {
    estimate: Estimate
    lines: EstimateLine[]
}

export interface Estimate {
    id: number
    title: string
    client_id: number
    client_name: string
    notes: string
    attachments: string
    total_item_price: number
    total_item_tax_amount: number
    total_item_discount_amount: number
    sub_total: number
    overall_discount_amount: number
    overall_tax_amount: number
    total: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
}

export interface EstimateLine {
    id: number
    info: string
    qty: number
    sale_id: number
    product_id: number
    price: number
    tax_amount: number
    discount_amount: number
    total_amount: number
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
}


export interface ReportTemplate {
    id: number
    name: string
    input_options: string
    report_type: string
    output_options: string
    query_template: string
    filter_script: string
    viewer_editable: boolean
    template: string
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
}

export interface SavedReport {
    id: number
    name: string
    template_id: number
    result: string
    created_by: number
    updated_by: number
    created_at: string
    updated_at: string
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

    generateLiveTxn = (pid: string, opts: Record<string, any>) => {
        return this.client.post(`books/${pid}/report/live`, opts)
    }

    exportData = (pid: string) => {
        return this.client.post(`books/${pid}/txn/export`)
    }

    importData = (pid: string, data: Record<string, any>) => {
        return this.client.post(`books/${pid}/txn/import`, data)
    }


    listCatagories = (pid: string) => {
        return this.client.get<Catagory[]>(`books/${pid}/inventory/catagories`)
    }

    addCatagory = (pid: string, data: Partial<Catagory>) => {
        return this.client.post(`books/${pid}/inventory/catagories`, data)
    }

    getCatagory = (pid: string, id: string) => {
        return this.client.get<Catagory>(`books/${pid}/inventory/catagories/${id}`)
    }

    updateCatagory = (pid: string, id: string, data: Partial<Catagory>) => {
        return this.client.post(`books/${pid}/inventory/catagories/${id}`, data)
    }

    deleteCatagory = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/inventory/catagories/${id}`)
    }

    // products

    listProducts = (pid: string) => {
        return this.client.get<Product[]>(`books/${pid}/inventory/products`)
    }

    addProduct = (pid: string, data: Partial<Product>) => {
        return this.client.post(`books/${pid}/inventory/products`, data)
    }

    getProduct = (pid: string, id: string) => {
        return this.client.get<Product>(`books/${pid}/inventory/products/${id}`)
    }

    updateProduct = (pid: string, id: string, data: Partial<Product>) => {
        return this.client.post(`books/${pid}/inventory/products/${id}`, data)
    }

    deleteProduct = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/inventory/products/${id}`)
    }

    // contacts

    listContacts = (pid: string) => {
        return this.client.get<Contact[]>(`books/${pid}/contacts`)
    }

    addContact = (pid: string, data: Partial<Contact>) => {
        return this.client.post(`books/${pid}/contacts`, data)
    }

    getContact = (pid: string, id: string) => {
        return this.client.get<Contact>(`books/${pid}/contacts/${id}`)
    }

    updateContact = (pid: string, id: string, data: Partial<Contact>) => {
        return this.client.post(`books/${pid}/contacts/${id}`, data)
    }

    deleteContact = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/contacts/${id}`)
    }

    // sales

    addSale = (pid: string, data: SalesData) => {
        return this.client.post(`books/${pid}/sales`, data)
    }

    getSale = (pid: string, id: string) => {
        return this.client.get<SalesData>(`books/${pid}/sales/${id}`)
    }

    listSales = (pid: string) => {
        return this.client.get<Sale[]>(`books/${pid}/sales`)
    }

    deleteSale = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/sales/${id}`)
    }

    // estimates

    listEstimate = (pid: string) => {
		return this.client.get<Estimate[]>(`books/${pid}/estimates`)
	}

	addEstimate = (pid: string, data: EstimateData) => {
		return this.client.post(`books/${pid}/estimates`, data)
	}

	getEstimate = (pid: string, id: string) => {
		return this.client.get<EstimateData>(`books/${pid}/estimates/${id}`)
	}

	deleteEstimate = (pid: string, id: string) => {
		return this.client.delete(`books/${pid}/estimates/${id}`)
	}


    // tax 
    listTax = (pid: string) => {
        return this.client.get<Tax[]>(`books/${pid}/tax`)
    }
    addTax = (pid: string, data: Partial<Tax>) => {
        return this.client.post(`books/${pid}/tax`, data)
    }
    getTax = (pid: string, id: string) => {
        return this.client.get<Tax>(`books/${pid}/tax/${id}`)
    }
    updateTax = (pid: string, id: string, data: Partial<Tax>) => {
        return this.client.post(`books/${pid}/tax/${id}`, data)
    }
    deleteTax = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/tax/${id}`)
    }

    listTaxProduct = (pid: string, tid?: string) => {
        return this.client.get<ProductTax[]>(`books/${pid}/tax/product?tid=${tid || ""}`)
    }

    addTaxProduct = (pid: string, id: string, data: Partial<ProductTax>) => {
        return this.client.post(`books/${pid}/tax/${id}/product`, data)
    }
    deleteTaxProduct = (pid: string, id: string, tpid: string) => {
        return this.client.delete(`books/${pid}/tax/${id}/product/${tpid}`)
    }

    // stocks

    listProductStockIn = (pid: string) => {
        return this.client.get<StockIn[]>(`books/${pid}/stocks`)
    }

    addProductStockIn = (pid: string, data: Partial<StockIn>) => {
        return this.client.post(`books/${pid}/stocks`, data)
    }

    getProductStockIn = (pid: string, id: string) => {
        return this.client.get<StockIn>(`books/${pid}/stocks/${id}`)
    }

    deleteProductStockIn = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/stocks/${id}`)
    }


        
	// report template
	

    listReportTemplate = (pid: string, offset?: string) => {
        return this.client.get<ReportTemplate[]>(`books/${pid}/reportTemplate?offset=${offset || ""}`)
    }

    addReportTemplate = (pid: string, data: Partial<ReportTemplate>) => {
        return this.client.post(`books/${pid}/reportTemplate`, data)
    }

    getReportTemplate = (pid: string, id: string) => {
        return this.client.get<ReportTemplate>(`books/${pid}/reportTemplate/${id}`)
    }

    updateReportTemplate = (pid: string, id: string, data: Partial<any>) => {
        return this.client.post(`books/${pid}/reportTemplate/${id}`, data)
    }

    deleteReportTemplate = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/reportTemplate/${id}`)
    }


    // report saved


    listReportSaved = (pid: string, offset?: string) => {
        return this.client.get<SavedReport[]>(`books/${pid}/reportSaved?offset=${offset || ""}`)
    }

    addReportSaved = (pid: string, data: Partial<SavedReport>) => {
        return this.client.post(`books/${pid}/reportSaved`, data)
    }

    getReportSavedGe = (pid: string, id: string) => {
        return this.client.get<SavedReport>(`books/${pid}/reportSaved/${id}`)
    }
    
    updateReportSaved = (pid: string, id: string, data: Partial<any>) => {
        return this.client.post(`books/${pid}/reportSaved/${id}`, data)
    }

    deleteReportSaved = (pid: string, id: string) => {
        return this.client.delete(`books/${pid}/reportSaved/${id}`)
    }

}


