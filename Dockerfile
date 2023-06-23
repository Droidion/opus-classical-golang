# Build app
FROM golang:1.20.5-buster as go-builder
WORKDIR /app
# Install packages
COPY go.* ./
RUN go mod download
# Build app
COPY . ./
RUN go build -v -o server github.com/droidion/opus-classical-golang/cmd/api

# Prepare runtime
FROM debian:buster-slim as runtime
WORKDIR /app
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy app
COPY --from=go-builder /app/server /app/server

# Run app
CMD ["/app/server"]