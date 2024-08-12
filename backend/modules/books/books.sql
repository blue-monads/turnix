
-- ACCOUNTING 

create table __project__Accounts(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    acc_type TEXT NOT NULL DEFAULT '',
    parent_id INTEGER NOT NULL DEFAULT 0,
    total_debit INTEGER NOT NULL DEFAULT 0,
    total_credit INTEGER NOT NULL DEFAULT 0,
    contact_id INTEGER NOT NULL DEFAULT 0,
    calculated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__Transactions(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    txn_type TEXT NOT NULL DEFAULT 'normal',
    linked_sales_id INTEGER NOT NULL DEFAULT 0,
    linked_invoice_id INTEGER NOT NULL DEFAULT 0,
    linked_stockin_id INTEGER NOT NULL DEFAULT 0,
    reference_id TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    txn_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__TransactionLines(
    id INTEGER PRIMARY KEY,
    account_id INTEGER NOT NULL,
    txn_id INTEGER NULL,
    debit_amount INTEGER NOT NULL DEFAULT 0,
    credit_amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table __project__ReportTemplates(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    input_options JSON NOT NULL DEFAULT '{}',
    output_options JSON NOT NULL DEFAULT '{}',
    query_template TEXT NOT NULL DEFAULT '',
    filter_script TEXT NULL DEFAULT ''
);
create table __project__Reports(
    id INTEGER PRIMARY KEY,
    template_id INTEGER NOT NULL DEFAULT 0,
    report_type TEXT NOT NULL DEFAULT 'custom',
    result JSON NOT NULL DEFAULT '{}',
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- INVOICING

-- Estimates has not direct link to Txn, it will be used as template to create actual invoice or directsales

-- z_1_Estimates
create table __project__Estimates(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',
    tax_id INTEGER NOT NULL DEFAULT 0,
    sub_total INTEGER NOT NULL DEFAULT 0,
    total INTEGER NOT NULL DEFAULT 0,
    txn_link_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__EstimateLines(
    id INTEGER PRIMARY KEY,
    info TEXT NOT NULL DEFAULT '',
    qty INTEGER NOT NULL DEFAULT 0,    
    estimate_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL DEFAULT 0,
    amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table __project__Invoices(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    client_id INTEGER NOT NULL DEFAULT 0,
    client_name TEXT NOT NULL DEFAULT '',

    notes TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',

    tax_id INTEGER NOT NULL DEFAULT 0,
    tax_amount INTEGER NOT NULL DEFAULT 0,

    discount_id INTEGER NOT NULL DEFAULT 0,
    discount_amount INTEGER NOT NULL DEFAULT 0,

    sub_total INTEGER NOT NULL DEFAULT 0,
    total INTEGER NOT NULL DEFAULT 0,
    txn_link_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    invalidated_reason TEXT NOT NULL DEFAULT '',

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__InvoiceLines(
    id INTEGER PRIMARY KEY,
    info TEXT NOT NULL DEFAULT '',
    product_id INTEGER NOT NULL,
    qty INTEGER NOT NULL DEFAULT 0,    
    invoice_id INTEGER NOT NULL,
    amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- SALES could be made two way 
-- 1. Invoice[credit] <-[paid]-> Txn
-- 2. Sales[debit] <-[paid]-> Txn

-- INVENTORY

create table __project__Catagories(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    parent_id INTEGER NOT NULL DEFAULT 0,
    image TEXT NOT NULL DEFAULT '',
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__Products(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    variant_id TEXT NOT NULL DEFAULT '',
    catagory_id INTEGER NOT NULL DEFAULT 0,
    price INTEGER NOT NULL DEFAULT 0,
    parent_id INTEGER NOT NULL DEFAULT 0,
    image TEXT NOT NULL DEFAULT '',
    alt_images TEXT NOT NULL DEFAULT '',
    epoch INTEGER NOT NULL DEFAULT 0, -- optimistic counter for stock count
    stock_count INTEGER NOT NULL DEFAULT 0, 
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__ProductStockIn(
    id INTEGER PRIMARY KEY,
    info TEXT NOT NULL DEFAULT '',
    amount INTEGER NOT NULL DEFAULT 0,
    vendor_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table __project__ProductStockInLines(
    id INTEGER PRIMARY KEY,
    info TEXT NOT NULL DEFAULT '',
    product_stockin_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    qty INTEGER NOT NULL DEFAULT 0,
    amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- SALES

create table __project__Sales(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    client_id INTEGER NOT NULL DEFAULT 0,
    client_name TEXT NOT NULL DEFAULT '',

    notes TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',

    tax_id INTEGER NOT NULL DEFAULT 0,
    tax_amount INTEGER NOT NULL DEFAULT 0,

    discount_id INTEGER NOT NULL DEFAULT 0,
    discount_amount INTEGER NOT NULL DEFAULT 0,

    sub_total INTEGER NOT NULL DEFAULT 0,
    total INTEGER NOT NULL DEFAULT 0,
    txn_link_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    invalidated_reason TEXT NOT NULL DEFAULT '',

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__SalesLines(
    id INTEGER PRIMARY KEY,
    info TEXT NOT NULL DEFAULT '',
    qty INTEGER NOT NULL DEFAULT 0,    
    sale_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL DEFAULT 0,
    amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table __project__Tax(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    ttype TEXT NOT NULL DEFAULT 'item_percent', -- item_percent, overall_percent, item_custom, overall_custom
    info TEXT NOT NULL DEFAULT '',
    rate INTEGER NOT NULL DEFAULT 0,
    strict BOOLEAN NOT NULL DEFAULT FALSE,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__Discounts(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    ttype TEXT NOT NULL DEFAULT 'item_percent', -- item_percent, item_fixed, item_custom,  overall_percent, overall_fixed, overall_custom
    info TEXT NOT NULL DEFAULT '',
    rate INTEGER NOT NULL DEFAULT 0,
    strict BOOLEAN NOT NULL DEFAULT FALSE,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__Contacts(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    ctype TEXT NOT NULL DEFAULT 'vendor', -- vendor, client
    info TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    phone TEXT NOT NULL DEFAULT '',
    phone2 TEXT NOT NULL DEFAULT '',
    phone3 TEXT NOT NULL DEFAULT '',
    address TEXT NOT NULL DEFAULT '',    
    address2 TEXT NOT NULL DEFAULT '',
    address3 TEXT NOT NULL DEFAULT '',
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);