-- name: CreateAccount :one
INSERT INTO accounts (
  firstname,
  lastname,
  email,
  hashed_password
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccountByEmail :one
SELECT * FROM accounts WHERE email = $1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAccountPassword :one
UPDATE accounts SET hashed_password = $1, updated_at = now()
WHERE id = $2 RETURNING *;

-- name: UpdateAccountStatus :one
UPDATE accounts SET is_verified = $1, updated_at = now() WHERE id = $2 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: DeleteAllAccounts :exec
DELETE FROM accounts;