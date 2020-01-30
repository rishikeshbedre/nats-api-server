#!/bin/sh

#rm -rf *.tar

docker build --build-arg http_proxy="$http_proxy" --build-arg https_proxy="$https_proxy" -t nats-api-server:0.0.1 .

docker run -it -p 4222:4222 -p 6060:6060 nats-api-server:0.0.1

#docker save -o nats-api-server.tar nats-api-server:0.0.1

#chmod 777 *.tar
