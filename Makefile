VERSION=V0.1.0
ifndef GOPATH
	$(error GOPATH is not set)
endif
GOAPTH = $(firstword $(subst :, ,${GOPATH}))

compile:
	cd cmd/envparser && go build -ldflags "-s -w \
	-X github.com/gokultp/go-envparser/internal/version.Version=${VERSION} \
	-X github.com/gokultp/go-envparser/internal/version.MinVersion=`git rev-parse HEAD` \
	-X github.com/gokultp/go-envparser/internal/version.BuildTime=`date +%FT%T%z` " \
	-o ${GOPATH}/bin/envparser

config:
	go get -d
build:
	config
	compile
.PHONY:
	build
