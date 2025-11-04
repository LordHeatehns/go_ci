package repository

import (
	"go_ci/servers"

	"github.com/jmoiron/sqlx"
)

type RepoAction interface {
	BeginX() (*sqlx.Tx, error)
	UsersRepository() UsersRepository
}

type repository struct {
	db *sqlx.DB
}

type Option func(repo *repository) error

func NewRepository(server *servers.Server, options ...Option) (RepoAction, error) {
	repo := &repository{}

	for _, option := range options {
		if err := option(repo); err != nil {
			return nil, err
		}
	}

	if repo.db == nil {
		repo.db = server.DB
	}

	return repo, nil
}

func (repo *repository) BeginX() (*sqlx.Tx, error) {
	tx, err := repo.db.Beginx()
	return tx, err
}

// func queryOptionBuilder(searchQuery squirrel.SelectBuilder, option bindings.IQueryOption) squirrel.SelectBuilder {
// 	if option != nil {
// 		if option.GetLimit() > 0 {
// 			searchQuery = searchQuery.Limit(uint64(option.GetLimit()))
// 		}
// 		if option.GetOffset() > 0 {
// 			searchQuery = searchQuery.Offset(uint64(option.GetOffset()))
// 		}

// 		//sorting
// 		if len(option.GetSort()) > 0 {
// 			orders := make([]string, 0)
// 			for _, order := range option.GetSort() {
// 				orderStr := fmt.Sprintf("%s %s", order.GetField(), string(order.GetOrder()))
// 				orders = append(orders, orderStr)
// 			}
// 			searchQuery = searchQuery.OrderBy(orders...)
// 		}
// 	}

// 	return searchQuery
// }
