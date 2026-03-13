# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM cgr.dev/chainguard/go:latest AS builder

WORKDIR /app

ARG TARGETOS
ARG TARGETARCH

COPY go.mod go.sum* ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags="-s -w" -o /app/sentry-pod main.go

FROM cgr.dev/chainguard/static:latest

WORKDIR /app
COPY --from=builder /app/sentry-pod /app/sentry-pod

USER 65532:65532
EXPOSE 8080

ENTRYPOINT ["/app/sentry-pod"]