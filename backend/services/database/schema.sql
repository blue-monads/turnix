
CREATE TABLE IF NOT EXISTS GlobalConfig (
  id INTEGER PRIMARY KEY, 
  key TEXT NOT NULL DEFAULT '', 
  group_name TEXT NOT NULL DEFAULT '',
  value TEXT NOT NULL DEFAULT '',
  unique(group_name, key)
);


CREATE TABLE IF NOT EXISTS Users (
  id INTEGER PRIMARY KEY, 
  username TEXT,
  email TEXT, 
  phone TEXT,

  name TEXT NOT NULL, 
  utype TEXT NOT NULL DEFAULT 'real',  -- super, real, device
  bio TEXT NOT NULL DEFAULT '', 
  password TEXT NOT NULL, 
  is_verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  owner_user_id INTEGER NOT NULL DEFAULT 0,
  owner_project_id INTEGER NOT NULL DEFAULT 0,
  extrameta JSON NOT NULL DEFAULT '{}',
  msg_read_head INTEGER NOT NULL DEFAULT 0,
  
  disabled BOOLEAN NOT NULL DEFAULT FALSE, 
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  unique(username),  
  unique(email),
  unique(phone)
);

CREATE TABLE IF NOT EXISTS UserConfig (
  id INTEGER PRIMARY KEY, 
  key TEXT NOT NULL DEFAULT '', 
  group_name TEXT NOT NULL DEFAULT '',
  value TEXT NOT NULL DEFAULT '',
  user_id INTEGER NOT NULL, 
  unique(user_id, group_name, key),
  FOREIGN KEY (user_id) REFERENCES Users(id)
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

CREATE TABLE IF NOT EXISTS ProjectConfig (
  id INTEGER PRIMARY KEY, 
  key TEXT NOT NULL DEFAULT '', 
  group_name TEXT NOT NULL DEFAULT '',
  value TEXT NOT NULL DEFAULT '',
  project_id INTEGER NOT NULL, 
  unique(project_id, group_name, key),
  FOREIGN KEY (project_id) REFERENCES Projects(id)
);

CREATE TABLE IF NOT EXISTS ProjectHooks (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  event TEXT NOT NULL,
  order_id INTEGER NOT NULL DEFAULT 0,
  runas_user_id INTEGER NOT NULL DEFAULT -1,
  hook_type TEXT NOT NULL DEFAULT 'script', -- script, webhook, email
  hook_code TEXT NOT NULL DEFAULT '',
  target TEXT NOT NULL DEFAULT '',
  envs JSON NOT NULL DEFAULT '{}',
  project_id INTEGER NOT NULL, 
  extrameta JSON NOT NULL DEFAULT '{}',
  FOREIGN KEY (project_id) REFERENCES Projects(id)
);

CREATE TABLE IF NOT EXISTS ProjectPlugins (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  ptype TEXT NOT NULL DEFAULT 'sideapp_v1', -- sideapp, inject 
  project_id INTEGER NOT NULL,
  server_code TEXT NOT NULL DEFAULT '',
  client_code TEXT NOT NULL DEFAULT '',
  created_by INTEGER NOT NULL,
  updated_by INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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

CREATE TABLE IF NOT EXISTS Files (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  ftype TEXT NOT NULL DEFAULT 'project', -- project, user
  is_folder BOOLEAN NOT NULL DEFAULT FALSE,
  path TEXT NOT NULL DEFAULT '', 
  size INTEGER NOT NULL DEFAULT 0, 
  mime TEXT NOT NULL DEFAULT '', 
  hash TEXT NOT NULL DEFAULT '',
  storeType INTEGER NOT NULL DEFAULT 0, -- 0: inline_blob, 1: external_blob, 2: mulit_part_blob
  preview BLOB, 
  blob BLOB,
  external BOOLEAN NOT NULL DEFAULT FALSE,
  owner_user_id INTEGER NOT NULL DEFAULT 0,
  owner_project_id INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (owner_user_id, owner_project_id, path, name)
);



CREATE TABLE IF NOT EXISTS FilePartedBlobs (
  id INTEGER PRIMARY KEY,
  file_id INTEGER NOT NULL,
  size INTEGER NOT NULL,
  part_id INTEGER NOT NULL,
  blob BLOB NOT NULL
);



CREATE TABLE IF NOT EXISTS FileShares (
  id TEXT PRIMARY KEY,
  file_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);