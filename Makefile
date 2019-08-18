VERSION=V0.1.0
ifndef GOPATH
	$(error GOPATH is not set)
endif
GOAPTH = $(firstword $(subst :, ,${GOPATH}))

build:
	cd cmd/envparser && go build -ldflags "-s -w \
	-Xgithub.com/gokultp/envparser/internal/version.Version=${VERSION} \
	-Xgithub.com/gokultp/envparser/internal/version.MinVersion=`git rev-parse HEAD` \
	-Xgithub.com/gokultp/envparser/internal/version.BuildTime=`date +%FT%T%z` " \
	-o ${GOPATH}/bin/envparser

.PHONY:
	build
