NAME= chromecast

GITSHA:=$(shell git rev-parse HEAD)
DATE:=$(shell date -u +%FT%TZ)
VERSION:=HEAD

all: build

build:
	@mkdir -p bin/
	go build -o bin/${NAME} -ldflags "-X github.com/teintuc/chromecast-cli/commands.Buildstamp=$(DATE) -X github.com/teintuc/chromecast-cli/commands.Githash=$(GITSHA) -X github.com/teintuc/chromecast-cli/commands.Version=$(VERSION)" .

fmt:
	go fmt github.com/teintuc/chromecast-cli/...

fclean:
	go clean
	rm -fr bin/${NAME}

.PHONY: fmt