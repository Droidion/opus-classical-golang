package main

import (
	"github.com/droidion/opus-classical-golang/internal/models"
)

type application struct {
	config    *config
	logger    LogsEvents
	repo      models.ProvidesData
	webserver servesRoutes
}

func main() {
	app := &application{}
	app.initConfig()
	app.logger = newLogger()
	db := app.initDb()
	defer db.Close()
	app.repo = &models.Repo{Db: db}
	app.createWebserver()
	err := app.webserver.startServer(app.config.ServerPort)
	if err != nil {
		app.logger.Fatal("Could not start web server", err)
	}
}
