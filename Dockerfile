FROM golang:1.20 as BUILDER

ENV CGO_ENABLED=0

WORKDIR /go_app

COPY . .

RUN echo " !! À chaque modifications pour run le programme : go run main.go !!"

CMD go run ./app