package handler

import (
	"go_project_structure_be/servers"
)

type Handler struct {
	server *servers.Server
}

func NewHandler(sv *servers.Server) Handler {
	return Handler{server: sv}
}
