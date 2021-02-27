FROM golang:1.15.2-alpine

RUN apk update && apk add --update-cache git ca-certificates tzdata make gcc g++ && git config --global http.sslVerify false && go get golang.org/x/tools/cmd/goimports && go get -u golang.org/x/lint/golint && go get golang.org/x/tools/gopls && go get -u github.com/kisielk/errcheck && go get -u github.com/golang/mock/mockgen && go get -u github.com/securego/gosec/cmd/gosec && go get -u github.com/google/wire/cmd/wire && go get -u github.com/fzipp/gocyclo

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app