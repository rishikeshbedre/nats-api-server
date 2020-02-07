FROM debian:stretch-slim as build

ENV GO_VERSION=1.12.6

RUN apt-get update \
	&& apt-get install -y --no-install-recommends build-essential wget curl ca-certificates git \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/*

RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
	&& tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
	&& rm go${GO_VERSION}.linux-amd64.tar.gz 

COPY apis nats-api-server/apis
COPY configuration nats-api-server/configuration
COPY lib nats-api-server/lib
COPY util nats-api-server/util
COPY scripts nats-api-server/scripts
COPY go.mod go.sum main.go Makefile nats-api-server/

WORKDIR /nats-api-server

RUN export PATH=$PATH:/usr/local/go/bin \
	&& ls \
	&& make 

FROM nats:2.1.2-alpine3.10 as middlelayer

FROM alpine:3.10

COPY --from=build /nats-api-server/nats-api-server /home/nats/
COPY --from=middlelayer /usr/local/bin/nats-server /home/nats/nats-server
COPY configuration /home/nats/configuration
COPY scripts /home/nats/scripts

EXPOSE 4222 8222 6222 6060

WORKDIR /home/nats

CMD setsid /home/nats/nats-api-server & /home/nats/nats-server -c /home/nats/configuration/nats-server.conf
