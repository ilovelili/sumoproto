#!/bin/sh
PKG_OK=$(command -v glide)
if [ "" = "$PKG_OK" ]; then
    echo "No glide. Setting up..."
    curl https://glide.sh/get | sh
fi

# this is necessary in bin
go get github.com/golang/protobuf/{proto,protoc-gen-go}

glide install

sudo docker build -t sumoproto:0.1 .

sudo docker-compose build && sudo docker-compose up
