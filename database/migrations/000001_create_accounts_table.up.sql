-- CREATE TYPE role AS ENUM (
--   'admin',
--   'user'
-- );
-- Create a PL/pgSQL function to create the role enum type if it does not exist
CREATE OR REPLACE FUNCTION create_role_enum_if_not_exists() RETURNS VOID AS $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
    CREATE TYPE role AS ENUM ('admin', 'user');
  END IF;
END;
$$ LANGUAGE plpgsql;

-- Call the function to create the role enum type if it does not exist
SELECT create_role_enum_if_not_exists();

CREATE TABLE "accounts" (
  "id" BIGSERIAL PRIMARY KEY,
  "firstname" varchar(256) NOT NULL,
  "lastname" varchar(256) NOT NULL,
  "email" varchar(256) UNIQUE NOT NULL,
  "role" role NOT NULL DEFAULT 'user',
  "is_verified" BOOLEAN DEFAULT FALSE,
  "hashed_password" varchar(256) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
)