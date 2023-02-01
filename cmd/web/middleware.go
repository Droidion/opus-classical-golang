package main

import "github.com/gofiber/fiber/v2"

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
func addSecurity(c *fiber.Ctx) error {
	c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	c.Set("X-Permitted-Cross-Domain-Policies", "none")
	c.Set("Referrer-Policy", "no-referrer")
	c.Set("Cross-Origin-Opener-Policy", "same-origin")
	// c.Set("Content-Security-Policy", app.cfg.Csp)
	c.Set("Permissions-Policy", "microphone=(), camera=()")
	return c.Next()
}
