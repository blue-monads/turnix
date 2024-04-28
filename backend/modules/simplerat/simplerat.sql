-- simplectl
create table DeviceBeacons(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    device_id INTEGER NOT NULL,
    lat INTEGER NOT NULL DEFAULT 0,
    lng INTEGER NOT NULL DEFAULT 0,
    beacon_type TEXT NULL 'p',
    beacon_data JSON NOT NULL '{}',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table Jobs(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    job_type TEXT NULL 's',
    contents TEXT NULL '',
    rule TEXT NULL '',
    extrameta JSON NOT NULL '{}',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table DeviceJobs(
    id INTEGER PRIMARY KEY,
    job_id INTEGER NOT NULL DEFAULT 0,
    device_id INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table JobOuputs(
    id INTEGER PRIMARY KEY,
    device_job_id INTEGER NOT NULL DEFAULT 0,
    data JSON NOT NULL '{}',
    result_data JSON NOT NULL '{}',
    device_id INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
