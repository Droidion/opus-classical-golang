# Opus Classic Golang version

# How to run

Have the latest version of Golang installed.

Create `app.env` with appropriate environment variables:

```dotenv
SERVER_PORT=8080
SENTRY_DSN=
DATABASE_CONNECTION_STRING=postgres://localhost:5432/opusclassical
```

Run app with `$ go run ./...`

# How to test

Run `$ go test ./... -cover`