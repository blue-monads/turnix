CREATE TABLE IF NOT EXISTS Users (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL, 
  utype TEXT NOT NULL DEFAULT 'real',  -- real, device
  email TEXT, 
  phone TEXT NOT NULL DEFAULT '', 
  bio TEXT NOT NULL DEFAULT '', 
  password TEXT NOT NULL, 
  isEmailVerified BOOLEAN NOT NULL DEFAULT FALSE, 
  extrameta JSON NOT NULL DEFAULT '{}', 
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  ownedBy INTEGER NOT NULL DEFAULT 0, 
  accessToken TEXT NOT NULL DEFAULT '', 
  disabled BOOLEAN NOT NULL DEFAULT FALSE, 
  messageReadHead INTEGER NOT NULL DEFAULT 0, 
  unique(email)
);

CREATE TABLE IF NOT EXISTS UserDevices (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  info TEXT NOT NULL DEFAULT '', 
  dtype TEXT NOT NULL DEFAULT 'sesssion', --  session token
  contentHash TEXT NOT NULL DEFAULT '', 
  contents TEXT NOT NULL DEFAULT '', 
  ownedBy INTEGER NOT NULL, 
  pinnedProjectId INTEGER NOT NULL DEFAULT 0,
  extrameta JSON NOT NULL DEFAULT '{}', 
  expiresOn TIMESTAMP not null, 
  FOREIGN KEY (ownedBy) REFERENCES Users(id)
);

create table UserMessages(
  id INTEGER PRIMARY KEY, 
  title text not null default '', 
  isRead boolean not null default false, 
  type text not null default "messsage", 
  contents text not null, 
  toUser text not null, 
  fromUser text not null default 0, 
  fromProject text not null default 0, 
  callbackToken text not null default 0, 
  warnLevel integer not null default 0, 
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  FOREIGN KEY (toUser) REFERENCES Users(id), 
  FOREIGN KEY (fromUser) REFERENCES Users(id), 
  FOREIGN KEY (fromProject) REFERENCES Projects(id)
);

CREATE TABLE IF NOT EXISTS Projects (
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  info TEXT NOT NULL DEFAULT '', 
  ptype TEXT NOT NULL DEFAULT 'onloop', 
  owner INTEGER NOT NULL, 
  extrameta JSON NOT NULL DEFAULT '{}', 
  isInitilized BOOLEAN NOT NULL DEFAULT FALSE, 
  FOREIGN KEY (owner) REFERENCES Users(id)
  isPublic BOOLEAN NOT NULL DEFAULT FALSE,
);

CREATE TABLE IF NOT EXISTS ProjectUsers (
  id INTEGER PRIMARY KEY, 
  userId INTEGER NOT NULL, 
  extrameta JSON NOT NULL DEFAULT '{}', 
  projectId INTEGER NOT NULL, 
  accessToken TEXT NOT NULL DEFAULT '', 
  FOREIGN KEY (projectId) REFERENCES Projects(id), 
  FOREIGN KEY (userId) REFERENCES Users(id), 
  unique(projectId, userId)
);
CREATE TABLE IF NOT EXISTS pt_onloop_Templates(
  id INTEGER PRIMARY KEY, 
  name TEXT NOT NULL DEFAULT '', 
  ttype TEXT NOT NULL DEFAULT 'bexec', 
  canReadResult BOOLEAN NOT NULL DEFAULT FALSE, 
  consts JSON NOT NULL DEFAULT '{}', 
  contentScript TEXT NOT NULL DEFAULT '', 
  extrameta JSON NOT NULL DEFAULT '{}', 
  projectId INTEGER NOT NULL, 
  FOREIGN KEY (projectId) REFERENCES Projects(id)
);
CREATE TABLE IF NOT EXISTS pt_onloop_QueueMessages(
  id INTEGER PRIMARY KEY, 
  submitter INTEGER NOT NULL, 
  contents JSON NOT NULL DEFAULT '{}', 
  result JSON NOT NULL DEFAULT '{}', 
  status TEXT NOT NULL, 
  projectId INTEGER NOT NULL, 
  templateId INTEGER NOT NULL, 
  extrameta JSON NOT NULL DEFAULT '{}', 
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  FOREIGN KEY (templateId) REFERENCES Templates(id), 
  FOREIGN KEY (projectId) REFERENCES Projects(id), 
  FOREIGN KEY (submitter) REFERENCES Users(id)
);