package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
	"strings"
)

func (app *application) handle404(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{"Title": "404", "Shared": app.sharedTemplateData})
}

func (app *application) handleError(c *fiber.Ctx) error {
	return c.Render("error", fiber.Map{"Title": "Error", "Shared": app.sharedTemplateData})
}

func (app *application) handleAbout(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{"Title": "About", "Shared": app.sharedTemplateData})
}

func (app *application) handlePeriods(c *fiber.Ctx) error {
	periods, err := app.repo.GetPeriods()
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
		"Shared":  app.sharedTemplateData,
		"Periods": periods,
		"Title":   "Composers",
	})
}

func (app *application) handleComposer(c *fiber.Ctx) error {
	slug := c.Params("slug")
	composer, err := app.repo.GetComposer(slug)
	if err != nil {
		return eris.Wrap(err, "Could not get composer from endpoint handler")
	}
	composer.Process()
	genres, err := app.repo.GetGenres(composer.Id)
	if err != nil {
		return eris.Wrap(err, "Could not get genres from endpoint handler")
	}
	for i := range genres {
		for j := range genres[i].Works {
			genres[i].Works[j].Process()
		}
	}
	return c.Render("composer", fiber.Map{
		"Shared":   app.sharedTemplateData,
		"Title":    composer.LastName,
		"Composer": composer,
		"Genres":   genres,
	})
}

func (app *application) handleWork(c *fiber.Ctx) error {
	composerSlug := c.Params("composer")
	workId, err := c.ParamsInt("work")
	if err != nil {
		return eris.Wrap(err, "Could not parse work id")
	}

	work, err := app.repo.GetWork(workId)
	if err != nil {
		return eris.Wrap(err, "Could not get work from endpoint handler")
	}
	work.Process()

	composer, err := app.repo.GetComposer(composerSlug)
	if err != nil {
		return eris.Wrap(err, "Could not get composer from endpoint handler")
	}
	composer.Process()

	childWorks, err := app.repo.GetChildWork(work.Id)
	if err != nil {
		return eris.Wrap(err, "Could not get child works from endpoint handler")
	}
	for i := range childWorks {
		childWorks[i].Process()
	}

	recordings, err := app.repo.GetRecordings(work.Id)
	if err != nil {
		return eris.Wrap(err, "Could not get recordings from endpoint handler")
	}
	for i := range recordings {
		recordings[i].Process()
	}

	r := strings.NewReplacer("&nbsp;<em>", " ", "</em>", "")

	return c.Render("work", fiber.Map{
		"Shared":          app.sharedTemplateData,
		"StaticAssetsUrl": app.config.CoversUri,
		"Title":           r.Replace(work.FullName),
		"Composer":        composer,
		"Work":            work,
		"ChildWorks":      childWorks,
		"Recordings":      recordings,
	})
}

func (app *application) handleSearch(c *fiber.Ctx) error {
	query := c.Query("q")
	composers, err := app.repo.SearchComposers(query, 5)
	if err != nil {
		return eris.Wrap(err, "Could not get composer from endpoint handler")
	}
	return c.JSON(composers)
}
