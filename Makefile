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

vendor:
	dep ensure

update:
	dep ensure -update

fclean:
	go clean
	rm -fr bin/${NAME}

test:
	go test github.com/teintuc/chromecast-cli/...

.PHONY: fmt vendor update test
