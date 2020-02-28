#!/bin/sh

if [ -d "/home/nats/configuration/authorization" ]
then
    echo "init-script: previous configuration present"
else
    cp -r /home/nats/data/configuration/* /home/nats/configuration/
    echo "init-script: starting and checking for configuration"
    echo "init-script: new configuration copied"
    echo "init-script: giving permissions"
    chmod -R 777 /home/nats/
fi

setsid ./nats-api-server &
./nats-server -c ./configuration/nats-server.conf