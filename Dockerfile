FROM golang:1.20 as BUILDER

ENV CGO_ENABLED=0

WORKDIR /go_app

COPY . .

CMD go run ./app