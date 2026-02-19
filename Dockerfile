FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/hospital

FROM alpine:3.23
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/app .
COPY --from=builder /build/migrations ./migrations

EXPOSE 8000
CMD ["./app"]