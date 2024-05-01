package services

import (
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/validators"

	"github.com/gin-gonic/gin"
)

type AccountsService struct {
	repo repository.AccountsRepository
}

func NewAccountService(repo repository.AccountsRepository) AccountsService {
	return AccountsService{
		repo: repo,
	}
}

func (as *AccountsService) GetAccounts(ctx *gin.Context, req validators.ListAccountInput) ([]database.Account, error) {

	accounts, err := as.repo.GetAccounts(ctx, req)
	if err != nil {
		return []database.Account{}, err
	}

	return accounts, nil
}

func (as *AccountsService) GetAccountByID(ctx *gin.Context, id int64) (database.Account, error) {

	account, err := as.repo.GetAccountByID(ctx, id)
	if err != nil {
		return database.Account{}, err
	}

	return account, nil
}

func (as *AccountsService) GetAccountByEmail(ctx *gin.Context, email string) (database.Account, error) {

	account, err := as.repo.GetAccountByEmail(ctx, email)
	if err != nil {
		return database.Account{}, err
	}

	return account, nil
}
