FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/client_run /app/client.go

ENTRYPOINT ["./client_run"]

CMD ["-rps", $RPS, "-duration", $DURATION]
