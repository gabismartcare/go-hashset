SHELL := /bin/bash
PKG := github.com/mathmasa/go-hashset
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

build:
	@go build $(GO_FILES)

lint:
	@golint -set_exit_status .
	@go vet .
	@go fmt .

test:
	@go test -short .

coverage:
	@go test -coverprofile=cov.out .
	@go tool cover -func=cov.out