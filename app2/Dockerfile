FROM golang:1.20
WORKDIR /app2
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /app2
EXPOSE 8888
CMD [ "/app2","--config=config/config.toml"]