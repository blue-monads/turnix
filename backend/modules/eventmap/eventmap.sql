

create table EventTypes(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    event_type TEXT NOT NULL 'e',
    icon TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);



create table Events(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    info TEXT NOT NULL  DEFAULT '',
    event_type TEXT NULL 'e',
    event_data JSON NOT NULL '{}',
    lat INTEGER NOT NULL DEFAULT 0,
    lng INTEGER NOT NULL DEFAULT 0,
    event_start TIMESTAMP,
    event_end TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);