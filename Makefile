.PHONY: all build clean default help init test format check-license
default: help

GO_TEST_PATH ?= ./...
GO_TEST_EXTRA ?=
GO_TEST_COVER_PROFILE ?= cover.out
GO_TEST_CODECOV ?=

BUILD ?= $(shell date +%FT%T%z)
GOVERSION ?= $(shell go version | cut -d " " -f3)
COMPONENT_VERSION ?= $(shell git describe --abbrev=0 --always)
COMPONENT_BRANCH ?= $(shell git describe --always --contains --all)
PMM_RELEASE_FULLCOMMIT ?= $(shell git rev-parse HEAD)
GO_BUILD_LDFLAGS = -X main.version=${COMPONENT_VERSION} -X main.buildDate=${BUILD} -X main.commit=${PMM_RELEASE_FULLCOMMIT} -X main.Branch=${COMPONENT_BRANCH} -X main.GoVersion=${GOVERSION} -s -w
NAME ?= rds_exporter
REPO ?= percona/$(NAME)
GORELEASER_FLAGS ?=
UID ?= $(shell id -u)

init:                       ## Install linters.
	go build -modfile=tools/go.mod -o bin/gofumports mvdan.cc/gofumpt/gofumports
	go build -modfile=tools/go.mod -o bin/gofumpt mvdan.cc/gofumpt
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	go build -modfile=tools/go.mod -o bin/reviewdog github.com/reviewdog/reviewdog/cmd/reviewdog

build:                      ## Compile using plain go build
	go build -ldflags="$(GO_BUILD_LDFLAGS)"  -o rds_exporter

release:                      ## Build the binaries using goreleaser
	docker run --rm --privileged \
		-v ${PWD}:/go/src/github.com/user/repo \
		-w /go/src/github.com/user/repo \
		--env GOPROXY=https://goproxy.cn \
		goreleaser/goreleaser release --snapshot --skip-publish --rm-dist

FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

format:                     ## Format source code.
	go mod tidy
	bin/gofumpt -w -s $(FILES)
	bin/gofumports -local github.com/eleven26/rds_exporter -l -w $(FILES)

check:                      ## Run checks/linters
	bin/golangci-lint run

check-license:              ## Check license in headers.
	@go run .github/check-license.go

help:                       ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
	awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

test:                   ## Run all tests.
	go test -v -timeout 30s ./...

test-race:              ## Run all tests with race flag.
	go test -race -v -timeout 30s ./...
