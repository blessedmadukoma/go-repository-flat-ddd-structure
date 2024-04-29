CREATE TABLE IF NOT EXISTS account_otps (
    id bigserial NOT NULL PRIMARY KEY,
    account_id bigint NOT NULL,
    otp text NOT NULL,
    type bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at  timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_accounts_account_otps FOREIGN KEY(account_id) REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE
);