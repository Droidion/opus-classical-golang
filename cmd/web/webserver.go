package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/jet/v2"
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
	//f.Use(cors.New())

	// Serve static assets
	static := f.Group("/static", addCache)
	static.Static("/", "./static")

	// Serve dynamic page
	f.Get("/404", app.handle404)
	f.Get("/error", app.handleError)
	f.Get("/about", app.handleAbout)
	f.Get("/composer/:slug", app.handleComposer)
	f.Get("/composer/:composer/work/:work", app.handleWork)
	f.Get("/api/search", app.handleSearch)
	f.Get("/", app.handlePeriods)

	app.webserver = &webserver{
		fiber: f,
	}
}

// startServer starts web server on a given port.
func (webserver *webserver) startServer(port int) error {
	err := webserver.fiber.Listen(":" + strconv.Itoa(port))
	if err != nil {
		return eris.Wrapf(err, "Failed to start web server on port %d", port)
	}
	return nil
}
