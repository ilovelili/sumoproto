FROM golang:1.10.0-alpine
LABEL maintainer="min<min_ju@invastsec.jp>"
# add git
RUN apk update && apk add git

RUN go get github.com/Masterminds/glide

COPY . /go/src/github.com/ilovelili/sumoproto

WORKDIR /go/src/github.com/ilovelili/sumoproto

RUN glide up