# Stage 1: Build the binary using the 'Developer' image
FROM cgr.dev/chainguard/go:latest AS builder
WORKDIR /app
COPY go.mod ./
# If you add dependencies later, you'd run 'go mod download' here
COPY main.go .
RUN go build -o sentry-pod main.go

# Stage 2: The hardened Production image
# This image contains NO shell, NO package manager, and NO root user.
FROM cgr.dev/chainguard/static:latest
WORKDIR /app
COPY --from=builder /app/sentry-pod .

# Run as the built-in 'nonroot' user (UID 65532)
USER 65532
ENTRYPOINT ["./sentry-pod"]