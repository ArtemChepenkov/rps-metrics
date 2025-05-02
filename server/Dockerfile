FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY server.go .

RUN go build -o server server.go

ENTRYPOINT ["./server"]
