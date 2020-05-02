SHELL := /bin/bash
PKG := "github.com/thayen/sample-rest"

fmt:
	@go fmt $(shell go list ${PKG}/... | grep -v /vendor/)

test:
	@go test $(shell go list ${PKG}/... | grep -v /vendor/)
	@go test -coverprofile=cov.out $(shell go list ${PKG}/... | grep -v /vendor/)
	@go tool cover -func=cov.out

dep:
	@go mod download

build:
	@go build

new-service:
	@read -p "Enter module name: " module; \
	mkdir -p $$module/domain/model; \
	mkdir -p $$module/domain/repository; \
	mkdir -p $$module/domain/service; \
	mkdir -p $$module/interface/persistence; \
	mkdir -p $$module/interface/protobuf; \
	mkdir -p $$module/interface/http; \
	mkdir -p $$module/registry; \
	mkdir -p $$module/usecase; \
	touch $$module/main.go; \
	touch $$module/Makefile; \
	cd $$module;

feature-start :
	@read -p "Enter feature name: " module; \
	git flow feature start $$module

hotfix-start :
	@read -p "Enter version: " module; \
	git flow hotfix start $$module

release-start :
	@read -p "Enter version: " module; \
	git flow release start $$module

support-start :
	@read -p "Enter version: " module; \
	git flow support start $$module

feature-finish : fmt test
	git flow feature finish -S

hotfix-finish : fmt test
	git flow hotfix finish -S

release-finish : fmt test
	git flow release finish -S
