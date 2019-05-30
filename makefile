OUT := emojipedia
PKG := github.com/gellel/emojipedia
VERSION := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: binary

binary:
	go build -i -v -o ${OUT} -ldflags="-X main.version=${VERSION}"

test:
	go test -short ${PKG_LIST}

clean:
rm ${OUT}