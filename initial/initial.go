package initial

import (
	"fmt"
	"go_project_structure_be/configurations"
	"go_project_structure_be/util"
	"net/url"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type initial struct {
	confs *configurations.Configs
}

type Initial interface {
	Database() (*sqlx.DB, error)
	WebsocketAsterisk() (*websocket.Conn, error)
}

func NewInitial(confs *configurations.Configs) Initial {
	return initial{confs: confs}
}

func (init initial) Database() (*sqlx.DB, error) {
	db := init.confs.ConfEnv
	dbStr := ""
	switch db.Provider {
	case "mysql":
		dbStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", db.User, db.Password, db.Host, db.DBPort, db.Dbname)
		database, err := sqlx.Connect(db.Provider, dbStr)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		for _, item := range init.confs.Conf.InitialScripts {
			sqlFile, err := os.ReadFile(item)
			if err != nil {
				log.Errorf("[%v]:%v", item, err)
			}

			queries := string(sqlFile)
			querySplit := strings.Split(queries, ";")

			for _, query := range querySplit {
				if util.StringIsEmpty(query) {
					continue
				}
				_, err = database.Exec(query)
				if err != nil {
					log.Errorf("[%v]:%v", item, err)
				}
			}

		}

		database.SetConnMaxLifetime(time.Duration(db.ConnectionMaxLifeTime) * time.Second)
		database.SetMaxIdleConns(db.MaxIdleConns)
		database.SetMaxOpenConns(db.MaxOpenConns)
		database.SetConnMaxIdleTime(time.Duration(db.MAX_IDLE_TIME) * time.Second)

		return database, nil

	case "postgres":
		dbStr = fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v", db.User, db.Password, db.Host, db.DBPort, db.Dbname, db.SSLMODE)
		database, err := sqlx.Connect(db.Provider, dbStr)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		for _, item := range init.confs.Conf.InitialScripts {
			sqlFile, err := os.ReadFile(item)
			if err != nil {
				log.Errorf("[%v]:%v", item, err)
			}

			queries := string(sqlFile)
			_, err = database.Exec(queries)
			if err != nil {
				log.Errorf("[%v]:%v", item, err)
			}

		}

		database.SetConnMaxLifetime(time.Duration(db.ConnectionMaxLifeTime) * time.Second)
		database.SetMaxIdleConns(db.MaxIdleConns)
		database.SetMaxOpenConns(db.MaxOpenConns)
		database.SetConnMaxIdleTime(time.Duration(db.MAX_IDLE_TIME) * time.Second)

		return database, nil
	}

	return nil, fmt.Errorf("not found db %v", db.Provider)

}

func (init initial) WebsocketAsterisk() (*websocket.Conn, error) {
	wsURL := url.URL{
		Scheme:   "ws",
		Host:     "",
		Path:     "/ari/events",
		RawQuery: fmt.Sprintf("api_key=%s&app=%s", "", ""),
	}

	conn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err != nil {
		log.Fatal("WebSocket connection error:", err)
	}

	return conn, err
}
