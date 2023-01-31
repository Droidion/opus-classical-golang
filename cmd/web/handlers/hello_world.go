package handlers

import "github.com/gofiber/fiber/v2"

func HandleHelloWorld(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World2!",
	}, "layouts/main")
}
