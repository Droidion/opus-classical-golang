package handlers

import (
	"github.com/droidion/opus-classical-golang/internal/models"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func HandlePeriods(c *fiber.Ctx) error {
	repo, err := utils.GetLocal[*models.Repo](c, "repo")
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Oops1")
	}
	periods, err := repo.GetPeriods()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Oops2")
	}
	for i := range periods {
		periods[i].Process()
		for j := range periods[i].Composers {
			periods[i].Composers[j].Process()
		}
	}
	return c.Render("periods", fiber.Map{
		"Periods": periods,
	}, "layouts/main")
}
