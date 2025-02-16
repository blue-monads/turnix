create table __project__Sites(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL DEFAULT '',
    api_key TEXT NOT NULL DEFAULT '',
    provider TEXT NOT NULL DEFAULT '', -- github pages, netlify, vercel, etc
    domain TEXT NOT NULL DEFAULT '',
    base_path TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_deployed_at TIMESTAMP,
    deploy_webhook TEXT NOT NULL DEFAULT '',
    deploy_branch TEXT NOT NULL DEFAULT '',
    deploy_repo TEXT NOT NULL DEFAULT '',
    hugo_config TEXT NOT NULL DEFAULT ''
);


create table __project__Posts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT NOT NULL DEFAULT '',
    title TEXT NOT NULL DEFAULT '',
    content TEXT NOT NULL DEFAULT '',
    author_id INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_published BOOLEAN NOT NULL DEFAULT FALSE
);

create table __project__Attachments(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL DEFAULT 0,
    file_id INTEGER NOT NULL DEFAULT 0
);

create table __project__Tags(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


create table __project__PostTags(
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (post_id, tag_id)
);