# Opus Classical

Curated catalogue of western classical music recordings. This version uses Go and Svelte.

[![Build → Test](https://github.com/Droidion/opus-classical-golang/actions/workflows/build.yml/badge.svg)](https://github.com/Droidion/opus-classical-golang/actions/workflows/build.yml)

![screenshot](screenshot.webp)

# How to run

Install:

- [The latest version of Go](https://go.dev/dl/).
- [LTS version of Node.js](https://nodejs.org/en/download/).
- [The latest version of pnpm](https://pnpm.io/installation).

Create `app.env` with appropriate environment variables:

```dotenv
# What port webserver runs on.
SERVER_PORT=8080
# Sentry.io DSN for submitting logs.
SENTRY_DSN=
# Postgres connection string.
DATABASE_CONNECTION_STRING=postgres://localhost:5432/opusclassical
# URI where covers reside.
COVERS_URI=https://opusclassical.zunh.nl-ams1.upcloudobjects.com/
# Content Security Policy HTTP header.
CSP="default-src 'none'; manifest-src 'self'; connect-src 'self' https://analytics.umami.is; script-src 'self' https://analytics.umami.is; style-src 'self'; img-src 'self' https://opusclassical.zunh.nl-ams1.upcloudobjects.com"
# URI for submitting client-side statustics with Umami.
UMAMI_URI=
# Website ID for submitting client-side statustics with Umami.
UMAMI_WEBSITE_ID=
```

Install packages:

- `$ pnpm i` for client-side.
- `$ go mod download`. for server-side.

Compile static assets:

- `$ pnpm run sass`
- `$ pnpm run build`

Build and run Go server

- `$ go run ./cmd/web`

Alternatively, build everything with `$ ./build.sh` and run server with `$ ./server`.

Alternatively, have Docker installed and run with `$ docker compose up`.

# Run unit tests

- `$ go test ./... -cover`

# Deploy

Opus Classical uses [render.com](https://render.com) for hosting. `main` branch automatically deploys to https://opusclassical.net