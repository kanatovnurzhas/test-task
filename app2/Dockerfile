FROM golang:1.20-alpine
WORKDIR /app2
COPY . .
RUN go build -o app2 cmd/main.go
CMD ["./app2","--config=config/config.toml"]