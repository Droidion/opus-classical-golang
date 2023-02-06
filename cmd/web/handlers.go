package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
	"strings"
)

// handle404 renders 404 page.
func (app *application) handle404(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{"Title": "404", "Shared": app.sharedTemplateData})
}

// handleError renders error page.
func (app *application) handleError(c *fiber.Ctx) error {
	return c.Render("error", fiber.Map{"Title": "Error", "Shared": app.sharedTemplateData})
}

// handleAbout renders About page.
func (app *application) handleAbout(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{"Title": "About", "Shared": app.sharedTemplateData})
}

// handlePeriods renders Periods (index) page.
func (app *application) handlePeriods(c *fiber.Ctx) error {
	periods, err := app.repo.GetPeriods()
	if err != nil {
		return eris.Wrap(err, "app.repo.GetPeriods")
	}
	for i := range periods {
		periods[i].EnrichForTemplate()
		for j := range periods[i].Composers {
			periods[i].Composers[j].EnrichForTemplate()
		}
	}
	return c.Render("periods", fiber.Map{
		"Shared":  app.sharedTemplateData,
		"Periods": periods,
		"Title":   "Composers",
	})
}

// handleComposer renders Compose page (all works of a given composer).
func (app *application) handleComposer(c *fiber.Ctx) error {
	slug := c.Params("slug")
	composer, err := app.repo.GetComposer(slug)
	if err != nil {
		return eris.Wrap(err, "app.repo.GetComposer")
	}
	composer.EnrichForTemplate()
	genres, err := app.repo.GetGenres(composer.Id)
	if err != nil {
		return eris.Wrap(err, "app.repo.GetGenres")
	}
	for i := range genres {
		for j := range genres[i].Works {
			genres[i].Works[j].EnrichForTemplate()
		}
	}
	return c.Render("composer", fiber.Map{
		"Shared":   app.sharedTemplateData,
		"Title":    composer.LastName,
		"Composer": composer,
		"Genres":   genres,
	})
}

// handleWork renders Work page (all recordings of a given work).
func (app *application) handleWork(c *fiber.Ctx) error {
	composerSlug := c.Params("composer")
	workId, err := c.ParamsInt("work")
	if err != nil {
		return eris.Wrap(err, "c.ParamsInt")
	}

	work, err := app.repo.GetWork(workId)
	if err != nil {
		return eris.Wrap(err, "app.repo.GetWork")
	}
	work.EnrichForTemplate()

	composer, err := app.repo.GetComposer(composerSlug)
	if err != nil {
		return eris.Wrap(err, "app.repo.GetComposer")
	}
	composer.EnrichForTemplate()

	childWorks, err := app.repo.GetChildWorks(work.Id)
	if err != nil {
		return eris.Wrap(err, "app.repo.GetChildWorks")
	}
	for i := range childWorks {
		childWorks[i].EnrichForTemplate()
	}

	recordings, err := app.repo.GetRecordings(work.Id)
	if err != nil {
		return eris.Wrap(err, "app.repo.GetRecordings")
	}
	for i := range recordings {
		recordings[i].EnrichForTemplate()
	}

	return c.Render("work", fiber.Map{
		"Shared":          app.sharedTemplateData,
		"StaticAssetsUrl": app.config.CoversUri,
		"Title":           strings.NewReplacer("&nbsp;<em>", " ", "</em>", "").Replace(work.FullName),
		"Composer":        composer,
		"Work":            work,
		"ChildWorks":      childWorks,
		"Recordings":      recordings,
	})
}

// handleSearch returns composer search result as JSON.
func (app *application) handleSearch(c *fiber.Ctx) error {
	query := c.Query("q")
	composers, err := app.repo.SearchComposers(query, 5)
	if err != nil {
		return eris.Wrap(err, "app.repo.SearchComposers")
	}
	return c.JSON(composers)
}
