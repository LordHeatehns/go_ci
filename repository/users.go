package repository

import "go_ci/responses"

type UsersRepository interface {
	GetUser() (*responses.Users, error)
}

func (r *repository) UsersRepository() UsersRepository {
	return r
}

func (r *repository) GetUser() (*responses.Users, error) {
	mock := responses.Users{
		Username: "nus",
		Email:    "nus@gmaill.com",
	}

	return &mock, nil
}
