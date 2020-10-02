FROM golang:1.10-alpine

RUN apk add --update tzdata bash wget curl git;

RUN mkdir -p $$GOPATH/bin && \
    go get -u github.com/golang/dep/cmd/dep

# directory
ADD . /go/src/rest_echo
WORKDIR /go/src/rest_echo

CMD dep ensure && go build main.go && ./main

EXPOSE 8000