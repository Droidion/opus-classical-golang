#!/usr/bin/env bash
# exit on error
set -o errexit

# Build and run app
pnpm i
pnpm run sass
pnpm run build
go mod download
go build -v -o server ./cmd/web