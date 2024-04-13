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
