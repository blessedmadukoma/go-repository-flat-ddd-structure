package models

import (
	database "goRepositoryPattern/database/sqlc"
	"time"
)

type ListAccountResponse struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Token     string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListAccountsResponse struct {
	Accounts []ListAccountsResponse `json:"accounts"`
}

func NewAccountResponse(account database.Account) ListAccountResponse {
	return ListAccountResponse{
		ID:        account.ID,
		FirstName: account.Firstname,
		LastName:  account.Lastname,
		Email:     account.Email,
		Role:      string(account.Role),
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
