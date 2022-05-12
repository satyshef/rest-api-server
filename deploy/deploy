#!/bin/bash

IMAGE_NAME=$1
APP_NAME="api-server"

if [ -z $IMAGE_NAME ]; then
    echo "Image Name not set"
    exit 1
fi

if [ ! -x "$(command -v docker)" ]; then
    echo "Docker not installed"
    exit 1
fi

echo "Remove old version"
docker rm $APP_NAME -f

echo "Deploy application"
docker run -d -p 443:7000 -p 80:7000 --restart always --name $APP_NAME $IMAGE_NAME


