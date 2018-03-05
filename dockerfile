FROM golang:1.10.0-alpine
LABEL maintainer="min<min_ju@invastsec.jp>"
# add git
RUN apk update && apk add git

RUN echo "Installing Go dependencies ... "
RUN go get github.com/Masterminds/glide
RUN go get github.com/golang/protobuf/{proto,protoc-gen-go}

RUN echo "Copying files ... "
COPY . /go/src/github.com/ilovelili/sumoproto

WORKDIR /go/src/github.com/ilovelili/sumoproto

RUN echo "Running glide ... "
RUN glide up