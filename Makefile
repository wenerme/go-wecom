# BAZEL=BAZELISK_BASE_URL=https://mirrors.huaweicloud.com/bazel `go env GOPATH`/bin/bazelisk
BAZEL=`go env GOPATH`/bin/bazelisk

ci:
	go env
	@$(MAKE) go-test-cover
	#command -v bazel > /dev/null || GOPROXY=https://goproxy.io,direct go install github.com/bazelbuild/bazelisk@latest
	#$(BAZEL) info

.PHONY: go-test-cover
go-test-cover: ## run test & generate coverage
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

### BEGIN COMMON
COLOR 	:= "\e[1;36m%s\e[0m\n"
RED 	:=   "\e[1;31m%s\e[0m\n"

ifdef GOROOT
PATH 	:= $(GOROOT)/bin:$(PATH)
endif

GOBIN 	:= $(if $(shell go env GOBIN),$(shell go env GOBIN),$(GOPATH)/bin)
PATH 	:= $(GOBIN):$(PATH)

GOOS 	?= $(shell go env GOOS)
GOARCH 	?= $(shell go env GOARCH)
GOPATH 	?= $(shell go env GOPATH)

MODDIR	:= $(shell dirname $(shell go env GOMOD))

##### Bazel #####

git-add-bazel:
	git add --ignore-error '**/BUILD.bazel' WORKSPACE deps.bzl BUILD.bazel

##### Bazel Docker #####

image: # build image
	@[ ! -e BUILD.bazel ] || bazel run $(shell bazel query 'kind(container_image, //build/...)')

image-ls: image
	tar -tvf bazel-bin/build/server-layer.tar

push: ## push image to registry
	@[ ! -e BUILD.bazel ] || bazel run $(shell bazel query 'kind(container_push, //build/...)')

##### Bazel Go #####

gazelle: ## bazel gazelle
	@bazel run --noshow_progress --noshow_loading_progress //:gazelle
	git add --ignore-error '**/BUILD.bazel'

gazelle-update-repos: tidy  ## bazel gazelle update-repos
	@bazel run --noshow_progress --noshow_loading_progress //:gazelle-update-repos

update: gazelle-update-repos gazelle

.PHONY: build
build: ## build binary
	@[ ! -e BUILD.bazel ] || bazel build $(shell bazel query 'kind(go_binary, //cmd/...)')
.PHONY: build
test: ## bazel test
	@[ ! -e BUILD.bazel ] || bazel test $(shell bazel query 'kind(go_test, //...)')

##### Golang #####

.PHONY: lint
lint: ## lint
	golangci-lint run

.PHONY: fmt
fmt: tidy ## tidy,format and imports
	gofumpt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
	goimports -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy

.PHONY: gen
gen: ## generate
	[ -e buf.gen.yaml ] && buf generate || true
	$(MAKE) fmt

go-info:
	@go version

update-dependencies: ## update go dependencies
	@printf $(COLOR) "Update dependencies..."
	@go get -u -t $(PINNED_DEPENDENCIES) ./...
	@go mod tidy

ensure-no-changes: ## ensure git doesn't have any changes
	@printf $(COLOR) "Check for local changes..."
	@printf $(COLOR) "========================================================================"
	@git diff --name-status --exit-code || (printf $(COLOR) "========================================================================"; printf $(RED) "Above files are not regenerated properly. Regenerate them and try again."; exit 1)

.PHONY: help
.DEFAULT_GOAL := help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
### END COMMON

docker:
	docker run --rm -it \
		-e GO111MODULE=on -e GOPROXY=https://goproxy.io \
		-v ~/go:/root/go \
		-v ~/.cache:/root/.cache \
		-v $(PWD):/host \
		--workdir /host \
		-e HOME=/root\
		--name go golang:1.19-bullseye bash

.PHONY: bin
bin:
	go build -o bin/wwfinance-libs ./cmd/wwfinance-libs
	go build -o bin/wwfinance ./cmd/wwfinance
