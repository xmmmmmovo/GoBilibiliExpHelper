PROJECT_NAME := "."
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
EXCLUDE_DIR := vendor,.idea

.PHONY: all dep lint vet test test-coverage build clean cloc

all: build

cloc: ## cloc code lines
	@cloc ${PROJECT_NAME} --exclude-dir ${EXCLUDE_DIR}

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -race -coverprofile=coverage.txt -covermode=atomic ${PKG_LIST}

build: dep ## Build the binary file
	@go build -i -o build/main $(PKG)

clean: ## Remove previous build
	@rm -f ./build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'