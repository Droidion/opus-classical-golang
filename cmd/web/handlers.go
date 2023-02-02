package main

import (
	"github.com/droidion/opus-classical-golang/internal/models"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

func getRepo(c *fiber.Ctx) (*models.Repo, error) {
	return utils.GetLocal[*models.Repo](c, "repo")
}

func Handle404(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{"Title": "404"})
}

func HandleError(c *fiber.Ctx) error {
	return c.Render("error", fiber.Map{"Title": "Error"})
}

func HandleAbout(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{"Title": "About"})
}

func HandlePeriods(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return eris.Wrap(err, "Could not get repo from endpoint handler")
	}
	periods, err := repo.GetPeriods()
	if err != nil {
		return eris.Wrap(err, "Could not get periods from endpoint handler")
	}
	for i := range periods {
		periods[i].Process()
		for j := range periods[i].Composers {
			periods[i].Composers[j].Process()
		}
	}
	return c.Render("periods", fiber.Map{
		"Periods": periods,
		"Title":   "Composers",
	})
}

func HandleComposer(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return eris.Wrap(err, "Could not get repo from endpoint handler")
	}
	slug := c.Params("slug")
	composer, err := repo.GetComposer(slug)
	if err != nil {
		return eris.Wrap(err, "Could not get composer from endpoint handler")
	}
	composer.Process()
	genres, err := repo.GetGenres(composer.Id)
	if err != nil {
		return eris.Wrap(err, "Could not get genres from endpoint handler")
	}
	for i := range genres {
		for j := range genres[i].Works {
			genres[i].Works[j].Process()
		}
	}
	return c.Render("composer", fiber.Map{
		"Title":    composer.LastName,
		"Composer": composer,
		"Genres":   genres,
	})
}
