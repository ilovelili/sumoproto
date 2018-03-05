#!/bin/sh
PKG_OK=$(command -v glide)
if [ "" = "$PKG_OK" ]; then
    echo "No glide. Setting up..."
    curl https://glide.sh/get | sh
fi

# this is necessary in bin
go get github.com/golang/protobuf/{proto,protoc-gen-go}

docker build -t sumoproto:0.1 .
docker-compose build && docker-compose up
