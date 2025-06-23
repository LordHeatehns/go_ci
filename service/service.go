package service

import (
	"go_project_structure_be/configurations"
	"go_project_structure_be/repository"
	"go_project_structure_be/servers"
)

type service struct {
	repo  repository.RepoAction
	confs *configurations.Configs
}

type Option func(sv *service) error

func New(server *servers.Server, options ...Option) (*service, error) {
	sv := &service{}

	for _, option := range options {
		if err := option(sv); err != nil {
			return nil, err
		}
	}

	if sv.confs == nil {
		sv.confs = server.Conf
	}

	// db
	if sv.repo == nil {
		repo, err := repository.NewRepository(server)
		if err != nil {
			return nil, err
		}

		sv.repo = repo

	}

	return sv, nil
}
