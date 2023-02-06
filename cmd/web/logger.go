package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
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
	FatalWithContext(msg string, err error, key string, value string)
}

// initSentry initialises sentry.io.
func initSentry(dsn string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return eris.Wrapf(err, "Failed to initialise sentry with DSN %s", dsn)
	}
	defer sentry.Flush(2 * time.Second)
	return nil
}

// newLogger returns new instance of logger.
func (app *application) newLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	var writer io.Writer
	if app.config.IsDev {
		writer = zerolog.ConsoleWriter{Out: os.Stderr}
	} else {
		writer = os.Stderr
	}
	app.logger = &Logger{
		zerolog: zerolog.New(writer).With().Timestamp().Caller().Logger(),
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
	logger.zerolog.Fatal().Stack().Err(err).Msg(msg)
}

func (logger *Logger) FatalWithContext(msg string, err error, key string, value string) {
	sentry.CaptureException(eris.Wrap(err, msg))
	logger.zerolog.Fatal().Stack().Str(key, value).Err(err).Msg(msg)
}
