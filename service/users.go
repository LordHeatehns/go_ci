package service

import (
	"go_ci/responses"
)

func (s *service) GetUsers() (*responses.Users, error) {

	datas, err := s.repo.UsersRepository().GetUser()
	if err != nil {
		return nil, err
	}

	return datas, err
}
