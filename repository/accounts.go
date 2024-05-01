package repository

import (
	"errors"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/validators"
	"log"

	"github.com/gin-gonic/gin"
)

type AccountsRepository interface {
	GetAccounts(ctx *gin.Context, req validators.ListAccountInput) ([]database.Account, error)
	GetAccountByID(ctx *gin.Context, id int64) (database.Account, error)
	GetAccountByEmail(ctx *gin.Context, email string) (database.Account, error)
	DeleteAccount(ctx *gin.Context, id int64) error
}

func (r *Repository) GetAccounts(ctx *gin.Context, req validators.ListAccountInput) ([]database.Account, error) {
	args := database.ListAccountsParams{
		Limit:  int32(req.PageSize),
		Offset: int32((req.PageNumber - 1) * req.PageSize),
	}

	// list accounts
	accounts, err := r.DB.ListAccounts(ctx, args)
	if err != nil {
		return []database.Account{}, err
	}

	log.Println("Account length:", len(accounts))
	for _, v := range accounts {
		log.Println("Account:", v)
	}

	return accounts, nil
}

func (r Repository) GetAccountByID(ctx *gin.Context, id int64) (database.Account, error) {
	account, err := r.DB.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return database.Account{}, messages.ErrUserNotExists
		}

		return database.Account{}, err
	}

	return account, nil
}

func (r Repository) GetAccountByEmail(ctx *gin.Context, email string) (database.Account, error) {
	account, err := r.DB.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return database.Account{}, messages.ErrUserNotExists
		}

		return database.Account{}, err
	}

	return account, nil
}

func (r Repository) DeleteAccount(ctx *gin.Context, id int64) error {
	_, err := r.DB.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return messages.ErrUserNotExists
		}

		return err
	}

	err = r.DB.DeleteAccount(ctx, id)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return messages.ErrUserNotExists
		}

		return err
	}

	return nil
}
