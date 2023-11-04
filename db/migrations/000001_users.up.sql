BEGIN;

CREATE TYPE roles AS ENUM ('participant', 'admin');

CREATE TABLE IF NOT EXISTS users(
   id UUID PRIMARY KEY,
   username VARCHAR(50) NOT NULL UNIQUE,
   role roles default 'participant'::roles,
   created_at timestamp default now(),
   updated_at timestamp default now(),
   created_by UUID
);

COMMIT;