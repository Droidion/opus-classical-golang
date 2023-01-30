package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWebserver(t *testing.T) {
	app := &application{
		config: &config{},
	}
	app.createWebserver()
	assert.NotNil(t, app.webserver)
}

func TestAddRoutes(t *testing.T) {
	webserver := &webserver{
		fiber: fiber.New(),
	}

	webserver.registerRoutes()
	routes := webserver.fiber.GetRoutes()

	assert.NotZero(t, len(routes))
}
