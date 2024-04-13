create table AccountType(
    id INTEGER PRIMARY KEY,
    slug TEXT NOT NULL DEFAULT '',
    unique(slug)
);

create table Accounts(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL DEFAULT '',
    acc_type TEXT NOT NULL DEFAULT '',
    total_debit INTEGER NOT NULL DEFAULT 0,
    total_credit INTEGER NOT NULL DEFAULT 0,
    contact_id INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL FALSE
);

create table TransactionType(
    id INTEGER PRIMARY KEY,
    slug TEXT NOT NULL DEFAULT '',
    unique(slug)
);

create table Transactions(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    account_id INTEGER NOT NULL,
    txn_group_id INTEGER NULL NULL,
    debit_amount INTEGER NOT NULL DEFAULT 0,
    credit_amount INTEGER NOT NULL DEFAULT 0,
    credit_amount INTEGER NOT NULL DEFAULT 0,
    linked_sales_id INTEGER NOT NULL DEFAULT 0,
    linked_invoice_id INTEGER NOT NULL DEFAULT 0,
    reference_id TEXT NOT NULL DEFAULT '',
    attachments TEXT NOT NULL DEFAULT '',
    created_by INTEGER NULL NULL,
    updated_by INTEGER NULL NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL FALSE
);

create table ReportTemplates(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    input_options JSON NOT NULL DEFAULT '{}',
    output_options JSON NOT NULL DEFAULT '{}',
    query_template TEXT NOT NULL DEFAULT '',
    filter_script TEXT NULL NULL DEFAULT ''
);

create table Reports(
    id INTEGER PRIMARY KEY,
    template_id INTEGER NOT NULL DEFAULT 0,
    result JSON NOT NULL DEFAULT '{}',
    created_by INTEGER NULL NULL,
    updated_by INTEGER NULL NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL FALSE
);