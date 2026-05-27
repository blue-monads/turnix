
CREATE TABLE IF NOT EXISTS MQSubscriptions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  install_id INTEGER NOT NULL,
  space_id INTEGER NOT NULL DEFAULT 0,
  event_key TEXT NOT NULL DEFAULT '',
  target_type TEXT NOT NULL DEFAULT '', -- webhook, script, space_method
  target_space_id INTEGER NOT NULL DEFAULT 0,
  target_endpoint TEXT NOT NULL DEFAULT '',
  target_options JSON NOT NULL DEFAULT '{}', -- it has creds, api keys and other options
  target_code TEXT NOT NULL DEFAULT '',
  rules JSON NOT NULL DEFAULT '{}',
  transform JSON NOT NULL DEFAULT '{}',
  delay_start INTEGER NOT NULL DEFAULT 0,
  retry_delay INTEGER NOT NULL DEFAULT 0,
  max_retries INTEGER NOT NULL DEFAULT 0,
  expires_on INTEGER NOT NULL DEFAULT 0, -- (created_at + expires_in > now) then status is expired
  collapse_interval INTEGER NOT NULL DEFAULT 0, -- 1 minute, 5 minute, 15 minute etc in seconds
  extrameta JSON NOT NULL DEFAULT '{}',
  created_by INTEGER NOT NULL DEFAULT 0,
  disabled BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- utc timestamp rounded to nearsest interval

CREATE TABLE IF NOT EXISTS MQEvents (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  install_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  payload BLOB NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status TEXT NOT NULL DEFAULT 'new', -- new, scheduled, processed
  extrameta JSON NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS MQEventTargets (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  collapse_key TEXT NOT NULL DEFAULT '',
  event_id INTEGER NOT NULL,
  subscription_id INTEGER NOT NULL,
  status TEXT NOT NULL DEFAULT 'new', -- new, processing, start_delayed, delayed, processed, failed, expired
  delayed_until INTEGER NOT NULL DEFAULT 0,
  retry_count INTEGER NOT NULL DEFAULT 0,
  error TEXT NOT NULL DEFAULT '',
  extrameta JSON NOT NULL DEFAULT '{}',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);