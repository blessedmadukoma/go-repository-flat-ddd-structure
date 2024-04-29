-- name: CreateOtp :one
INSERT INTO account_otps (
  account_id,
otp,
type
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetOtpByID :one
SELECT * FROM account_otps WHERE id = $1;

-- name: GetOtpByAccountID :one
SELECT * FROM account_otps WHERE account_id = $1;

-- name: GetOtpByAccountIDAndType :one
SELECT * FROM account_otps WHERE account_id = $1 AND type = $2;

-- name: ListOtps :many
SELECT * FROM account_otps ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateOtp :one
UPDATE account_otps SET otp = $1, updated_at = now()
WHERE account_id = $2 AND type = $3 RETURNING *;

-- name: DeleteOtp :exec
DELETE FROM account_otps WHERE id = $1 AND account_id = $2 AND type = $3;

-- name: DeleteAllOtps :exec
DELETE FROM account_otps;