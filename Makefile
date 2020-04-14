# ref. 改訂2版 みんなのGo言語
NAME := calico-server
VERSION := $(godump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## install dependencies
.PHONY: deps
deps:
	go get -v -d

## setup
.PHONY: deps
devel-deps: deps
	GO111MODULE=off go get \
	github.com/gin-gonic/gin	\
	github.com/go-sql-driver/mysql \
	github.com/Songmu/make2help/cmd/make2help

## test
.PHONY: test
test: deps
	go test ./...

## lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status -min_confidence=0.1 ./...

## build binaries
# bin/%: main.go deps
# 	go build -ldflags "$(LDFLAGS)" -o $@ $<

## build binary
.PHONY: build
# build: bin/%
build:
	go build -o bin/$(NAME)/main -v

# run
.PHONY: run
run: build
	bin/$(NAME)/main

## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
