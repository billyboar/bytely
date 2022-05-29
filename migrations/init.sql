CREATE TABLE IF NOT EXISTS urls (
  short_url_key varchar(7) PRIMARY KEY,
  created_at timestamp NOT NULL DEFAULT now(),
  original_url text,
  clicks integer DEFAULT 0
);