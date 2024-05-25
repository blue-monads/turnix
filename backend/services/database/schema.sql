CREATE TABLE IF NOT EXISTS Users (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL, 
  utype TEXT NOT NULL DEFAULT 'real',  -- real, device
  email TEXT, 
  phone TEXT NOT NULL DEFAULT '', 
  bio TEXT NOT NULL DEFAULT '', 
  password TEXT NOT NULL, 
  email_verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  owner_user_id INTEGER NOT NULL DEFAULT 0,
  owner_project_id INTEGER NOT NULL DEFAULT 0,

  disabled BOOLEAN NOT NULL DEFAULT FALSE, 
  msg_read_head INTEGER NOT NULL DEFAULT 0,
  extrameta JSON NOT NULL DEFAULT '{}',
  unique(email, owner_project_id)
);

CREATE TABLE IF NOT EXISTS UserDevices (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  info TEXT NOT NULL DEFAULT '', 
  dtype TEXT NOT NULL DEFAULT 'sesssion', --  session token
  hash TEXT NOT NULL DEFAULT '', 
  contents TEXT NOT NULL DEFAULT '', 
  user_id INTEGER NOT NULL, 
  project_id INTEGER NOT NULL DEFAULT 0,
  extrameta JSON NOT NULL DEFAULT '{}', 
  expires_on TIMESTAMP not null, 
  FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS UserMessages(
  id INTEGER PRIMARY KEY, 
  title text not null default '', 
  is_read boolean not null default FALSE, 
  type text not null default "messsage", 
  contents text not null, 
  to_user text not null, 
  from_user_id text not null default 0, 
  from_project_id text not null default 0, 
  callback_token text not null default 0, 
  warn_level integer not null default 0, 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  FOREIGN KEY (to_user) REFERENCES Users(id), 
  FOREIGN KEY (from_user_id) REFERENCES Users(id), 
  FOREIGN KEY (from_project_id) REFERENCES Projects(id)
);

CREATE TABLE IF NOT EXISTS Projects (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  info TEXT NOT NULL DEFAULT '', 
  ptype TEXT NOT NULL DEFAULT 'onloop', 
  owned_by INTEGER NOT NULL, 
  extrameta JSON NOT NULL DEFAULT '{}', 
  is_initilized BOOLEAN NOT NULL DEFAULT FALSE, 
  is_public BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (owned_by) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS ProjectHooks (
  id INTEGER PRIMARY KEY, 
  event TEXT NOT NULL,
  hook_type TEXT NOT NULL DEFAULT 'script', -- script, webhook, email
  hook_code TEXT NOT NULL DEFAULT '',
  envs JSON NOT NULL DEFAULT '{}',
  project_id INTEGER NOT NULL, 
  extrameta JSON NOT NULL DEFAULT '{}',
  FOREIGN KEY (project_id) REFERENCES Projects(id)
);


CREATE TABLE IF NOT EXISTS ProjectUsers (
  id INTEGER PRIMARY KEY, 
  user_id INTEGER NOT NULL, 
  project_id INTEGER NOT NULL, 
  scope TEXT NOT NULL DEFAULT '', 
  extrameta JSON NOT NULL DEFAULT '{}', 
  token TEXT NOT NULL DEFAULT '', 
  FOREIGN KEY (project_id) REFERENCES Projects(id), 
  FOREIGN KEY (user_id) REFERENCES Users(id), 
  unique(project_id, user_id)
);
