CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "firstname" varchar(256) NOT NULL,
  "lastname" varchar(256) NOT NULL,
  "email" varchar(256) UNIQUE NOT NULL,
  "hashed_password" varchar(256) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
)