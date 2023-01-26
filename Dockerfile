FROM golang

RUN apt-get update

WORKDIR /usr/src/app

COPY . .

RUN go get
RUN go build
