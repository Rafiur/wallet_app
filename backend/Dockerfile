# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/wallet_app ./main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/seeder ./cmd/seeder

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=build /out/wallet_app ./wallet_app
COPY --from=build /out/seeder ./seeder

EXPOSE 8000
CMD ["./wallet_app"]
