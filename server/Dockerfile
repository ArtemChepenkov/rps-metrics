FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/server_run /app/server.go

CMD ["/app/server_run"]
