package handler

import (
	"go_project_structure_be/servers"
	"go_project_structure_be/service"
)

type Handler struct {
	server *servers.Server
	sv     *service.MockService
}

func NewHandler(sv *servers.Server) Handler {
	return Handler{server: sv}
}

func NewHandlerMock(sv *service.MockService) Handler {
	return Handler{sv: sv}
}
