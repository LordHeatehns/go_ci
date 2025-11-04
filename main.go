package main

import (
	"fmt"
	"go_project_structure_be/configurations"
	"go_project_structure_be/handler"
	"go_project_structure_be/initial"
	middlewares "go_project_structure_be/middleware"
	"go_project_structure_be/routers"
	"go_project_structure_be/servers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	"github.com/robfig/cron/v3"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conf, err := configurations.LoadConfigFile()
	if err != nil {
		log.Error(err)
	}

	mw := middlewares.NewMiddleWare(app, *conf)
	mw.Logger()
	mw.Cors()
	init := initial.NewInitial(conf)

	NewServer, err := servers.NewServer(init, conf)
	if err != nil {
		log.Error(err)
	}

	hand := handler.NewHandler(NewServer)
	api := routers.NewapiRegistry(&hand, app, mw)
	api.UsersAPi()
	newCron := cron.New()

	crons := routers.NewCronRegistry(&hand, *NewServer, newCron)
	_ = crons

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + conf.Conf.Port); err != nil {
			log.Fatal(err)
		}
	}()

	<-c
	fmt.Println("Shutting down server...")

	if newCron != nil {
		if err := newCron.Stop(); err != nil {
			log.Errorf("cron job stop error: %v", err)
		}
	}

	if NewServer.WebseocketConn != nil {
		if err := NewServer.WebseocketConn.Close(); err != nil {
			log.Errorf("websocket close error: %v", err)
		}
	}

	if NewServer.DB != nil {
		if err := NewServer.DB.Close(); err != nil {
			log.Errorf("database close error: %v", err)
		}
	}

	if err := app.Shutdown(); err != nil {
		log.Errorf("server shutdown error: %v", err)
	}

}
