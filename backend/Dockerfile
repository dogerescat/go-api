FROM golang:1.18-alpine

RUN apk update && apk add git

COPY ./src/ /app/src

WORKDIR /app/src

RUN go mod tidy

RUN go build -o ../main

WORKDIR /app

ENV GO111MODULE=on

EXPOSE 80

CMD ["./main"]
