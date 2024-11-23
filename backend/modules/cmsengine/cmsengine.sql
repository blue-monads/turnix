

create table __project__Pages(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT NOT NULL DEFAULT '',
    title TEXT NOT NULL DEFAULT '',
    content TEXT NOT NULL DEFAULT '',
    created_by INTEGER NOT NULL,
    updated_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    layout TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'published', -- published, draft
);

create table __project__Blocks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT NOT NULL DEFAULT '',
    title TEXT NOT NULL DEFAULT '',
    style TEXT NOT NULL DEFAULT '',
);
