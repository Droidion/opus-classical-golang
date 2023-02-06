package main

import (
	"github.com/dchest/uniuri"
	"github.com/droidion/opus-classical-golang/internal/models"
)

// application contains all application-level data.
type application struct {
	config             *config
	logger             LogsEvents
	repo               models.ProvidesData
	webserver          servesRoutes
	sharedTemplateData struct {
		appRunId       string
		umamiUri       string
		umamiWebsiteId string
	}
	appRunId string
}

func main() {
	app := &application{}
	app.initConfig()
	_ = initSentry(app.config.SentryDsn)
	app.sharedTemplateData.appRunId = uniuri.New()
	app.sharedTemplateData.umamiUri = app.config.UmamiUri
	app.sharedTemplateData.umamiWebsiteId = app.config.UmamiWebsiteId
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
