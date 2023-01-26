FROM golang

RUN apt-get update

WORKDIR /usr/src/app

COPY . .

EXPOSE 8000

RUN go get
RUN go build
