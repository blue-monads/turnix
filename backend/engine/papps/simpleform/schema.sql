

create table Forms(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    submitter_type TEXT NOT NULL, -- user, user_or_guest, verified_guest, guest
    post_trackable BOOLEAN NOT NULL DEFAULT FALSE,
    post_editable BOOLEAN NOT NULL DEFAULT FALSE,
    enable_captcha BOOLEAN NOT NULL DEFAULT FALSE,
    style_override TEXT NOT NULL DEFAULT '',
    script_override TEXT NOT NULL DEFAULT '',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table FormLayouts(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    title TEXT NOT NULL,
    ltype TEXT NOT NULL,
    html_attrs JSON NOT NULL DEFAULT '{}',
    parent_id INTEGER NOT NULL DEFAULT 0,
);


create table FormElements(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    title TEXT NOT NULL,
    elem_type TEXT NOT NULL, -- text, number, range, colorpicker, date, time, array
    extrameta JSON NOT NULL DEFAULT '{}',
    html_attrs JSON NOT NULL DEFAULT '{}',
    required BOOLEAN NOT NULL DEFAULT FALSE,
    options JSON NOT NULL DEFAULT '{}',
    layout_id INTEGER NOT NULL DEFAULT FALSE
);


create table FormSubmission(
    id INTEGER PRIMARY KEY,
    submitter INTEGER NOT NULL DEFAULT 0,
    locked BOOLEAN NOT NULL FALSE,
    submitter_token TEXT NULL NULL DEFAULT '',
    data JSON TEXT NOT NULL '{}',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

