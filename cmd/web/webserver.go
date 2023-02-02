package main

import (
	"github.com/droidion/opus-classical-golang/internal/models"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/template/jet"
	"github.com/rotisserie/eris"
	"strconv"
)

// webserver wraps internal implementation of web framework.
type webserver struct {
	fiber *fiber.App
}

type servesRoutes interface {
	startServer(port int) error
}

func (webserver *webserver) registerRoutes() {
	webserver.fiber.Use(helmet.New())
	webserver.fiber.Use(denyCache)
	webserver.fiber.Use(addSecurity)
	webserver.fiber.Use(recover.New())
	webserver.fiber.Use(logger.New())

	static := webserver.fiber.Group("/static", addCache)
	static.Static("/", "./static")

	webserver.fiber.Get("/404", Handle404)
	webserver.fiber.Get("/error", HandleError)
	webserver.fiber.Get("/about", HandleAbout)
	webserver.fiber.Get("/composer/:slug", HandleComposer)
	webserver.fiber.Get("/composer/:composer/work/:work", HandleWork)
	webserver.fiber.Get("/api/search", HandleSearch)
	webserver.fiber.Get("/", HandlePeriods)
}

// injectRepo is a middleware that inject database repository into all web handlers context.
func injectRepo(repo models.ProvidesData) func(c *fiber.Ctx) error {
	foo := func(c *fiber.Ctx) error {
		utils.SetLocal[models.ProvidesData](c, "repo", repo)
		c.Locals("repo", repo)
		return c.Next()
	}

	return foo
}

func injectConfig(cfg *config) func(c *fiber.Ctx) error {
	foo := func(c *fiber.Ctx) error {
		utils.SetLocal[*config](c, "config", cfg)
		c.Locals("config", cfg)
		return c.Next()
	}

	return foo
}

// addMiddleware registers all middlewares for a fiber webserver.
func (webserver *webserver) addMiddleware(repo models.ProvidesData, config *config) {
	webserver.fiber.Use(recover.New())
	webserver.fiber.Use(injectRepo(repo))
	webserver.fiber.Use(injectConfig(config))
	webserver.fiber.Use(cors.New())
}

// createWebserver creates new web server.
func (app *application) createWebserver() {
	engine := jet.New("./views", ".jet")

	srv := &webserver{
		fiber: fiber.New(fiber.Config{
			Views:        engine,
			ViewsLayout:  "layouts/main",
			ErrorHandler: app.errorInterceptor,
		}),
	}
	srv.addMiddleware(app.repo, app.config)
	srv.registerRoutes()
	app.webserver = srv
}

// startServer starts web server on a given port.
func (webserver *webserver) startServer(port int) error {
	err := webserver.fiber.Listen(":" + strconv.Itoa(port))
	if err != nil {
		return eris.Wrap(err, "Could not start web server")
	}
	return nil
}
