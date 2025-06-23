package routers

import (
	"go_project_structure_be/handler"
	"go_project_structure_be/servers"

	"github.com/robfig/cron/v3"
)

type CronRegistry interface {
}

type cronRegistry struct {
	hand   *handler.Handler
	server servers.Server
	cron   *cron.Cron
}

func NewCronRegistry(hand *handler.Handler, server servers.Server, cron *cron.Cron) CronRegistry {
	return &cronRegistry{hand: hand, server: server, cron: cron}
}
