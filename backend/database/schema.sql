CREATE TABLE IF NOT EXISTS Users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  utype TEXT NOT NULL DEFAULT 'real', -- real, device
  email TEXT,
  phone TEXT NOT NULL DEFAULT '',
  bio TEXT NOT NULL DEFAULT '',
  password TEXT NOT NULL,
  isEmailVerified BOOLEAN NOT NULL DEFAULT FALSE,
  extrameta JSON NOT NULL DEFAULT '{}',
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  ownedBy INTEGER NOT NULL DEFAULT 0,
  accessToken TEXT NOT NULL DEFAULT '',
  unique(email)
);

CREATE TABLE IF NOT EXISTS Projects (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL DEFAULT '',
  ptype TEXT NOT NULL DEFAULT 'onloop',
  owner INTEGER NOT NULL,
  extrameta JSON NOT NULL DEFAULT '{}',
  FOREIGN KEY (owner) REFERENCES Users(id)
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