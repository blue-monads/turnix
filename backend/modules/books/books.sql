
-- ACCOUNTING 

create table Accounts__project__(
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

create table Transactions__project__(
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

create table TransactionLines__project__(
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

create table ReportTemplates__project__(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    input_options JSON NOT NULL DEFAULT '{}',
    output_options JSON NOT NULL DEFAULT '{}',
    query_template TEXT NOT NULL DEFAULT '',
    filter_script TEXT NULL DEFAULT ''
);
create table Reports__project__(
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

-- Estimates has not direct link to Txn, it will be used as template to create actual invoice or sales

create table Estimates__project__(
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

create table EstimateLines__project__(
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


create table Invoices__project__(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',
    tax_id INTEGER NOT NULL DEFAULT 0,
    sub_total INTEGER NOT NULL DEFAULT 0,
    client_id INTEGER NOT NULL DEFAULT 0,
    total INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table InvoiceLines__project__(
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

create table Catagories__project__(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    parent_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table Products__project__(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    variant_id TEXT NOT NULL DEFAULT '',
    price INTEGER NOT NULL DEFAULT 0,
    parent_id INTEGER NOT NULL DEFAULT 0,
    images TEXT NOT NULL DEFAULT '',
    -- epoch INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table ProductStockIn__project__(
    id INTEGER PRIMARY KEY,
    info TEXT NOT NULL DEFAULT '',
    amount INTEGER NOT NULL DEFAULT 0,
    vendor_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table ProductStockInLines__project__(
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

create table Sales__project__(
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


create table SalesLines__project__(
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


create table Tax__project__(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    rate INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NULL,
    updated_by INTEGER NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);