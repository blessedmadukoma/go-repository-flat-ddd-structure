package repository

import (
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/validators"
	"log"

	"github.com/gin-gonic/gin"
)

type AccountsRepository interface {
	GetAccounts(ctx *gin.Context, req validators.ListAccountInput) ([]database.Account, error)
	GetAccountByID(ctx *gin.Context, id int64) (database.Account, error)
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
		return database.Account{}, err
	}

	return account, nil
}
