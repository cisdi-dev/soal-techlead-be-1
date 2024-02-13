FROM golang:1.21.6-bullseye AS builder

ENV GO111MODULE="auto"

RUN apt-get -y install git

WORKDIR /opt

COPY . .

RUN go build .

ENTRYPOINT ["/opt/soal-techlead-be-1"]