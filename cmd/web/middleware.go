package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

// addCache middleware adds long caching response headers.
func addCache(c *fiber.Ctx) error {
	c.Set("Cache-Control", "max-age=31536000, immutable")
	return c.Next()
}

// denyCache middleware adds caching response header with no cache directive.
func denyCache(c *fiber.Ctx) error {
	c.Set("Cache-Control", "private, max-age=0")
	return c.Next()
}

// addSecurity middleware adds security headers.
func (app *application) addSecurity(c *fiber.Ctx) error {
	c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	c.Set("X-Permitted-Cross-Domain-Policies", "none")
	c.Set("Referrer-Policy", "no-referrer")
	c.Set("Cross-Origin-Opener-Policy", "same-origin")
	c.Set("Content-Security-Policy", app.config.Csp)
	c.Set("Permissions-Policy", "microphone=(), camera=()")
	return c.Next()
}

func (app *application) errorInterceptor(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code == 404 {
		app.logger.InfoError("errorHandler caught 404, doing redirect", err)
		if c.Path() == "/" {
			return c.Redirect("/", fiber.StatusSeeOther)
		}
		return c.Redirect("/404", fiber.StatusSeeOther)
	}

	app.logger.Error("errorInterceptor caught error", err)
	sentry.CaptureException(err)
	return c.Redirect("/error", fiber.StatusSeeOther)
}
