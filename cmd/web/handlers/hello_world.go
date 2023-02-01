package handlers

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleHelloWorld(c *fiber.Ctx) error {
	repo, err := utils.GetRepo(c)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Oops1")
	}
	periods, err := repo.GetPeriods()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Oops2")
	}
	return c.Render("index", fiber.Map{
		"Title":   "Hello, World2!",
		"Periods": periods,
	}, "layouts/main")
}
