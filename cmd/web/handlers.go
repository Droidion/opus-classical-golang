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

func getConfig(c *fiber.Ctx) (*config, error) {
	return utils.GetLocal[*config](c, "config")
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

func HandleWork(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return eris.Wrap(err, "Could not get repo from endpoint handler")
	}
	config, err := getConfig(c)
	if err != nil {
		return eris.Wrap(err, "Could not get config from endpoint handler")
	}

	composerSlug := c.Params("composer")
	workId, err := c.ParamsInt("work")
	if err != nil {
		return eris.Wrap(err, "Could not parse work id")
	}

	work, err := repo.GetWork(workId)
	if err != nil {
		return eris.Wrap(err, "Could not get work from endpoint handler")
	}
	work.Process()

	composer, err := repo.GetComposer(composerSlug)
	if err != nil {
		return eris.Wrap(err, "Could not get composer from endpoint handler")
	}
	composer.Process()

	childWorks, err := repo.GetChildWork(work.Id)
	if err != nil {
		return eris.Wrap(err, "Could not get child works from endpoint handler")
	}
	for i := range childWorks {
		childWorks[i].Process()
	}

	recordings, err := repo.GetRecordings(work.Id)
	if err != nil {
		return eris.Wrap(err, "Could not get recordings from endpoint handler")
	}
	for i := range recordings {
		recordings[i].Process()
	}

	return c.Render("work", fiber.Map{
		"Title":           work.FullName,
		"Composer":        composer,
		"Work":            work,
		"ChildWorks":      childWorks,
		"Recordings":      recordings,
		"StaticAssetsUrl": config.CoversUri,
	})
}

func HandleSearch(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return eris.Wrap(err, "Could not get repo from endpoint handler")
	}
	query := c.Query("q")
	composers, err := repo.SearchComposers(query, 5)
	if err != nil {
		return eris.Wrap(err, "Could not get composer from endpoint handler")
	}
	return c.JSON(composers)
}
