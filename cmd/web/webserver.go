package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

// createWebserver creates new web server.
func (app *application) createWebserver() {
	engine := jet.New("./views", ".jet")

	f := fiber.New(fiber.Config{
		Views:        engine,
		ViewsLayout:  "layouts/main",
		ErrorHandler: app.errorInterceptor,
	})

	// Add middleware
	f.Use(helmet.New())
	f.Use(denyCache)
	f.Use(app.addSecurity)
	f.Use(recover.New())
	f.Use(cors.New())

	// Serve static assets
	static := f.Group("/static", addCache)
	static.Static("/", "./static")

	// Serve dynamic page
	f.Get("/404", app.Handle404)
	f.Get("/error", app.HandleError)
	f.Get("/about", app.HandleAbout)
	f.Get("/composer/:slug", app.HandleComposer)
	f.Get("/composer/:composer/work/:work", app.HandleWork)
	f.Get("/api/search", app.HandleSearch)
	f.Get("/", app.HandlePeriods)

	app.webserver = &webserver{
		fiber: f,
	}
}

// startServer starts web server on a given port.
func (webserver *webserver) startServer(port int) error {
	err := webserver.fiber.Listen(":" + strconv.Itoa(port))
	if err != nil {
		return eris.Wrap(err, "Could not start web server")
	}
	return nil
}
