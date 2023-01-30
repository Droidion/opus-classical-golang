package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntiSentryReturnsError(t *testing.T) {
	dns := "asd"
	err := initSentry(dns)
	assert.NotNil(t, err)
}

func TestIntiSentryDoesNotReturnError(t *testing.T) {
	dns := "https://secret@secret.ingest.sentry.io/secret"
	err := initSentry(dns)
	assert.Nil(t, err)
}
