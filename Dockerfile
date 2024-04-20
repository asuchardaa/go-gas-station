FROM golang:1.22
LABEL authors="Adam"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
