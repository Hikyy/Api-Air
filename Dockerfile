FROM golang:1.19
ENV CGO_ENABLED=0
WORKDIR /go_app
COPY --chown=golang:golang . .

RUN apt update && apt install nodejs npm -y
RUN npm i -g nodemon
#RUN nodemon --watch './**/*.go' --watch './**/**/*.go' --watch './templates/**/*.gohtml' --signal SIGTERM --exec 'go' run /go_app/app/main.go

#CMD go run /go_app/app/main.go