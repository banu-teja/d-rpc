FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/drpcd ./cmd/drpcd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/drpcd /app/drpcd

EXPOSE 8080

# Set environment variables with default values
ENV RPC_URL=http://localhost:8545 \
    PORT=8080

# Run the binary
CMD ["/app/drpcd"]
