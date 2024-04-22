package repository

import (
	database "goRepositoryPattern/database/sqlc"
)

type Repository struct {
	DB database.Store
}

func NewRepository(db database.Store) *Repository {
	return &Repository{
		DB: db,
	}
}

// func (repo Repository) Paginator() *paginate.Pagination {
// 	return paginate.New(&paginate.Config{
// 		DefaultSize:          10,
// 		FieldSelectorEnabled: true,
// 	})
// }
