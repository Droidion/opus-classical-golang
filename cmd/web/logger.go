package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"os"
	"time"
)

// Logger contains internal implementation of logger and public methods for logging.
type Logger struct {
	zerolog zerolog.Logger
}

// LogsEvents exposes methods for logging events across the app.
type LogsEvents interface {
	Info(msg string)
	InfoError(msg string, err error)
	Error(msg string, err error)
	Fatal(msg string, err error)
}

// initSentry initialises sentry.io.
func initSentry(dsn string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return eris.Wrap(err, "Could not initialise sentry")
	}
	defer sentry.Flush(2 * time.Second)
	return nil
}

// newLogger returns new instance of logger.
func newLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &Logger{
		zerolog: zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger(),
	}
}

// Info handles event of info level, outputting simple text.
func (logger *Logger) Info(msg string) {
	logger.zerolog.Info().Msg(msg)
}

// InfoError handles event of info level, outputting text description and Go error.
// Use it for handling Golang errors which are not really errors.
func (logger *Logger) InfoError(msg string, err error) {
	logger.zerolog.Info().Err(err).Msg(msg)
}

// Error handles event of error level, outputting text and Go error and submitting to sentry.
func (logger *Logger) Error(msg string, err error) {
	sentry.CaptureException(eris.Wrap(err, msg))
	logger.zerolog.Error().Stack().Err(err).Msg(msg)
}

// Fatal handles event of fatal level, outputting text and Go error and submitting to sentry.
func (logger *Logger) Fatal(msg string, err error) {
	sentry.CaptureException(eris.Wrap(err, msg))
	logger.zerolog.Fatal().Err(err).Msg(msg)
}
