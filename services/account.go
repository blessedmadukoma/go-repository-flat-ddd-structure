package services

import (
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/validators"
	"log"

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
		log.Fatal("error getting accounts:", err)
		return []database.Account{}, err
	}

	return accounts, nil
}
