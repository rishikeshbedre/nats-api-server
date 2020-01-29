FROM debian:stretch-slim

ENV GO_VERSION=1.12.6

RUN apt-get update \
	&& apt-get install -y build-essential wget curl ca-certificates git

RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz 

COPY apis nats-api-server/apis
COPY configuration nats-api-server/configuration
COPY lib nats-api-server/lib
COPY util nats-api-server/util
COPY go.mod go.sum main.go Makefile nats-api-server/

RUN export PATH=$PATH:/usr/local/go/bin \
    && cd nats-api-server \
    && ls \
	&& make 

FROM alpine:3.10

COPY --from=0 /nats-api-server/nats-api-server /home/nats/
COPY --from=nats:2.1.2-alpine3.10 /usr/local/bin/nats-server /home/nats/nats-server
COPY configuration /home/nats/configuration

EXPOSE 4222 8222 6222 6060

CMD setsid /home/nats/nats-api-server & /home/nats/nats-server -c /home/nats/configuration/nats-server.conf
