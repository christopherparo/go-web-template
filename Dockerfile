# syntax=docker/dockerfile:1

# ── Build stage ──────────────────────────────────────────────────────────────
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /build

# Cache dependencies separately from source.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=0   → static binary, no libc dependency
# -trimpath       → reproducible builds, no local paths in binary
# -ldflags "-s -w" → strip debug info + DWARF symbols
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /build/server \
    ./cmd/server

# ── Runtime stage ────────────────────────────────────────────────────────────
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -H -s /sbin/nologin appuser

COPY --from=builder /build/server /usr/local/bin/server

USER appuser

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["server"]
