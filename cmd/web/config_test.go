package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	app := &application{}

	app.initConfig()

	assert.NotNil(t, app.config)
}
