FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . ./
RUN go mod tidy
RUN go build -o /weather-bot cmd/main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /weather-bot /weather-bot
COPY --from=builder app/.env .env
COPY --from=builder app/config.json config.json
COPY --from=builder app/internal/database/postgres/migrations/ internal/database/postgres/migrations/
CMD ["/weather-bot"]