FROM golang:alpine

ADD . /go/src/github.com/dunglas/calavera

RUN apk add -U git && \
  go get github.com/dunglas/calavera && \
  go install github.com/dunglas/calavera && \
  rm -rf /go/pkg && \
  rm -rf /go/src && \
  rm -rf /go/cache/apk/*

ENTRYPOINT ["/go/bin/calavera"]
