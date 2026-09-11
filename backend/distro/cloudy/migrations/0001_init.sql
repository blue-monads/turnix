CREATE TABLE IF NOT EXISTS CloudyUsers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fullname TEXT NOT NULL,
    email TEXT UNIQUE,
    password TEXT,
    tenant_key TEXT NOT NULL UNIQUE,
    utype TEXT NOT NULL DEFAULT 'normal' CHECK (utype IN ('admin', 'normal')),
    pricing_tier TEXT NOT NULL DEFAULT 'free',
    is_verified INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
