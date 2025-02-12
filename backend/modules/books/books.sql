
-- ACCOUNTING 

create table __project__Accounts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
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
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    txn_type TEXT NOT NULL DEFAULT 'normal', -- manual, sales, stockin
    reference_id TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    txn_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_editable BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__TransactionLines(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    txn_id INTEGER NULL,
    debit_amount INTEGER NOT NULL DEFAULT 0,
    credit_amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    linked_sales_id INTEGER NOT NULL DEFAULT 0,
    linked_stockin_id INTEGER NOT NULL DEFAULT 0
);

create table __project__ReportTemplates(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL DEFAULT '',
    report_type TEXT NOT NULL DEFAULT 'html_report', -- html_report, sql_report
    input_options JSON NOT NULL DEFAULT '{}',
    output_options JSON NOT NULL DEFAULT '{}',
    viewer_editable BOOLEAN NOT NULL DEFAULT FALSE,
    template TEXT NOT NULL DEFAULT '',
    filter_script TEXT NULL DEFAULT '',
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table __project__SavedReports(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL DEFAULT '',
    template_id INTEGER NOT NULL DEFAULT 0,
    result TEXT NOT NULL DEFAULT '',
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- INVOICING

-- Estimates has not direct link to Txn, it will be used as template to create actual sales

-- z_1_Estimates
create table __project__Estimates(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL DEFAULT '',
    client_id INTEGER NOT NULL DEFAULT 0,
    client_name TEXT NOT NULL DEFAULT '',

    notes TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',

    total_item_price INTEGER NOT NULL DEFAULT 0,
    total_item_tax_amount INTEGER NOT NULL DEFAULT 0,
    total_item_discount_amount INTEGER NOT NULL DEFAULT 0,

    sub_total INTEGER NOT NULL DEFAULT 0, -- total_item_price + total_item_tax_amount - total_item_discount_amount

    overall_discount_amount INTEGER NOT NULL DEFAULT 0,
    overall_tax_amount INTEGER NOT NULL DEFAULT 0,
    
    total INTEGER NOT NULL DEFAULT 0, -- sub_total +  overall_tax_amount - overall_discount_amount
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,    
    sales_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table __project__EstimateLines(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    info TEXT NOT NULL DEFAULT '',
    qty INTEGER NOT NULL DEFAULT 0,    
    estimate_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL DEFAULT 0,

    price INTEGER NOT NULL DEFAULT 0, -- original item price
    tax_amount INTEGER NOT NULL DEFAULT 0,
    discount_amount INTEGER NOT NULL DEFAULT 0,

    total_amount INTEGER NOT NULL DEFAULT 0, -- total_amount = item_price + item_tax_amount - discount_amount
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);




-- INVENTORY

create table __project__Catagories(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    product_class TEXT NOT NULL DEFAULT 'physical_item', -- physical_item, service, digital_item        
    parent_id INTEGER NOT NULL DEFAULT 0,
    image TEXT NOT NULL DEFAULT '',
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__Products(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
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
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__ProductStockIn(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    info TEXT NOT NULL DEFAULT '',
    amount INTEGER NOT NULL DEFAULT 0,
    vendor_id INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table __project__ProductStockInLines(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    info TEXT NOT NULL DEFAULT '',
    product_stockin_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    qty INTEGER NOT NULL DEFAULT 0,
    amount INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- SALES

create table __project__Sales(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL DEFAULT '',
    client_id INTEGER NOT NULL DEFAULT 0,
    client_name TEXT NOT NULL DEFAULT '',

    notes TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',

    total_item_price INTEGER NOT NULL DEFAULT 0,
    total_item_tax_amount INTEGER NOT NULL DEFAULT 0,
    total_item_discount_amount INTEGER NOT NULL DEFAULT 0,

    sub_total INTEGER NOT NULL DEFAULT 0, -- total_item_price + total_item_tax_amount - total_item_discount_amount

    overall_discount_amount INTEGER NOT NULL DEFAULT 0,
    overall_tax_amount INTEGER NOT NULL DEFAULT 0,
    
    total INTEGER NOT NULL DEFAULT 0, -- sub_total +  overall_tax_amount - overall_discount_amount
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,    
    sales_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    invalidated_reason TEXT NOT NULL DEFAULT '',
    payment_status TEXT NOT NULL DEFAULT 'unpaid', -- unpaid, paid, partially_paid, refunded    
    txn_accounted BOOLEAN NOT NULL DEFAULT FALSE,

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__SalesLines(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    info TEXT NOT NULL DEFAULT '',
    qty INTEGER NOT NULL DEFAULT 0,    
    sale_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL DEFAULT 0,

    price INTEGER NOT NULL DEFAULT 0, -- original item price
    tax_amount INTEGER NOT NULL DEFAULT 0,
    discount_amount INTEGER NOT NULL DEFAULT 0,

    total_amount INTEGER NOT NULL DEFAULT 0, -- total_amount = item_price + item_tax_amount - discount_amount
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table __project__Tax(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL DEFAULT '',
    ttype TEXT NOT NULL DEFAULT 'item_percent', -- item_percent, category_percent
    info TEXT NOT NULL DEFAULT '',
    rate INTEGER NOT NULL DEFAULT 0,
    strict BOOLEAN NOT NULL DEFAULT FALSE,
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


create table __project__ProductTaxes(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    catagory_id INTEGER NOT NULL DEFAULT 0,
    product_id INTEGER NOT NULL DEFAULT 0,
    tax_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table __project__Contacts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
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
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__Notepads(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    note_type TEXT NOT NULL DEFAULT 'normal',
    attachments TEXT NOT NULL DEFAULT '',
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    linked_client_id INTEGER NOT NULL DEFAULT 0
);