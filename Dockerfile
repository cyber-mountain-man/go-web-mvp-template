# ----------------------
# Build Stage
# ----------------------
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/api

# ----------------------
# Final Stage
# ----------------------
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/static ./static

USER nonroot:nonroot

CMD ["/app/app"]
