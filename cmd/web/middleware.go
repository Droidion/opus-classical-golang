package main

import (
	"github.com/gofiber/fiber/v2"
)

// addCache middleware adds long caching response headers.
func addCache(c *fiber.Ctx) error {
	c.Set(fiber.HeaderCacheControl, "max-age=31536000, immutable")
	return c.Next()
}

// denyCache middleware adds caching response header with no cache directive.
func denyCache(c *fiber.Ctx) error {
	c.Set(fiber.HeaderCacheControl, "private, max-age=0")
	return c.Next()
}

// addSecurity middleware adds security headers.
func (app *application) addSecurity(c *fiber.Ctx) error {
	c.Set(fiber.HeaderStrictTransportSecurity, "max-age=31536000; includeSubDomains")
	c.Set(fiber.HeaderXPermittedCrossDomainPolicies, "none")
	c.Set(fiber.HeaderReferrerPolicy, "no-referrer")
	c.Set("Cross-Origin-Opener-Policy", "same-origin")
	c.Set(fiber.HeaderContentSecurityPolicy, app.config.Csp)
	c.Set(fiber.HeaderPermissionsPolicy, "microphone=(), camera=()")
	return c.Next()
}

// errorInterceptor catches all errors from handlers
func (app *application) errorInterceptor(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code == 404 {
		app.logger.InfoError("errorInterceptor 404", err)
		if c.Path() == "/" {
			return c.Redirect("/", fiber.StatusSeeOther)
		}
		return c.Redirect("/404", fiber.StatusSeeOther)
	}

	app.logger.Error("errorInterceptor non 404", err)
	return c.Redirect("/error", fiber.StatusSeeOther)
}
