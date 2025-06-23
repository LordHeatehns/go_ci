package servers

import (
	"go_project_structure_be/configurations"
	"go_project_structure_be/initial"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DB             *sqlx.DB
	Conf           *configurations.Configs
	WebseocketConn *websocket.Conn
}

func NewServer(init initial.Initial, conf *configurations.Configs) (*Server, error) {
	sv := &Server{}
	sv.Conf = conf

	// db, err := init.Database()
	// if err != nil {
	// 	return nil, err
	// }

	// sv.DB = db

	return sv, nil
}
