FROM golang:1.17.4-alpine3.15

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 5000

CMD ["./main"]