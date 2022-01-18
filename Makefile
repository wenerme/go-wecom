# BAZEL=BAZELISK_BASE_URL=https://mirrors.huaweicloud.com/bazel `go env GOPATH`/bin/bazelisk
BAZEL=`go env GOPATH`/bin/bazelisk

GOROOT 	?= `go env GOROOT`
PATH 	:= $(GOROOT):$(PATH)

ci:
	go env
	@$(MAKE) go-test-cover
	#command -v bazel > /dev/null || GOPROXY=https://goproxy.io,direct go install github.com/bazelbuild/bazelisk@latest
	#$(BAZEL) info

##### Common #####

COLOR 	:= "\e[1;36m%s\e[0m\n"
RED 	:=   "\e[1;31m%s\e[0m\n"

MODULE_ROOT := $(lastword $(shell grep -e "^module " go.mod))

GOBIN 	:= $(if $(shell go env GOBIN),$(shell go env GOBIN),$(GOPATH)/bin)
PATH 	:= $(GOBIN):$(PATH)

GOOS 	?= $(shell go env GOOS)
GOARCH 	?= $(shell go env GOARCH)
GOPATH 	?= $(shell go env GOPATH)

##### Bazel Docker #####

image: # build image
	bazel run $(shell bazel query 'kind(container_image, //build/...)')

image-ls: image
	tar -tvf bazel-bin/build/server-layer.tar

push:
	bazel run $(shell bazel query 'kind(container_push, //build/...)')

##### Bazel Go #####

gazelle:
	bazel run //:gazelle
gazelle-update-repos: tidy
	bazel run //:gazelle-update-repos

.PHONY: build
build:
	bazel build $(shell bazel query 'kind(go_binary, //cmd/...)')

.PHONY: test
test:
	bazel test $(shell bazel query 'kind(go_test, //...)')

##### Golang #####

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt: tidy ## tidy,format and imports
	gofumpt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
	goimports -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy

.PHONY: go-test-cover
go-test-cover: ## run test & generate coverage
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: gen
gen: ## generate
	[ -e buf.gen.yaml ] && buf generate || true
	$(MAKE) fmt

update-dependencies:
	@printf $(COLOR) "Update dependencies..."
	@go get -u -t $(PINNED_DEPENDENCIES) ./...
	@go mod tidy

ensure-no-changes:
	@printf $(COLOR) "Check for local changes..."
	@printf $(COLOR) "========================================================================"
	@git diff --name-status --exit-code || (printf $(COLOR) "========================================================================"; printf $(RED) "Above files are not regenerated properly. Regenerate them and try again."; exit 1)

.PHONY: help
.DEFAULT_GOAL := help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
