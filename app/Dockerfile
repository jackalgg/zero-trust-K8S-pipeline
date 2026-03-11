FROM cgr.dev/chainguard/go:latest AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/sentry-pod main.go

FROM cgr.dev/chainguard/static:latest
WORKDIR /app
COPY --from=builder /app/sentry-pod /app/sentry-pod
USER 65532
ENTRYPOINT ["/app/sentry-pod"]