// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateOtp(ctx context.Context, arg CreateOtpParams) (AccountOtp, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAllAccounts(ctx context.Context) error
	DeleteAllOtps(ctx context.Context) error
	DeleteOtp(ctx context.Context, arg DeleteOtpParams) error
	GetAccountByEmail(ctx context.Context, email string) (Account, error)
	GetAccountByID(ctx context.Context, id int64) (Account, error)
	GetOtpByAccountID(ctx context.Context, accountID int64) (AccountOtp, error)
	GetOtpByAccountIDAndType(ctx context.Context, arg GetOtpByAccountIDAndTypeParams) (AccountOtp, error)
	GetOtpByID(ctx context.Context, id int64) (AccountOtp, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListOtps(ctx context.Context, arg ListOtpsParams) ([]AccountOtp, error)
	UpdateAccountPassword(ctx context.Context, arg UpdateAccountPasswordParams) (Account, error)
	UpdateAccountStatus(ctx context.Context, arg UpdateAccountStatusParams) (Account, error)
	UpdateOtp(ctx context.Context, arg UpdateOtpParams) (AccountOtp, error)
}

var _ Querier = (*Queries)(nil)
