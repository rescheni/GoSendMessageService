#!/bin/sh
version=1.11
docker build -t send_message_service-go:$version .

docker stop send_message_service-go

docker rm send_message_service-go

docker run -d \
    -p 38080:8080 \
    -p 33000:3000 \
    --name send_message_service-go \
    send_message_service-go:$version

docker logs -f send_message_service-go