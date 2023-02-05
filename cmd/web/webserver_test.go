package main

import (
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
