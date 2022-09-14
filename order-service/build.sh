#!/bin/bash

echo "Building binary..."
env GOOS=linux GOARCH=amd64 go build .
echo "Building docker image..."
docker build -t order-service .
