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
		return eris.Wrap(err, "Failed to get periods from repo")
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
		return eris.Wrapf(err, "Failed to get composer from repo by slug %s", slug)
	}
	composer.EnrichForTemplate()
	genres, err := app.repo.GetGenres(composer.Id)
	if err != nil {
		return eris.Wrapf(err, "Failed to get genres from repo by composer id %d", composer.Id)
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
		return eris.Wrapf(err, "Failed to parse work id to number from string %s", c.Params("work"))
	}

	work, err := app.repo.GetWork(workId)
	if err != nil {
		return eris.Wrapf(err, "Failed to get work from repo by work id %d", workId)
	}
	work.EnrichForTemplate()

	composer, err := app.repo.GetComposer(composerSlug)
	if err != nil {
		return eris.Wrapf(err, "Failed to get composer from repo by composer slug %s", composerSlug)
	}
	composer.EnrichForTemplate()

	childWorks, err := app.repo.GetChildWorks(work.Id)
	if err != nil {
		return eris.Wrapf(err, "Failed to get children works from repo by work id %d", work.Id)
	}
	for i := range childWorks {
		childWorks[i].EnrichForTemplate()
	}

	recordings, err := app.repo.GetRecordings(work.Id)
	if err != nil {
		return eris.Wrapf(err, "Failed to get recordings from repo by work id %d", work.Id)
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
	limit := 5
	composers, err := app.repo.SearchComposers(query, limit)
	if err != nil {
		return eris.Wrapf(err, "Failed to get composer search results from repo with query %s and result count %d", query, limit)
	}
	return c.JSON(composers)
}
