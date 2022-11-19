# BAZEL=BAZELISK_BASE_URL=https://mirrors.huaweicloud.com/bazel `go env GOPATH`/bin/bazelisk
BAZEL=`go env GOPATH`/bin/bazelisk

ci:
	go env
	@$(MAKE) go-test-cover
	#command -v bazel > /dev/null || GOPROXY=https://goproxy.io,direct go install github.com/bazelbuild/bazelisk@latest
	#$(BAZEL) info

.PHONY: go-test-cover
go-test-cover: ## run test & generate coverage
	LD_LIBRARY_PATH=$(PWD)/WeWorkFinanceSDK/libs go test -race -coverprofile=cover.out -coverpkg=./... ./...
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
		-e GO111MODULE=on -e GOPROXY=https://goproxy.io -e GOPATH=/root/go \
		-v ~/go:/root/go \
		-v ~/.cache:/root/.cache \
		-v $(PWD):/host \
		--workdir /host \
		-e HOME=/root\
		--name go golang:1.19-bullseye bash

.PHONY: bin
bin:
	CGO_ENABLED=0 go build -o bin/wwfinance-libs -trimpath -ldflags "-s -w" github.com/wenerme/go-wecom/cmd/wwfinance-libs
	go build -o bin/wwfinance-poller -trimpath -ldflags "-s -w" github.com/wenerme/go-wecom/cmd/wwfinance-poller

install:
	go install mvdan.cc/gofumpt@latest

image:
	docker build --pull -t wener/go-wecom .

run-image:
	docker run --rm -it -v $(PWD)/.env:/app/.env -v $(PWD)/data:/app/data wener/go-wecom
introspect-image:
	docker run --rm -it --entrypoint bash --env-file .env -v $(PWD)/data:/app/data wener/go-wecom
run-poller:
	LD_LIBRARY_PATH=$(PWD)/WeWorkFinanceSDK/libs go run ./cmd/wwfinance-poller/main.go
