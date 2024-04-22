package repository

import (
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/token"
)

type Repository struct {
	DB    database.Store
	Token *token.JWTToken
}

func NewRepository(db database.Store, token *token.JWTToken) *Repository {
	return &Repository{
		DB:    db,
		Token: token,
	}
}

// func (repo Repository) Paginator() *paginate.Pagination {
// 	return paginate.New(&paginate.Config{
// 		DefaultSize:          10,
// 		FieldSelectorEnabled: true,
// 	})
// }
