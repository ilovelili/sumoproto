FROM golang:1.10.0-alpine
LABEL maintainer="min<min_ju@invastsec.jp>"
# add git
RUN apk update && apk add git

# Installing Go dependencies
RUN go get github.com/Masterminds/glide

ENV SRC_DIR=/go/src/github.com/ilovelili/sumoproto
WORKDIR $SRC_DIR

# Copying files
COPY . $SRC_DIR

# Running glide
RUN glide up